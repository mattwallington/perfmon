package main

import (
	"code.google.com/p/gorest"
	"net/http"
)

func main() {
	// Start API service.
	gorest.RegisterService(new(ApiService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":31337", nil)

	// Start listening to sensors.
}
