package main

import (
	"fmt"
)

func main() {
    welcome := "Welcome to user input"
    fmt.Println(welcome)

    // Taking user input
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}