package main

import "fmt"

type object_value struct {
	objName string
	id      int
	radius  float64
	height  float64
}

type inter interface {
	name() string
	area() float64
}

func (v object_value) name() string {
	return v.objName
}

func (v object_value) area() float64 {
	return 2 * v.height * v.radius
}

func main() {
	object1 := object_value{"table", 12, 12, 1} 
	fmt.Println( object_value{"table", 12, 12, 1} .name())
	fmt.Println(object1.area())
}
