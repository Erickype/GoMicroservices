package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello word!!")
		data, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Oops", http.StatusBadRequest)
			return
		}
		_, _ = fmt.Fprintf(writer, "Hello %s!", data)
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
