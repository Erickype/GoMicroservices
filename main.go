package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello word!!")
		data, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("Data: %s\n", data)
		_, err = writer.Write(data)
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
