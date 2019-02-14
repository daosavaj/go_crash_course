package main

import "fmt"

type Position struct {
	lon string
	lat string
}

func main() {
	ids := []int{11, 22, 33, 44, 55, 66, 87, 98}

	for index, id := range ids {
		fmt.Printf("%d - ID: %d\n", index, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Sum
	//sum := 0

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	fmt.Println(sum(ids))

	pos := Position{
		lon: "40",
		lat: "40",
	}

	fmt.Println(pos.lat)

}

func sum(values []int) int {

	value := 0

	for _, id := range values {
		value += id
	}

	return value
}
