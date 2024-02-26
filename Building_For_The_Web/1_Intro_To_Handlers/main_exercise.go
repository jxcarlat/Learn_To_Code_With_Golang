package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to the Home page!")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of most populous cities.\n")
	for _,city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cities", cityList)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}