package main

import "fmt"

func main() {
	if true && false {
		fmt.Println("This is so true")
	} else if n := "apple"; n == "apple" {
		fmt.Println("OK, We are Here")
	} else {
		fmt.Println("Here's ELSE")
	}

	if true || false {
		fmt.Println("This is also true")
	}
}
