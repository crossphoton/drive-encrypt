package http

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/crossphoton/drive-encrypt/src"
	"github.com/gorilla/mux"
)

func cmdMux(mux *mux.Router) {
	mux.Use(authMiddleware)
	mux.PathPrefix("/list/{path:.*}").Methods("GET").HandlerFunc(list)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("success", r.URL)
	})
	return
}

func list(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(mux.Vars(r)["path"])
	log.Println("list -", path)

	files, err := src.ListPath(path, global_password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(files)
	w.Write(data)
}
