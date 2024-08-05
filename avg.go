package main

import (
	"fmt"
)

func main() {
	var name string
	var total_courses int
	var subject_grade = make(map[string]float64)

	fmt.Println("Enter your name")
	fmt.Scanln(&name)

	fmt.Println("Enter the number of courses you took")
	fmt.Scanln(&total_courses)

	var subject string
	var grade float64

	for i := 0; i < total_courses; i++ {
		fmt.Println("Enter subject")
		fmt.Scanln(&subject)

		fmt.Println("Enter grade")
		fmt.Scanln(&grade)

		subject_grade[subject] = grade
	}

	var total_grade = 0.0

	for s, g := range subject_grade {
		fmt.Println(s, g)
		total_grade += g
	}

	var avg = total_grade / float64(total_courses)
	fmt.Println("Your average is ", avg)
}
