package caddysave

import (
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/caddyserver/caddy/v2"
    "github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
    "github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
    "github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
    caddy.RegisterModule(Save{})
    httpcaddyfile.RegisterHandlerDirective("save", parseCaddyfile)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Save
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

type Save struct {
    FilePath string `json:"file_path,omitempty"`
}

func (Save) CaddyModule() caddy.ModuleInfo {
    return caddy.ModuleInfo{
        ID:  "http.handlers.save",
        New: func() caddy.Module { return new(Save) },
    }
}

func (s Save) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
    file, err := os.Create(s.FilePath)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    _, err = io.Copy(file, r.Body)
    if err != nil {
        return fmt.Errorf("failed to write request body to file: %w", err)
    }

    return next.ServeHTTP(w, r)
}

func (s *Save) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
    for d.Next() {
        if !d.NextArg() {
            return d.ArgErr()
        }
        s.FilePath = d.Val()
    }
    return nil
}

var (
    _ caddy.Module                = (*Save)(nil)
    _ caddyhttp.MiddlewareHandler = (*Save)(nil)
    _ caddyfile.Unmarshaler       = (*Save)(nil)
)

