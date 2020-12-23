package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server run on port: 8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Println(err)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
