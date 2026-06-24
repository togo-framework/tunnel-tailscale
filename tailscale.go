// Package tailscale is a togo tunnel driver for Tailscale Funnel. It wraps the
// `tailscale` CLI: Funnel publishes a local port on your tailnet's MagicDNS name
// to the public internet over HTTPS.
//
// Install: `togo install togo-framework/tunnel-tailscale`, set TUNNEL_DRIVER=tailscale.
// Requires the `tailscale` CLI, an authenticated node, and Funnel enabled for the
// tailnet (https://tailscale.com/kb/1223/funnel).
package tailscale

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/togo-framework/togo"
	"github.com/togo-framework/tunnel"
)

func init() {
	tunnel.RegisterDriver("tailscale", func(k *togo.Kernel) (tunnel.Tunnel, error) {
		return &driver{bin: envOr("TAILSCALE_BIN", "tailscale")}, nil
	})
}

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

type driver struct {
	bin string

	mu      sync.Mutex
	running bool
	url     string
}

type tsStatus struct {
	Self struct {
		DNSName string `json:"DNSName"`
	} `json:"Self"`
}

// funnelURL turns a tailscale MagicDNS name (with trailing dot) into the public
// Funnel URL.
func funnelURL(dnsName string) string {
	host := strings.TrimSuffix(dnsName, ".")
	if host == "" {
		return ""
	}
	return "https://" + host
}

// parseDNSName extracts Self.DNSName from `tailscale status --json` output.
func parseDNSName(jsonOut []byte) (string, error) {
	var s tsStatus
	if err := json.Unmarshal(jsonOut, &s); err != nil {
		return "", err
	}
	return s.Self.DNSName, nil
}

func (d *driver) Start(ctx context.Context, addr string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	port := tunnel.PortOf(addr)

	// Enable Funnel to the local port in the background.
	cmd := exec.CommandContext(ctx, d.bin, "funnel", "--bg", port)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("tunnel-tailscale: funnel --bg %s: %w: %s", port, err, strings.TrimSpace(string(out)))
	}

	statusOut, err := exec.CommandContext(ctx, d.bin, "status", "--json").Output()
	if err != nil {
		return "", fmt.Errorf("tunnel-tailscale: status: %w", err)
	}
	dnsName, err := parseDNSName(statusOut)
	if err != nil {
		return "", err
	}
	url := funnelURL(dnsName)
	if url == "" {
		return "", fmt.Errorf("tunnel-tailscale: could not determine the tailnet DNS name")
	}
	d.running = true
	d.url = url
	return url, nil
}

func (d *driver) Stop(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Best-effort reset of serve/funnel config.
	_ = exec.CommandContext(ctx, d.bin, "funnel", "reset").Run()
	_ = exec.CommandContext(ctx, d.bin, "serve", "reset").Run()
	d.running = false
	d.url = ""
	return nil
}

func (d *driver) Status(context.Context) (tunnel.Status, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return tunnel.Status{Running: d.running, URL: d.url}, nil
}
