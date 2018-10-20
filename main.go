package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World!"}`))
	})
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Printf("[main] ListenAndServe error: %v", err)
	}
}
