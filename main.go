package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server run on port: 8080")
	err := http.ListenAndServe(":8080", &indexHandler{})

	if err != nil {
		log.Println(err)
	}

}

type indexHandler struct{}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(("Hello, GOPHER")))
}
