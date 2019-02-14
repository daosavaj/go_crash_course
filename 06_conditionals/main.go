package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := []string{"red", "blue", "yellow", "green", "purple", ""}
	rand.Seed(time.Now().UnixNano())
	var randomColor = rand.Intn(len(color) - 0)

	switch color[randomColor] {
	case "red":
		fmt.Printf("color is %v\n", color[randomColor])
	case "blue":
		fmt.Printf("color is %v\n", color[randomColor])
	case "yellow":
		fmt.Printf("color is %v\n", color[randomColor])
	case "green":
		fmt.Printf("color is %v\n", color[randomColor])
	case "purple":
		fmt.Printf("color is %v\n", color[randomColor])
	default:
		fmt.Println("SOL")
	}
}
