package main

import (
	"Pprof/internal/app/handlers/hello"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello.Handler)
	http.ListenAndServe(":9090", nil)
}
