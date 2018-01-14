package main

import "fmt"
import "net/http"

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hehheehhe...")
}

func Xixi(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "xixiixixix...")
}

func main() {
	http.HandleFunc("/kkk", Hello)
	http.HandleFunc("/vvv", Xixi)
	http.ListenAndServe(":8080", nil)
}