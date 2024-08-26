package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple","Orange","Anjeer"}
	fmt.Printf("Type of fruitList is %T\n",fruitList)

	fruitList = append(fruitList, "Mango","Banana")
	fmt.Println("fruitList",fruitList)

	fruitListCpy := append(fruitList[1:]) // last range is non-inclusive
	fmt.Println("fruitListCpy",fruitListCpy)

	fruitListCpy2 := append(fruitList[:3]) // last range is non-inclusive
	fmt.Println("fruitListCpy2",fruitListCpy2)

	fruitListCpy3 := append(fruitList[1:3]) // last range is non-inclusive
	fmt.Println("fruitListCpy3",fruitListCpy3)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
    
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)
    fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
    fmt.Println(sort.IntsAreSorted(highScores))

	
}