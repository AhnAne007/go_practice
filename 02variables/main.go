package main

import "fmt"

const LoginToken = "gsfhs" //public
// in this way we shwo that this is a public variable

func main() {
    // Use camelCase for variable names
    var userAge int = 25 // Explicit type declaration
    fmt.Println("User Age:", userAge)
	fmt.Printf("The variable is of type: %T \n", userAge);

    // Use short declaration for local variables
    userName := "Ahnuf" // Type inferred from value
    fmt.Println("User Name:", userName)
    
    // userName = 3
    // the above code will not work as the variable is implicitely declared as string

    // Declare multiple variables in a single line for better readability
    var city, country = "Dubai", "UAE"
    fmt.Println("Location:", city, country)

    // Use type inference where possible
    userHeight, userWeight := 5.9, 70.5
    fmt.Println("Height & Weight:", userHeight, userWeight)

    // Use zero values appropriately (avoid unnecessary initialization)
    var defaultInt int       // Default is 0
    var defaultString string // Default is ""
    var defaultBool bool     // Default is false
    fmt.Println("Default Values:", defaultInt, defaultString, defaultBool)

    // Declare constants with meaningful names
    const piValue float64 = 3.14159
    fmt.Println("Pi Value:", piValue)

    // Use grouped declarations for better organization
    var (
        firstName  = "John"
        lastName   = "Doe"
        isStudent  = true
        birthYear  = 2000
    )
    fmt.Println("Person Info:", firstName, lastName, isStudent, birthYear)

    // Use pointers when needed, avoid unnecessary usage
    var number int = 10
    var numberPtr *int = &number
    fmt.Println("Pointer Value:", numberPtr)   // Prints memory address
    fmt.Println("Pointer Dereference:", *numberPtr) // Prints the value stored at the memory address


    fmt.Println("Login Token:", LoginToken)
    // LoginToken = "1+1" // Due to constant we cannot change it
	fmt.Printf("The variable is of type: %T \n", LoginToken);
}
