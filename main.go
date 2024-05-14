package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	new()
	var handeler = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %v", r.URL.Path)
	}
	http.HandleFunc("POST /hello", handeler)
	var error = http.ListenAndServe(":9000", nil)
	fmt.Println("sever is running at port 9000")
	if error!=nil{
		log.Fatal(error)
	}
}
func new(){
	var handeler = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %v", r.URL.Path)
	}
	http.HandleFunc("POST /hello", handeler)
	var error = http.ListenAndServe(":9001", nil)
	fmt.Println("sever is running at port 9000 second")
	if error!=nil{
		log.Fatal(error)
	}
}
