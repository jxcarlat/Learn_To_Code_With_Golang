package main

import "fmt"

func main() {
	product := "T-shirt"
	cost := 20

	fmt.Println("product's value is:", product)
	fmt.Printf("product's type is: %T\n", product)

	cost = 15

	fmt.Println("product's value is:", cost)
	fmt.Printf("product's type is: %T\n", cost)
}
