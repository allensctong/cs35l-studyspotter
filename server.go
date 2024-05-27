package main

import (
//	"fmt"
	"net/http"
	"github.com/rs/cors"
)

type userHandler struct{}

func (userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello world\nMethod type:%s\n", r.Method)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write([]byte("{\"GET\": \"hello world\"}\n"))
	} else if r.Method == "POST" {
		w.Write([]byte("{\"POST\": \"hello world\"}\n"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/user", userHandler{})
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}\n"))
	})
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
