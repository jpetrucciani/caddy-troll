package troll

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var b Troll
	err := b.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return b, err
}

// func parseStringArg(d *caddyfile.Dispenser, out *string) error {
// 	if !d.Args(out) {
// 		return d.ArgErr()
// 	}
// 	return nil
// }
