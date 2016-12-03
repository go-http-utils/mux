package mux

import (
	"net/http"
	"strings"

	"github.com/DavidCai1993/routing"
	"github.com/go-http-utils/headers"
)

// Version is this package's version number.
const Version = "0.0.1"

// Handler responds to an HTTP request.
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request, map[string]string)
}

// Mux is the HTTP request multiplexer.
type Mux struct {
	root map[string]*routing.Node
}

// New returns a new mux.
func New() *Mux {
	return &Mux{}
}

// Handle registers the handler for the given method and url.
func (m *Mux) Handle(method string, url string, handler Handler) *Mux {
	if _, ok := m.root[method]; !ok {
		m.root[method] = routing.New()
	}

	m.root[method].Define(url, handler)

	return m
}

func (m Mux) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	method := req.Method

	if method == http.MethodOptions {
		allows := []string{}
		res.WriteHeader(http.StatusNoContent)

		for m, n := range m.root {
			if _, _, ok := n.Match(uri); ok {
				allows = append(allows, m)
			}
		}

		res.Header().Add(headers.Allow, strings.Join(allows, ", "))

		return
	}

	if method == http.MethodHead {
		method = http.MethodGet
	}

	node, ok := m.root[method]

	if !ok {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))

		return
	}

	handler, params, ok := node.Match(uri)

	if !ok {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(http.StatusText(http.StatusNotFound)))

		return
	}

	handler.(Handler).ServeHTTP(res, req, params)
}
