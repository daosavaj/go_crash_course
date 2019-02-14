package main

import "fmt"

func main() {

	//i:=1

	//for i<=10 {
	//	fmt.Println(i)
	//	i++
	//}
	//
	//for j:=1;j<10;j++{
	//	fmt.Println(j)
	//}

	//FizzBuzz 3 5

	for k := 0; k < 100; k++ {
		if k%3 == 0 {
			fmt.Print("Fizz")
		}
		if k%5 == 0 {
			fmt.Print("Buzz")
		}
		if k%3 != 0 && k%5 != 0 {
			fmt.Print(k)
		}
		fmt.Println("")
	}
}
