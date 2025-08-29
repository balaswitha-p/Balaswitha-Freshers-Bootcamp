package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Book represents a book with its details and status
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
}

// Library holds a collection of books
type Library struct {
	Books []Book `json:"books"`
}

const dataFile = "library.json"

// It reads the library data from a JSON file
func loadData() *Library {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Library{}
		}
		panic(fmt.Sprintf("failed to read data file: %v", err))
	}
	var library Library
	if err := json.Unmarshal(data, &library); err != nil {
		panic(fmt.Sprintf("failed to unmarshal JSON: %v", err))
	}
	return &library
}

// It writes the library data to a JSON file
func saveData(library *Library) {
	data, err := json.MarshalIndent(library, "", " ")
	if err != nil {
		panic(fmt.Sprintf("failed to marshal JSON: %v", err))
	}
	if err := ioutil.WriteFile(dataFile, data, 0644); err != nil {
		panic(fmt.Sprintf("failed to write data file: %v", err))
	}
}

// It prints all books in the library
func viewAllBooks(library *Library) {
	if len(library.Books) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	fmt.Println("\n--- All Books ---")
	for _, book := range library.Books {
		fmt.Printf("ID: %d | Title: %s | Author:%s | Status: %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

// It prompts the user for book details and adds a new book to the library
func addBook(library *Library) {
	var title, author string
	fmt.Print("Enter book title:")
	fmt.Scanln(&title)
	fmt.Print("Enter book author:")
	fmt.Scanln(&author)
	id := len(library.Books) + 1

	newBook := Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "available",
	}

	library.Books = append(library.Books, newBook)
	saveData(library)
	fmt.Printf("Book '%s' added successfully.\n", newBook.Title)
}

// It finds a book by its ID and returns a pointer to it
func findBook(library *Library, id int) *Book {
	for i := range library.Books {
		if library.Books[i].ID == id {
			return &library.Books[i]
		}
	}
	return nil
}

// It changes a book's status to "on loan"
func borrowBook(library *Library) {
	var idstr string
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scanln(&idstr)

	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	book := findBook(library, id)
	if book == nil {
		fmt.Println("Book not found")
		return
	}
	if book.Status == "on loan" {
		fmt.Printf("book '%s' is already on loan \n", book.Title)
		return
	}
	book.Status = "on loan"
	saveData(library)
	fmt.Printf("Book '%s' has been successfully borrowed\n", book.Title)
}

// It changes a book's status back to "available"
func returnBook(library *Library) {
	var idstr string
	fmt.Print("Enter book ID to return: ")
	fmt.Scanln(&idstr)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Invalid")
		return
	}
	book := findBook(library, id)
	if book == nil {
		fmt.Println("Book not found")
		return
	}
	if book.Status == "available" {
		fmt.Printf("Book '%s' is already available\n", book.Title)
		return
	}
	book.Status = "available"
	saveData(library)
	fmt.Printf("Book '%s' has been successfully returned\n", book.Title)
}

func main() {
	library := loadData()
	for {
		fmt.Println("--- Library Management System ---")
		fmt.Println("1. View all books")
		fmt.Println("2. Add a new book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice: ")

		var choice string
		fmt.Scanln(&choice)
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			viewAllBooks(library)
		case "2":
			addBook(library)
		case "3":
			borrowBook(library)
		case "4":
			returnBook(library)
		case "5":
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
