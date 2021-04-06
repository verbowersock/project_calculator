package main

import (
	"calculator/database"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	_ "github.com/mattn/go-sqlite3"
)

var router *gin.Engine

func main() {
	database.InitDB("calculator.db")

	router = gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		ValidateHeaders: false,
	}))

	api := router.Group("/api")

	//    html := router.Group("/"
	initializeRoutes(api)
	// initializeRoutes(html)

	router.Use(static.Serve("/", static.LocalFile("./frontend/client/build", true)))

	router.Run() // listen and serve on 0.0.0.0:8080

}
