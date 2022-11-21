package main

import (
	"net/http"

	"github.com/maple-soft/shared"
)

type userHandler struct{}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// all users request are going to be routed here
	shared.Log("looking up")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/pay/", &userHandler{})
	http.ListenAndServe(":8080", mux)
}
