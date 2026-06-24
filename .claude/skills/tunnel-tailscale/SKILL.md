---
name: tunnel-tailscale
description: Expose a local togo app publicly via Tailscale Funnel — set TUNNEL_DRIVER=tailscale and call tunnel.Start
---

# togo tunnel-tailscale

Tailscale Funnel driver for the togo `tunnel` subsystem.

## Setup

```bash
togo install togo-framework/tunnel
togo install togo-framework/tunnel-tailscale
```

1. Install the `tailscale` CLI and `tailscale up` (log in).
2. Enable **Funnel** for the tailnet (https://tailscale.com/kb/1223/funnel).
3. `.env`:
   ```bash
   TUNNEL_DRIVER=tailscale
   ```

## Use

```go
import (
	_ "github.com/togo-framework/tunnel"
	_ "github.com/togo-framework/tunnel-tailscale"
	"github.com/togo-framework/tunnel"
)

if tn, ok := tunnel.FromKernel(k); ok {
	url, _ := tn.Start(ctx, "8080") // https://<host>.<tailnet>.ts.net
	defer tn.Stop(ctx)
}
```

## Notes
- Funnel serves HTTP from your local port at the tailnet `*.ts.net` URL.
- Set `TAILSCALE_FUNNEL_URL` to skip URL discovery in headless/CI runs.
- The funnel stops when the process is stopped (`tn.Stop`).
