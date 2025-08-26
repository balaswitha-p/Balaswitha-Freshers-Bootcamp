package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var store = make(map[string]string)

var mu sync.Mutex

// handles SET operation
func setHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/set/"):]
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "Key or value cannot be empty", http.StatusBadRequest)
		return
	}
	mu.Lock()
	store[key] = value
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "key '%s' set to value '%s'", key, value)
}

// handles the GET operation
func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/get/"):]
	if key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}
	mu.Lock()
	value, ok := store[key]
	mu.Unlock()
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Value for key '%s' is '%s'", key, value)
}

// handles the DELETE operation
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/delete/"):]
	if key == "" {
		http.Error(w, "Key canot be empty", http.StatusBadRequest)
		return
	}
	mu.Lock()
	_, ok := store[key]
	if !ok {
		mu.Unlock()
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	delete(store, key)
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key '%s' deleted", key)
}

// provides a JSON dump of entire store
func dumpHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	storeCopy := make(map[string]string)
	for k, v := range store {
		storeCopy[k] = v
	}
	mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storeCopy)
}

func main() {
	http.HandleFunc("/set/", setHandler)
	http.HandleFunc("/get/", getHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/dump/", dumpHandler)

	fmt.Println("Starting key-value store server on : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
