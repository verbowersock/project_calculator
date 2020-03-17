package main

import (
	"calculator/board"
	"calculator/database"
	"calculator/finish"
	"calculator/handlers"
	"calculator/material"
	"calculator/products"
	"calculator/wood"
	_ "database/sql"
	"github.com/gin-gonic/gin"
)

var (
	db = database.InitDB("calculator.db")
	boardRepo = database.NewBoardRepo(db)
	boardService = board.NewBoardService(boardRepo)
	boardHandler = handlers.NewBoardHandler(boardService)
	finishRepo=database.NewFinishRepository(db)
	finishService= finish.NewFinishService(finishRepo)
	finishHandler = handlers.NewFinishHandler(finishService)
	materialRepo =database.NewMaterialRepository(db)
	materialService = material.NewMaterialService(materialRepo)
	materialHandler = handlers.NewMaterialHandler(materialService)
	productRepo=database.NewProductRepository(db)
	productService = products.NewProductService(productRepo)
	productHandler = handlers.NewProductHandler(productService)
	woodRepo =database.NewWoodRepository(db)
	woodService = wood.NewWoodService(woodRepo)
	woodHandler = handlers.NewWoodHandler(woodService)
)

func initializeRoutes(router *gin.RouterGroup) {
	initializeApiRoutes(router)
//	initializeHtmlRoutes(router, c)
}
func initializeApiRoutes(api *gin.RouterGroup){

	api.GET("/boards", boardHandler.Get)
	api.POST("/boards", boardHandler.Create)
	api.GET("/boards/:id", boardHandler.GetById)
	api.PUT("/boards/:id", boardHandler.Update)
	api.DELETE("/boards/:id", boardHandler.Delete)



	api.GET("/finishes", finishHandler.Get)
	api.POST("/finishes", finishHandler.Create)
	api.GET("/finishes/:id", finishHandler.GetById)
	api.PUT("/finishes/:id", finishHandler.Update)
	api.DELETE("/finishes/:id", finishHandler.Delete)



	api.GET("/materials", materialHandler.Get)
	api.POST("/materials", materialHandler.Create)
	api.GET("/materials/:id", materialHandler.GetById)
	api.PUT("/materials/:id/:price", materialHandler.Update)
	api.DELETE("/materials/:id", materialHandler.Delete)



	api.GET("/products", productHandler.Get)
	api.GET("/products/:id", productHandler.GetById)
	api.PUT("/products/:id", productHandler.UpdateProd)
	api.DELETE("/products/:id", productHandler.DeleteProd)
	api.DELETE("/products/:id/:mid", productHandler.DeleteMat)
	api.POST("/products", productHandler.Create)
	api.PATCH("/products/:id", productHandler.AddMat)
	api.PATCH("/products/:id/:mid", productHandler.UpdateMatQ)



	api.GET("/wood", woodHandler.Get)
	api.POST("/wood", woodHandler.Create)
	api.GET("/wood/:id", woodHandler.GetById)
	api.PUT("/wood/:id", woodHandler.Update)
	api.DELETE("/wood/:id", woodHandler.Delete)
}

/*func initializeHtmlRoutes(router *gin.RouterGroup, c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.GET("/", boardHandler.Get)
}
*/