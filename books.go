package main
 
import (
	"net/http"
)
type BookHandler struct {

}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request){}

func (b BookHandler) GetBook(w http.ResponseWriter, r *http.Request){}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request){}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request){}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request){}