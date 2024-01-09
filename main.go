package main

import (
	"fmt"
	"log"
	"net/http"
	"groupie-tracker/server"
)

func main() {
	port := "8080"
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/filter", server.HandlerFilter)
	http.HandleFunc("/static/", server.ServeStatic)
	http.HandleFunc("/artist", server.ArtistHandler)
	http.HandleFunc("/suggest", server.SuggestHandler)
	http.HandleFunc("/search", server.SearchHandler)

	fmt.Printf("Server listeneing on port: %s\n", port)
	fmt.Printf("http://localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
