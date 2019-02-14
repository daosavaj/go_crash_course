package main

import "fmt"

func main() {

	//emails := make(map[int] string)

	emails := map[int]string{6: "hello"}

	emails[1] = "1@gmail.com"
	emails[2] = "3@gmail.com"
	emails[3] = "2@gmail.com"

	fmt.Println(emails)
	delete(emails, 2)
	fmt.Println(emails)

}
