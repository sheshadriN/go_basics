package main

import "fmt"

func main() {
    // Defer a function call to handle any panics that occur within the main function
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    // Simulate a panic by dividing by zero
    panic("Oops! Something went wrong")
}
