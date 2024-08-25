package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok syntax || comma err syntax

	input, _ := reader.ReadString('\n');
	// input, err := reader.ReadString('\n');
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of Rating, %T ", input)

}