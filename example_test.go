package mux_test

import (
	"fmt"
	"net/http"

	"github.com/go-http-utils/mux"
)

func Example() {
	m := mux.New()

	m.Get("/:type(a|b)/:id", mux.HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
		res.WriteHeader(http.StatusOK)

		fmt.Println(params["type"])
		fmt.Println(params[":id"])

		res.Write([]byte("Hello Worlkd"))
	}))

	http.ListenAndServe(":8080", m)
}
