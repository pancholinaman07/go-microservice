package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World...")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(rw, "hellpoooo naman %s", d)
}