package main

import (
	"fmt"
	"net/http"
)

func main() {
	// very simple server
	fmt.Println("http sample")

	http.HandleFunc("/", root)
	http.HandleFunc("/healthz", sample)

	http.ListenAndServe(":8080", nil)
	fmt.Println("server: 0.0.0.0:8080")
}

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Println("root OK")
	w.Write([]byte("very simple server"))
}

func sample(w http.ResponseWriter, req *http.Request) {
	fmt.Println("healthz OK")
	w.Write([]byte("OK"))
}
