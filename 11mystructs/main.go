package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No Inheritance in Golang, No super or parent

	 manthan := User{"Manthan","manthan@go.dev",true,25}
	 fmt.Println("User details:",manthan)
	 fmt.Printf("manthan details are :%+v\n",manthan)
	 fmt.Printf("Name is %v and email is %v \n",manthan.Name, manthan.Email)

}

// First Letter Capital is Important for a struct to be public, because it might be exported
type User struct{
	Name string
	Email string
	Status bool
	Age int
}