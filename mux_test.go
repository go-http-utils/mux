package mux

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MuxSuite struct {
	suite.Suite

	server *httptest.Server
}

func (s *MuxSuite) SetupSuite() {
	mux := New()

	mux.Get("/get/:id", HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(params["id"]))
	}))

	mux.Post("/post/:id", HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(params["id"]))
	}))

	mux.Put("/put/:id", HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(params["id"]))
	}))

	mux.Delete("/delete/:id", HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(params["id"]))
	}))

	mux.Head("/head/:id", HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(params["id"]))
	}))

	mux.Patch("/patch/:id", HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(params["id"]))
	}))

	s.server = httptest.NewServer(mux)
}

func (s *MuxSuite) TestMethods() {
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodHead,
		http.MethodPatch,
	}

	for _, method := range methods {
		req, err := http.NewRequest(method, s.server.URL+"/"+strings.ToLower(method)+"/123", nil)

		s.Nil(err)

		res, err := sendRequest(req)

		s.Nil(err)
		s.Equal(http.StatusOK, res.StatusCode)

		if method != http.MethodHead {
			s.Equal([]byte("123"), getResRawBody(res))
		}
	}
}

func TestMux(t *testing.T) {
	suite.Run(t, new(MuxSuite))
}

func sendRequest(req *http.Request) (*http.Response, error) {
	cli := &http.Client{}
	return cli.Do(req)
}

func getResRawBody(res *http.Response) []byte {
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return bytes
}
