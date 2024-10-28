package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicradon/internpulse/stage4/models"
	response "github.com/vicradon/internpulse/stage4/pkg"
	"github.com/vicradon/internpulse/stage4/utils"
)

// @Summary     Create a book in the library system
// @Description Create a book by supplying a minimum of title and author
// @Tags        books
// @Produce     json
// @Param       book body models.Book true "Book Data"
// @Success     201 {object} models.Book "Book created successfully"
// @Failure     400 {object} response.ErrorResponse "Bad request"
// @Failure     500 {object} response.ErrorResponse "Server error"
// @Router      /api/v1/books [post]
func CreateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			utils.HandleValidationError(c, err)
			return
		}

		if _, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author); err != nil {
			log.Println(err)
			response.Error(c, http.StatusInternalServerError, "Error inserting book to db")
			return
		}

		response.Success(c, http.StatusCreated, "book created successfully")
	}
}

// @Summary     Get all the books in the library system
// @Description Get all the books
// @Tags        books
// @Produce     json
// @Success     200 {object} response.BookSuccessResponse "Books fetched successfully"
// @Failure     500 {object} response.ErrorResponse "Server error"
// @Router      /api/v1/books [get]
func GetBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM books")
		if err != nil {
			log.Println(err)
			response.Error(c, http.StatusInternalServerError, "An error occurred while fetching the books")
			return
		}
		defer rows.Close()

		var books []models.Book
		for rows.Next() {
			var book models.Book
			if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Genre, &book.IsAvailable, &book.Summary, &book.Edition); err != nil {
				log.Println(err)
				response.Error(c, http.StatusInternalServerError, "An error occurred while fetching the books")
				return
			}
			books = append(books, book)
		}

		response.Success(c, http.StatusOK, "books fetched successfully", books)
	}
}

// @Summary     Get a single book in the library
// @Description Get a single book using the id
// @Tags        books
// @Produce     json
// @Success     200 {object} response.BookSuccessResponse "Book details fetched successfully"
// @Failure     500 {object} response.ErrorResponse "Server error"
// @Router      /api/v1/books/:id [get]
func GetBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book models.Book

		if err := db.QueryRow("SELECT * FROM BOOKS WHERE id = ?", id).Scan(&book.Id, &book.Title, &book.Author, &book.Genre, &book.IsAvailable, &book.Summary, &book.Edition); err != nil {
			log.Println(err)
			response.Error(c, http.StatusInternalServerError, "An error occurred while fetching the book")
			return
		}

		response.Success(c, http.StatusOK, "successfully fetched the book", book)
	}
}

// @Summary     Update a book in the library
// @Description Update a single book using the id
// @Tags        books
// @Produce     json
// @Success     200 {object} response.BookSuccessResponse "Book updated successfully"
// @Failure     500 {object} response.ErrorResponse "Server error"
// @Router      /api/v1/books/:id [put]
func UpdateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book models.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			utils.HandleValidationError(c, err)
			return
		}

		result, err := db.Exec("UPDATE books SET title=?, author=? WHERE id = ?", book.Title, book.Author, id)
		if err != nil {
			log.Println(err)
			response.Error(c, http.StatusInternalServerError, "error updating book details in database")
			return
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			response.Error(c, http.StatusBadRequest, fmt.Sprintf("No such book with id %v", id))
			return
		}

		var updatedBook models.Book
		err = db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&updatedBook.Id, &updatedBook.Title, &updatedBook.Author, &updatedBook.Edition, &updatedBook.Genre, &updatedBook.IsAvailable, &updatedBook.Summary)
		if err != nil {
			log.Println(err)
			response.Error(c, http.StatusInternalServerError, "Error returning updated book from db")
			return
		}

		response.Success(c, http.StatusOK, "Book updated successfully", updatedBook)
	}
}

// @Summary     Delete a book in the library
// @Description Delete a single book using the id
// @Tags        books
// @Produce     json
// @Success     200 {object} response.BookSuccessResponse "Book deleted successfully"
// @Failure     500 {object} response.ErrorResponse "Server error"
// @Router      /api/v1/books/:id [delete]
func DeleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		result, err := db.Exec("DELETE FROM books WHERE id = ?", id)
		if err != nil {
			log.Println(err)
			response.Error(c, http.StatusInternalServerError, "could not delete book from db")
			return
		}

		rowsAffected, err := result.RowsAffected()
		if rowsAffected == 0 {
			log.Println(err)
			response.Error(c, http.StatusBadRequest, "No such book with id")
			return
		}

		response.Success(c, http.StatusOK, "Book deleted successfully")
	}
}
