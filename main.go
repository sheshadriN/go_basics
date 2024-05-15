package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	array()
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


func array(){
	// var arr = [10]string{"1","2","3","4"}
	arr2 := [10]int{1,2,3,4,5,6,7,8,9,0}
	for i:=0;i<10;i++{
		fmt.Print((arr2[i]))
	}
}
