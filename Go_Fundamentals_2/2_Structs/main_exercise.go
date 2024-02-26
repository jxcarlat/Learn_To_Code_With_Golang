package main

import "fmt"

type Student struct {
	id   int
	name string
}

type Classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []Student
}

func main() {
	c1 := Classroom{id: 1, capacity: 20, subject: "Computer Science", studentList: []Student{{id: 1, name: "Jeremy"}, {id: 2, name: "Jessie"},},}
	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 100
	c2.subject = "Spanish"
	c2.studentList = []Student {
		{
			id: 30,
			name: "Vinny",
		},
		{
			id: 35,
			name: "Jhon",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)
}