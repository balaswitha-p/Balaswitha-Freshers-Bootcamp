package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[4] = "Mango"
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	vege := [2]string{"Potato", "Tomato"}
	fmt.Println(vege)
	fmt.Println(len(vege))
}
