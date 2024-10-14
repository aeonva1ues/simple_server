package middleware

import (
	"fmt"
	"net/http"
)

func NewLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s \n", r.Method, r.URL.Path)
		next.ServeHTTP(rw, r)
	})
}