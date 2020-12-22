package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server run on port: 8080")
	err := http.ListenAndServe(":8080", http.HandlerFunc(handlerFunc))

	if err != nil {
		log.Println(err)
	}

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Index Page"))
	default:
		w.Write([]byte("404 Page Not Found"))
	}
}
