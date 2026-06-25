# tunnel-tailscale — docs

**Tailscale Funnel.** Expose a local port over Tailscale Funnel via the `tailscale` CLI.

## Install

```bash
togo install togo-framework/tunnel-tailscale
```

Registers on the [`tunnel`](https://github.com/togo-framework/tunnel) base; select it with **tunnel.provider in togo.yaml (or TUNNEL_DRIVER)**, then use **`togo tunnel`**.

## Interface

`Tunnel` — `Start(ctx, addr) -> publicURL`, `Stop`, `Status`.

## Usage & notes

Requires `tailscale` installed and the node joined to your tailnet with Funnel enabled. Runs `tailscale funnel` and resolves the public HTTPS URL from `tailscale status --json`.

## Example

```bash
togo tunnel:start --provider tailscale
```

## Links

- [Tailscale Funnel](https://tailscale.com/kb/1223/funnel)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/tunnel-tailscale)
