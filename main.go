package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
	})
	log.Println("Web server started!")
	http.ListenAndServe(":80", nil)
}