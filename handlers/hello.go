package handlers

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops, Error!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
	
}