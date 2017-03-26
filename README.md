# overridemethod

[![GoDoc](https://godoc.org/github.com/i2bskn/overridemethod?status.svg)](https://godoc.org/github.com/i2bskn/overridemethod)
[![Build Status](https://travis-ci.org/i2bskn/overridemethod.svg?branch=master)](https://travis-ci.org/i2bskn/overridemethod)
[![codecov](https://codecov.io/gh/i2bskn/overridemethod/branch/master/graph/badge.svg)](https://codecov.io/gh/i2bskn/overridemethod)

## Dependencies

- [Go](https://golang.org/) 1.7 or lator

No dependency on the third party library.

## Installation

```
go get -u github.com/i2bskn/overridemethod
```

## Usage

Example for `net/http`:

```Go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/i2bskn/overridemethod"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Override %s with %s", overridemethod.Origin(r), r.Method)
	})
	mw := overridemethod.Middleware()

	log.Fatal(http.ListenAndServe(":8080", mw(mux)))
}
```

See also [GoDoc](https://godoc.org/github.com/i2bskn/overridemethod).

## License

overridemethod is available under the MIT.