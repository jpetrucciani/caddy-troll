package troll

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

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

func NaughtyResponse(b Troll, res http.ResponseWriter, req *http.Request) {
	if b.DisableNaughtyStrings {
		return
	}

	jsonFile, _ := os.Open("blns.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []string
	json.Unmarshal([]byte(byteValue), &result)
	randomIndex := rand.Intn(len(result))

	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(result[randomIndex]))
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

var serverHeaders = []string{
	"Apache/2.4.1 (Unix)",
	"nginx",
	"cloudflare",
	"Akamai Resource Optimizer",
	"proxygen-bolt",
	"ATS",
}

func RandomServerHeader(b Troll, res http.ResponseWriter, req *http.Request) {
	if b.DisableRandomServerHeader {
		res.Header().Set("Server", "Caddy")
	} else {
		randomIndex := rand.Intn(len(serverHeaders))
		res.Header().Set("Server", serverHeaders[randomIndex])
	}
}

func (b Troll) ServeHTTP(res http.ResponseWriter, req *http.Request, next caddyhttp.Handler) error {
	b.log.Info(
		"TROLL",
		zap.String("ip", req.RemoteAddr),
		zap.String("url", req.URL.String()),
		zap.String("user-agent", req.Header.Get("User-Agent")),
	)

	functions := []func(b Troll, res http.ResponseWriter, req *http.Request){}

	if !b.DisableRedirects {
		functions = append(functions, RedirectLocalhost)
		functions = append(functions, RedirectSelf)
		functions = append(functions, RedirectRickRoll)
	}

	if !b.DisableGzips {
		functions = append(functions, GzipSmall)
		functions = append(functions, GzipLarge)
		functions = append(functions, GzipBomb)
	}

	if !b.DisableXmls {
		functions = append(functions, XMLBomb)
	}

	if !b.DisableRandomServerHeader {
		functions = append(functions, RandomServerHeader)
	}

	// if (!b.DisableNaughtyStrings) {
	// 	functions = append(functions, NaughtyResponse)
	// }

	randomIndex := rand.Intn(len(functions))
	randomFunction := functions[randomIndex]
	randomFunction(b, res, req)

	return next.ServeHTTP(res, req)
}
