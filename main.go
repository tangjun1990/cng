package main

import "fmt"
import "net/http"

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, Welcome to go web programming...")
}

func main() {
	http.HandleFunc("/kkk", Hello)
	http.ListenAndServe(":8080", nil)
}