package response

import (
	"github.com/gin-gonic/gin"
	"github.com/vicradon/internpulse/stage4/models"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, message string, data ...interface{}) {
	response := Response{
		Status:  "success",
		Message: message,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	c.JSON(code, response)
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
	})
}

// ErrorREsponse represents the structure of the error response
// @Description Error response model
// @ID ErrorResponse
type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"An error occured"`
}

// BookSuccessResponse represents a successful response of a single book
type BookSuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Data fetched successfully"`
	Data    models.Book `json:"data,omitempty"`
}

// BooksSuccessREsponse represents a successful response of multiple books
type BooksSuccessResponse struct {
	Status  string        `json:"status" example:"success"`
	Message string        `json:"message" example:"Data fetched successfully"`
	Data    []models.Book `json:"data,omitempty"`
}
