package main

import "fmt"

func main() {
	fmt.Println("Slices")

	var fruits = []string{"Apple", "Banana"}
	fmt.Println(fruits)
	fmt.Printf("%T\n", fruits)

	fruits = append(fruits, "Hello", "World")
	fmt.Println(fruits)

	fruits = append(fruits[1:])
	fmt.Println(fruits)

	high := make([]int, 4)
	fmt.Println(high)
	fmt.Printf("%T", high)

	var courses = []string{"A", "B", "C", "D"}
	fmt.Println(courses)

	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses)
}
