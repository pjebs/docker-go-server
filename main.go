package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

// handler is the http server's endpoint handler.
func handler(w http.ResponseWriter, r *http.Request) {
	output := "hello world"
	w.Write([]byte(output))
}
