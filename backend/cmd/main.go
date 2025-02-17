package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Message string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Message)
}

func startServer(url string) {
	mux := http.NewServeMux()
	h := &Handler{"Hello from server!"}
	mux.Handle("/", h)
	mux.HandleFunc("/test",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("You are in the /test"))
		})
	http.ListenAndServe(url, mux)
}

func main() {
	fmt.Print("Server is starting...\n")
	startServer(":8080")
}
