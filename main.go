package main

import (
    "calculator/database"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    cors "github.com/itsjamie/gin-cors"
    _ "github.com/mattn/go-sqlite3"
    "time"
)
var router *gin.Engine


func main() {
    database.InitDB("calculator.db")
    router = gin.Default()


    api := router.Group("/api")

        //    html := router.Group("/"
        initializeRoutes(api)
        // initializeRoutes(html)

    router.Use(static.Serve("/", static.LocalFile("./frontend/client/build", true)))

    router.Use(cors.Middleware(cors.Config{
        Origins: "*",
        Methods:         "GET, PUT, POST, DELETE",
        RequestHeaders:  "Content-Type",
        ExposedHeaders:  "",
        MaxAge:          50 * time.Second,
        ValidateHeaders: false,
    }))



    router.Run() // listen and serve on 0.0.0.0:8080

}

