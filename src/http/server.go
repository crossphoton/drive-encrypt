package http

import (
	// "log"
	"net/http"
)

func getRouter() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/login", Login)
	return
}
