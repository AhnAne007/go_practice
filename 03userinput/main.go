package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    welcome := "Welcome to user input"
    fmt.Println(welcome)

    // Method 1: Using fmt.Scanln()
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)

    // Method 2: Using fmt.Scan()
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Println("You are", age, "years old.")

    // Method 3: Using bufio.NewReader() (Handles spaces and multiple words)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your full name: ")
    fullName, _ := reader.ReadString('\n')
    fmt.Println("Your full name is:", fullName)

}


//Will do shift this code in functions practice
// func main() {
//     welcomeMessage()
//     name := getName()
//     fmt.Println("Hello,", name)

//     age := getAge()
//     fmt.Println("You are", age, "years old.")

//     fullName := getFullName()
//     fmt.Println("Your full name is:", fullName)
// }

// // Function to print welcome message
// func welcomeMessage() {
//     fmt.Println("Welcome to user input")
// }

// // Function to get user's name
// func getName() string {
//     var name string
//     fmt.Print("Enter your name: ")
//     fmt.Scanln(&name)
//     return name
// }

// // Function to get user's age
// func getAge() int {
//     var age int
//     fmt.Print("Enter your age: ")
//     fmt.Scan(&age)
//     return age
// }

// // Function to get full name using bufio.NewReader()
// func getFullName() string {
//     reader := bufio.NewReader(os.Stdin)
//     fmt.Print("Enter your full name: ")
//     fullName, _ := reader.ReadString('\n')
//     return fullName
// }