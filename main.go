package main

import (
	"github.com/Erickype/GoMicroservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "bonsai-api", log.LstdFlags)

	helloHandler := handlers.NewHello(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)

	_ = http.ListenAndServe(":9090", serveMux)
}
