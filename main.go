package main

import "fmt"
import "net/http"

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "a...bbbbbbbbbbb")
}

func Xixi(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "b..nnnnnnnnnnnnn.")
}

func main() {
	http.HandleFunc("/kkk", Hello)
	http.HandleFunc("/vvv", Xixi)
	http.ListenAndServe(":8331", nil)
}