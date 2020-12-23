package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server run on port: 8080")
	err := http.ListenAndServe(":8080", http.HandlerFunc(mux))

	if err != nil {
		log.Println(err)
	}

}

func mux(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Page Not Found"))
}
