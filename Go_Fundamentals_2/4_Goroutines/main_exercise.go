package main

import (
	"fmt"
	"time"
)

func printStrings(strings []string) {
	for _, item := range strings {
		fmt.Println(item)
	}
}


func main() {
	start := time.Now()
	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}
	go printStrings(colorNames)
	go printStrings(colorCodes)
	duration := time.Since(start)
	fmt.Println("Duration of execution: " + duration.String())
	time.Sleep(time.Second)
}