package util

import (
	"fmt"
	"log"
	"net/http"
)

//NotFoundHandler logs message before sending response
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("--> %s %s", r.Method, r.URL.Path)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, "Endpoint not found")
		return
	})
}
