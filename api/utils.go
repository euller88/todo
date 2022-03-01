package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func writeErrorMessage(status int, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	data := struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Detail string `json:"detail"`
	}{
		Code:   status,
		Status: http.StatusText(status),
		Detail: err.Error(),
	}
	_ = json.NewEncoder(w).Encode(data)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

//newLoggingResponseWriter retorna um LoggingResponseWriter
func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

//WriteHeader é análogo a http.ResponseWriter.WriteHeader, porém expõe o código de status da requisição
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lwr := newLoggingResponseWriter(w)

		inner.ServeHTTP(lwr, r)

		log.Printf(
			"%s %s %d %s %s\n",
			name,
			r.Method,
			lwr.StatusCode,
			r.RequestURI,
			time.Since(start),
		)
	})
}
