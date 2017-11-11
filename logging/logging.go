package logging

import (
	"log"
	"net/http"
)

//Logger logs request/responses to console
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.UserAgent())
		log.Printf("--> %s %s", r.Method, r.URL.Path)

		lrw := NewLoggingResponseWriter(w)

		next.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		log.Printf("<-- %d %s", statusCode, http.StatusText(statusCode))
	})
}

//LogWriter is a struct to grab status codes for logging
type LogWriter struct {
	http.ResponseWriter
	statusCode int
}

//NewLoggingResponseWriter constructor  for LoggingResponseWriter
func NewLoggingResponseWriter(w http.ResponseWriter) *LogWriter {
	return &LogWriter{w, http.StatusOK}
}

//WriteHeader pulls status message from status code
func (lrw *LogWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
