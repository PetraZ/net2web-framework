package main

import (
	"net/http"

	"github.com/PetraZ/net2web-framework/framework"
)

func main() {
	// handler has a ServeHTTP method to serve all traffic
	handler := framework.NewCore()
	http.ListenAndServe("localhost:8088", handler)
}
