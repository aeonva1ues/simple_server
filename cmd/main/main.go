package main

import (
	"fmt"
	"net/http"

	"github.com/aeonva1ues/dockered_go/internal/middleware"
)

func RunServer(host string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  "test_cookie",
			Value: "1",
		})
		fmt.Fprintf(w, "<body>testpage</body>")
	})
	fmt.Println("Server listen on host: " + host)
	http.ListenAndServe(host, middleware.NewLoggerMiddleware(mux))
}

func main() {
	RunServer(":8080")
}
