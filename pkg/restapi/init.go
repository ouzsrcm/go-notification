package restapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	r.HandleFunc("/api/run", SendAllHandler).Methods("GET").Schemes("http")
	return r
}

func Run() {
	srv := &http.Server{
		Handler:      routes(),
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
