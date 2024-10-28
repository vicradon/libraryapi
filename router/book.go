package router

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vicradon/internpulse/stage4/controller"
)

func Book(r *gin.Engine, ApiVersion string, db *sql.DB) *gin.Engine {

	bookUrls := r.Group(fmt.Sprintf("%s/books", ApiVersion))
	{
		bookUrls.POST("/", controller.CreateBook(db))
		bookUrls.GET("/", controller.GetBooks(db))
		bookUrls.GET("/:id", controller.GetBook(db))
		bookUrls.PUT("/:id", controller.UpdateBook(db))
		bookUrls.DELETE("/:id", controller.DeleteBook(db))
	}

	return r
}
