package main

import (
	"fmt"
	"log"
	"net/http"
)

func msgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HandlerFunc!")
}

func main02() {
	mux := http.NewServeMux()

	handler := http.HandlerFunc(msgHandler)

	mux.Handle("/hello", handler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
