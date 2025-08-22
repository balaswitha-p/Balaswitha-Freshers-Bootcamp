package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wel := "Welcome to Golang"
	fmt.Println(wel)

	var name string
	var age int
	fmt.Println("name and age are:")
	fmt.Scan(&name, &age)
	fmt.Printf("%s %d", name, age)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Printf(text)
}
