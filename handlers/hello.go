package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{
		logger: logger,
	}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello word!!")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(rw, "Hello %s!", data)
}
