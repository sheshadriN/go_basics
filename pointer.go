package main

import "fmt"

func main(){
	name := "shesha"
	// var ptr *string
	var ptr = &name
	fmt.Println(ptr)
	fmt.Println(name)
}