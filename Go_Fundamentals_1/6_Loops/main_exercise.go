package main

import "fmt"

func fizzbuzz(n int) []string {
	var fizzbuzz_strings []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fizzbuzz_strings = append(fizzbuzz_strings, "Fizzbuzz")
		} else if i%3 == 0 {
			fizzbuzz_strings = append(fizzbuzz_strings, "Fizz")
		} else if i%5 == 0 {
			fizzbuzz_strings = append(fizzbuzz_strings, "Fizz")
		} else {
			fizzbuzz_strings = append(fizzbuzz_strings, fmt.Sprint(i))
		}
	}
	return fizzbuzz_strings
}


func main() {
	fmt.Println(fizzbuzz(16))
}