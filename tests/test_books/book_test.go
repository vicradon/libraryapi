package controller_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vicradon/internpulse/stage4/controller"
)

func setupTestDB() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(`CREATE TABLE books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		genre TEXT,
		isAvailable BOOLEAN,
		summary TEXT,
		edition TEXT
	)`)
	return db
}

func TestCreateBook(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	router := gin.Default()
	router.POST("/books", controller.CreateBook(db))

	t.Run("should create a book successfully", func(t *testing.T) {
		bookJSON := `{"title": "Test Book", "author": "Test Author"}`
		req, _ := http.NewRequest("POST", "/books", strings.NewReader(bookJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, resp.Code)
		}

		var responseData map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}
		if message, ok := responseData["message"].(string); !ok || message != "book created successfully" {
			t.Errorf("Unexpected message: got %v", responseData["message"])
		}
	})
}

func TestGetBooks(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	db := setupTestDB()
	defer db.Close()

	db.Exec(`INSERT INTO books (title, author) VALUES (?, ?), (?, ?)`, "Merlin", "Author Pendragon", "Seeker", "Richard Cipher")

	router := gin.Default()
	router.GET("/books", controller.GetBooks(db))
	router.Use(gin.Recovery())

	t.Run("should fetch all books successfully", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/books", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.Code)
		}

		var responseData map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}
		if message, ok := responseData["message"].(string); !ok || message != "books fetched successfully" {
			t.Errorf("Unexpected message: got %v", responseData["message"])
		}
		if _, ok := responseData["data"]; !ok {
			t.Errorf("Expected data in response, got none")
		}
		data, ok := responseData["data"].([]interface{})
		if !ok {
			t.Errorf("Expected data to be a slice, but got %T", responseData["data"])
			return
		}

		if len(data) != 2 {
			t.Errorf("Expected 2 books in data but got %v", data)
		}
	})
}

func TestGetBook(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	// Insert a sample book with ID 1 for testing
	db.Exec(`INSERT INTO books (title, author) VALUES (?, ?)`, "Test Book", "Test Author")

	router := gin.Default()
	router.GET("/books/:id", controller.GetBook(db))

	t.Run("should fetch a book by ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/books/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.Code)
		}

		var responseData map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}
		if message, ok := responseData["message"].(string); !ok || message != "successfully fetched the book" {
			t.Errorf("Unexpected message: got %v", responseData["message"])
		}
		if _, ok := responseData["data"]; !ok {
			t.Errorf("Expected data in response, got none")
		}
	})
}

func TestDeleteBook(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	// Insert a sample book with ID 1 for testing
	db.Exec(`INSERT INTO books (id, title, author) VALUES (?, ?, ?)`, 1, "Test Book", "Test Author")

	router := gin.Default()
	router.DELETE("/books/:id", controller.DeleteBook(db))

	t.Run("should delete a book successfully", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/books/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.Code)
		}

		var responseData map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}
		if message, ok := responseData["message"].(string); !ok || message != "Book deleted successfully" {
			t.Errorf("Unexpected message: got %v", responseData["message"])
		}
	})
}

func TestUpdateBook(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	db.Exec(`INSERT INTO books (title, author) VALUES (?, ?)`, "Old Title", "Old Author")

	router := gin.Default()
	router.PUT("/books/:id", controller.UpdateBook(db))

	t.Run("should update a book successfully", func(t *testing.T) {
		updateData := `{"title": "Updated Title", "author": "Updated Author"}`
		req, _ := http.NewRequest("PUT", "/books/1", strings.NewReader(updateData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.Code)
		}

		var responseData map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}

		if message, ok := responseData["message"].(string); !ok || message != "Book updated successfully" {
			t.Errorf("Unexpected message: got %v", responseData["message"])
		}

		if data, ok := responseData["data"].(map[string]interface{}); !ok {
			t.Errorf("Expected data in response to be a map, but got %T", responseData["data"])
		} else {
			if data["title"] != "Updated Title" {
				t.Errorf("Expected title to be 'Updated Title', got %v", data["title"])
			}
			if data["author"] != "Updated Author" {
				t.Errorf("Expected author to be 'Updated Author', got %v", data["author"])
			}
		}
	})
}
