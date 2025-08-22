package main

import "fmt"

func main() {
	fmt.Println("Maps")
	lang := make(map[string]string)
	lang["Fruit"] = "Apple"
	lang["Vege"] = "Tomato"
	lang["Movie"] = "War"
	lang["Drink"] = "Tea"

	fmt.Println(lang)

	fmt.Println(lang["Drink"])

	delete(lang, "Vege")
	fmt.Println(lang)

	for key, value := range lang {
		fmt.Printf("%v,%v", key, value)
	}
}
