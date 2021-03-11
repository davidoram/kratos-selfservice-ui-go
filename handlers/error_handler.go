package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler renders a response when an error occurs
func ErrorHandler(w http.ResponseWriter, _ *http.Request, err error) {
	log.Printf("ErrorHandler returning 500")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(500)
	fmt.Fprintln(w, err.Error())
}
