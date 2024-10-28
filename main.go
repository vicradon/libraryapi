// @title Library API
// @version 1.0
// @description This API returns content of a library including books
// @host {host}
// @BasePath /
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vicradon/internpulse/stage4/database"
	"github.com/vicradon/internpulse/stage4/docs"
	_ "github.com/vicradon/internpulse/stage4/docs"
	"github.com/vicradon/internpulse/stage4/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error reading .env file %v", err)
	}

	databaseFile := os.Getenv("DATABASE_FILE")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if host == "" {
		host = fmt.Sprintf("localhost:%s", port)
	}

	docs.SwaggerInfo.Host = host

	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.InitDB(db)

	r := router.Setup(db)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Printf("Server is running on port %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
