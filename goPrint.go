package main

import (
    "fmt"
)

func main() {
    // print message in GO
    fmt.Println("This is a message in Go")

    // Go loop that loops through numbers 1 - 10
    for i := 1; i <= 10; i++ {
        // Go command to calculate square of a number
        square := i * i
        
        // Print the number and its square
        fmt.Printf("Number: %d, Square: %d\n", i, square)
    }
}
