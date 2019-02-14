package main

import "fmt"

func main() {
	// Arrays are fixed length and type

	var fruitArr [2]string
	var fruitArrAssign = [2]string{"Apple", "Orange"}

	fruitDeclare := [2]string{"peach", "eggplant"}

	var fruitSli []string

	var mapFruit map[int]string

	fruitArr[0] = "Dragon Fruit"
	fruitArr[1] = "Mango"

	fruitSlice := []string{"banana", "strawberry", "grape"}

	fmt.Println(fruitArr, fruitArrAssign, fruitDeclare)

	fmt.Println(fruitSlice, fruitSli)
	fmt.Println(fruitSlice[1:2])
	fmt.Println(mapFruit)

}
