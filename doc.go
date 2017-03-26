/*
Package overridemethod provides HTTP method overriding.

Example for net/http:
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
*/
package overridemethod
