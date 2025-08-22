package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	num := 10
	var ptr = &num
	fmt.Println(ptr, *ptr)
}
