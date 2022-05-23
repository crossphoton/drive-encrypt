package http

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	mux := mux.NewRouter()

	// Login
	mux.HandleFunc("/login", Login).Methods("GET")

	// cmd subrouter
	cmdMux(mux.Methods("GET").PathPrefix("/cmd/").Subrouter())

	return mux
}
