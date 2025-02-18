package main

import (
	"fmt"
	"net/http"

	"github.com/lyfoore/weather-app/internal/router"
)

type Handler struct {
	Message string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Message)
}

func startServer(url string) {
	fmt.Print("Server is starting...\n")
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
	// // startServer(":8080")
	// err := godotenv.Load("../configs/.env")
	// if err != nil {
	// 	fmt.Printf("Error while loading .env: %s", err)
	// }

	// apiKey := os.Getenv("OWM_API_key")
	// fmt.Println("got api")
	// fmt.Println(apiKey)

	router.StartRouter()
}
