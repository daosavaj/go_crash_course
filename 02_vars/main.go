package main

import (
	"fmt"
)

func main() {

	var name = "Jameson"
	var age int32 = 37
	var abool = true

	email, size := "jdoasavanh@gmail.com", 264

	fmt.Println(name, age, size, email, abool)
	fmt.Printf("%T\n%T\n", name, age)
}
