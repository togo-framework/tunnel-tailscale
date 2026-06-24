<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/tunnel-tailscale</h1>
  <p>Tailscale Funnel driver for togo tunnel — publish a local port on your tailnet to the public internet.</p>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/tunnel-tailscale"><img src="https://pkg.go.dev/badge/github.com/togo-framework/tunnel-tailscale.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/tunnel-tailscale
```
<!-- /togo-header -->

**Tailscale Funnel** driver for togo's [`tunnel`](https://github.com/togo-framework/tunnel)
subsystem. Wraps the `tailscale` CLI to publish a local port on your tailnet's
MagicDNS name to the public internet over HTTPS.

Requires the `tailscale` CLI, an authenticated node, and
[Funnel enabled](https://tailscale.com/kb/1223/funnel) for the tailnet.

## Config

| Env | Meaning |
|-----|---------|
| `TUNNEL_DRIVER` | set to `tailscale` |
| `TAILSCALE_BIN` | path to the `tailscale` CLI (default: `tailscale` on PATH) |

```go
svc, _ := tunnel.FromKernel(k)
url, _ := svc.Start(ctx, "8080")   // → https://<host>.<tailnet>.ts.net
defer svc.Stop(ctx)
```

`Start` runs `tailscale funnel --bg <port>` and reads the public URL from
`tailscale status --json`. `Stop` resets the serve/funnel config.

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- togo-sponsors -->
