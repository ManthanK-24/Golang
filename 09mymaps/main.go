package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
    
	fmt.Println("List of all Languages:",languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all Languages:",languages)
   
	// loops are interesting in Golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n",key, value) // %v for value, %T for Type
	}

}