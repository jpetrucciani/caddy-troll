package troll

import (
	"math/rand"
	"net/http"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func WrongContentSmall(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Length", "1000")
	res.Write([]byte(""))
}

func WrongContentLarge(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Length", "10000000000000")
	res.Write(b.billionLaughs)
}

func GzipSmall(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte("WAZAAAA"))
}

func GzipLarge(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "text/html")
	res.Write(b.billionLaughs)
}

func _redirect(b Troll, res http.ResponseWriter, req *http.Request, location string) {
	req.URL.Host = location
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}
	http.Redirect(res, req, req.URL.String(), 301)
}

func RedirectLocalhost(b Troll, res http.ResponseWriter, req *http.Request) {
	_redirect(b, res, req, "127.0.0.1")
}

func RedirectSelf(b Troll, res http.ResponseWriter, req *http.Request) {
	_redirect(b, res, req, req.RemoteAddr)
}

func RedirectRickRoll(b Troll, res http.ResponseWriter, req *http.Request) {
	_redirect(b, res, req, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}

func XMLBomb(b Troll, res http.ResponseWriter, req *http.Request) {
	// https://en.wikipedia.org/wiki/Billion_laughs_attack
	if rand.Intn(100) < 50 {
		res.Header().Set("Content-Type", "text/xml")
	} else {
		res.Header().Set("Content-Type", "application/xml")
	}
	res.Write(b.billionLaughs)
}

func GzipBomb(b Troll, res http.ResponseWriter, req *http.Request) {
	// https://rehmann.co/blog/10-gb-27-kb-gzip-file-present-http-scanners/
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "text/html")
	res.Write(b.gzipBomb)
}

func RandomServerHeader(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "nginx")
}

func (b Troll) ServeHTTP(res http.ResponseWriter, req *http.Request, next caddyhttp.Handler) error {
	b.log.Info(
		"TROLL",
		zap.String("ip", req.RemoteAddr),
		zap.String("url", req.URL.String()),
		zap.String("user-agent", req.Header.Get("User-Agent")),
	)

	RandomServerHeader(b, res, req)

	// RedirectRickRoll(b, res, req)
	// GzipSmall(b, res, req)
	// XMLBomb(b, res, req)
	GzipBomb(b, res, req)
	// RedirectSelf(b, res, req)
	return next.ServeHTTP(res, req)
}
