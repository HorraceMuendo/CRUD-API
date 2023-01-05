package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {

	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":3100", nil))
}
