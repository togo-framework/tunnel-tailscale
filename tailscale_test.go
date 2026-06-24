package tailscale

import (
	"testing"

	"github.com/togo-framework/tunnel"
)

func TestFunnelURL(t *testing.T) {
	cases := map[string]string{
		"myhost.tail1234.ts.net.": "https://myhost.tail1234.ts.net",
		"box.example.ts.net":      "https://box.example.ts.net",
		"":                        "",
		".":                       "",
	}
	for in, want := range cases {
		if got := funnelURL(in); got != want {
			t.Errorf("funnelURL(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestParseDNSName(t *testing.T) {
	js := []byte(`{"Self":{"DNSName":"myhost.tail1234.ts.net.","HostName":"myhost"}}`)
	name, err := parseDNSName(js)
	if err != nil {
		t.Fatalf("parseDNSName: %v", err)
	}
	if name != "myhost.tail1234.ts.net." {
		t.Errorf("got %q", name)
	}
}

func TestDriverRegistered(t *testing.T) {
	found := false
	for _, n := range tunnel.Drivers() {
		if n == "tailscale" {
			found = true
		}
	}
	if !found {
		t.Fatal("tailscale driver not registered on tunnel base")
	}
}
