package troll

import (
	"io"
	"net/http"
	"math/rand"
	
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func WrongContentSmall(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Length", "1000")
	res.Write([]byte(""))
}

func WrongContentLarge(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Length", "10000000000000")
	io.Copy(res, b.billionsOfLol)
}

func GzipSmall(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte("WAZAAAA"))
}

func GzipLarge(b Troll, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "text/html")
	io.Copy(res, b.billionsOfLol)
}

func RedirectLocalhost(b Troll, res http.ResponseWriter, req *http.Request) {
	req.URL.Host = "127.0.0.1"
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}
	http.Redirect(res, req, req.URL.String(), 301)
}

func RedirectOwnIP(b Troll, res http.ResponseWriter, req *http.Request) {
	req.URL.Host = req.RemoteAddr
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}
	http.Redirect(res, req, req.URL.String(), 301)
}

func XMLBomb(b Troll, res http.ResponseWriter, req *http.Request) {
	// https://en.wikipedia.org/wiki/Billion_laughs_attack
	if rand.Intn(100) < 50 {
		res.Header().Set("Content-Type", "text/xml")
	} else {
		res.Header().Set("Content-Type", "application/xml")
	}
	io.Copy(res, b.billionsOfLol)
}

func GzipBomb(b Troll, res http.ResponseWriter, req *http.Request) {
	// https://rehmann.co/blog/10-gb-27-kb-gzip-file-present-http-scanners/
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "text/html")
	io.Copy(res, b.gzipBomb)
}

func (b Troll) ServeHTTP(res http.ResponseWriter, req *http.Request, next caddyhttp.Handler) error {
	fullPath := req.URL.Path
	if fullPath == "" {
		fullPath = "/"
	}

	XMLBomb(b, res, req)
	return nil
}
