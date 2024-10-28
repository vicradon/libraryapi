package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Setup(db *sql.DB) *gin.Engine {
	r := gin.Default()

	API_VERSION := "api/v1"

	Book(r, API_VERSION, db)

	return r
}
