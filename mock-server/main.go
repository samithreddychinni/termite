package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/woc-webhook", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		log.Printf("Received WoC Webhook: %s", string(body))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/aiverse-webhook", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		log.Printf("Received AI-Verse Webhook: %s", string(body))
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Mock server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
