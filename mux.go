package mux

import (
	"net/http"
	"strings"

	"github.com/DavidCai1993/routing"
	"github.com/go-http-utils/headers"
)

// Version is this package's version number.
const Version = "0.0.2"

// Handler responds to a HTTP request.
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request, map[string]string)
}

// The HandlerFunc type is an adapter to allow the use of an ordinary function
// as a Handler.
type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

func (f HandlerFunc) ServeHTTP(res http.ResponseWriter, req *http.Request, params map[string]string) {
	f(res, req, params)
}

// Mux is the HTTP request multiplexer.
type Mux struct {
	root map[string]*routing.Node
}

// New returns a new mux.
func New() *Mux {
	return &Mux{root: map[string]*routing.Node{}}
}

// Handle registers the handler for the given method and url.
func (m *Mux) Handle(method string, url string, handler Handler) *Mux {
	method = strings.ToUpper(method)

	if _, ok := m.root[method]; !ok {
		m.root[method] = routing.New()
	}

	m.root[method].Define(url, handler)

	return m
}

// Get registers a handler for the GET http method.
func (m *Mux) Get(url string, handler Handler) *Mux {
	return m.Handle(http.MethodGet, url, handler)
}

// Post registers a handler for the POST http method.
func (m *Mux) Post(url string, handler Handler) *Mux {
	return m.Handle(http.MethodPost, url, handler)
}

// Put registers a handler for the PUT http method.
func (m *Mux) Put(url string, handler Handler) *Mux {
	return m.Handle(http.MethodPut, url, handler)
}

// Delete registers a handler for the DELETE http method.
func (m *Mux) Delete(url string, handler Handler) *Mux {
	return m.Handle(http.MethodDelete, url, handler)
}

// Head registers a handler for the HEAD http method.
func (m *Mux) Head(url string, handler Handler) *Mux {
	return m.Handle(http.MethodHead, url, handler)
}

// Patch registers a handler for the PATCH http method.
func (m *Mux) Patch(url string, handler Handler) *Mux {
	return m.Handle(http.MethodPatch, url, handler)
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
