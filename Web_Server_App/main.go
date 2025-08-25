package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// renders home.html template
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tepl, err := template.ParseFiles("Templates/home.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Error parsing home.html %v", err)
		return
	}
	tepl.Execute(w, nil)
}

// renders about.html template
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tepl, err := template.ParseFiles("Templates/about.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Error parsing about.html: %v", err)
	}
	tepl.Execute(w, nil)
}

// reders contact.html template
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tepl, err := template.ParseFiles("Templates/contact.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Error parsing contact.html: %v", err)
		return
	}
	tepl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	port := ":8081"
	fmt.Printf("Server starting on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
