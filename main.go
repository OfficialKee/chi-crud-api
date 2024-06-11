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

	// using the created routes for the book resource all the routes are mounted to as new router which is then mounted to main router
	r.Mount("/books", BookRoutes())
	

// server on port 3000
	http.ListenAndServe(":3000", r)


}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
    
	bookHandler := BookHandler{}

	r.Get("/",bookHandler.ListBooks)
	r.Post("/",bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBook)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)

	return r
}