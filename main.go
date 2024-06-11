package main 

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
// middleware logger for http reuquestsa
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("OK"))
	})

// server on port 3000
	http.ListenAndServe(":3000", r)

	// i
}