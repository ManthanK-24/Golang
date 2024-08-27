package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")

	loginCount := 10
	var result string
       
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10{
		result = "Watch out"
	} else{
		result = "Exactly 10 loginCount"
	}

	fmt.Println(result)

	if num := 3; num<10{
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is Not less than 10")
	}
}