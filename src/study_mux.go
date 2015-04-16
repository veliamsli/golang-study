package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("hello go world!!")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	s := r.PathPrefix("/products").Subrouter()

	s.HandleFunc("/{key}", ProductsHandler)
	r.HandleFunc("/articles/{category}", ArticlesHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticlesHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

}

func HomeHandler(http.ResponseWriter, *http.Request) {
	fmt.Println("test home handler")
}

func ProductsHandler(http.ResponseWriter, *http.Request) {
	fmt.Println("test product handler")
}

func ArticlesHandler(http.ResponseWriter, *http.Request) {
	fmt.Println("test articles handler")
}
