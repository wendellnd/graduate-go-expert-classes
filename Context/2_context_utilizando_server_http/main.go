package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Started request")
	defer log.Println("Request finished")

	select {
	case <-time.After(5 * time.Second):
		message := "Request successfully processed"
		log.Println(message)
		w.Write([]byte(message))
	case <-ctx.Done():
		message := "Request cancelled by client"
		log.Println(message)
		http.Error(w, message, http.StatusRequestTimeout)
	}
}
