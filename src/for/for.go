package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		fmt.Println("*")
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		if j%2 == 0 {
			fmt.Println("+<><")
		} else {
			fmt.Println("---")
		}
	}
}
