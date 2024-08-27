package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4] string;
	fruitList[0] = "Apple";
	fruitList[1] = "Banana"; 
	fruitList[3] = "Chikoo";
	
	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List Len is: ", len(fruitList))
  
	var vegList = [3] string {"potato", "beans","mushroom"}
	fmt.Println("VegList is: ",vegList)
	fmt.Println("VegList is: ",len(vegList))
}