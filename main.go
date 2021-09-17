package main

import (
	"net/http"

	"github.com/joe-davidson1802/go-pact-provider/handlers"
)

func main() {
	h := http.HandlerFunc(handlers.GetTimeHandler)
	panic(http.ListenAndServe(":8081", h))
}
