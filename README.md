# mux
[![Build Status](https://travis-ci.org/go-http-utils/mux.svg?branch=master)](https://travis-ci.org/go-http-utils/mux)
[![Coverage Status](https://coveralls.io/repos/github/go-http-utils/mux/badge.svg?branch=master)](https://coveralls.io/github/go-http-utils/mux?branch=master)

HTTP mux for Go.

## Installation

```
go get -u github.com/go-http-utils/mux
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/go-http-utils/mux

## Usage

```go
import (
  "github.com/go-http-utils/mux"
)
```

```go
m := mux.New()

m.Get("/:type(a|b)/:id", mux.HandlerFunc(func(res http.ResponseWriter, req *http.Request, params map[string]string) {
  res.WriteHeader(http.StatusOK)

  fmt.Println(params["type"])
  fmt.Println(params[":id"])

  res.Write([]byte("Hello Worlkd"))
}))

http.ListenAndServe(":8080", m)
```
