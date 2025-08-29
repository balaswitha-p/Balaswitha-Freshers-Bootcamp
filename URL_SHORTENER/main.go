package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
	Error    string `json:"error,omitempty"`
}

var db *sql.DB

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS urls (
	short_code TEXT PRIMARY KEY,
	original_url TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Database initialised successfully.")
}

func generateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invlaid request body", http.StatusBadRequest)
		return
	}
	if req.URL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}
	const maxAttempts = 5
	var shortCode string
	for i := 0; i < maxAttempts; i++ {
		shortCode = generateShortCode(6)
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE short_code=?)", shortCode).Scan(&exists)
		if err != nil {
			log.Printf("Error checking short code existence: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !exists {
			break
		}
		if i == maxAttempts-1 {
			http.Error(w, "Could not generate unique short code, please try again", http.StatusInternalServerError)
			return
		}
	}

	_, err = db.Exec("INSERT INTO urls (short_code,original_url) VALUES (?,?)", shortCode, req.URL)
	if err != nil {
		log.Printf("Error inserting URL into database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fullShortURL := fmt.Sprintf("http//localhost:8080/%s", shortCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ShortenResponse{ShortURL: fullShortURL})

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	if shortCode == "" {
		http.Error(w, "Short code not provided", http.StatusBadRequest)
		return
	}
	var originalURL string
	err := db.QueryRow("SELECT original_url FOM urls WHERE short_code=?", shortCode).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Short URL not found", http.StatusNotFound)
		} else {
			log.Printf("Error querying database for shirt cide %s:%v", shortCode, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}

func main() {
	initDB("./urls.db")
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)
	port := ":8080"
	fmt.Printf("URL Shortener service starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
