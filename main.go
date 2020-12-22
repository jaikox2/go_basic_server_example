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
	w.Write([]byte(("Hello, GOPHER")))
}
