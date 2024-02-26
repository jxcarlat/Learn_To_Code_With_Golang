package main

import "fmt"

func main() {
	number := 125
	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number > 0 && number < 100 {
		fmt.Println(number, "is positive")
	} else {
		fmt.Println(number, "is positive and is a large number!")
	}
}