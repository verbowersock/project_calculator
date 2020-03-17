package handlers

import (
	"calculator/wood"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WoodHandler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type woodHandler struct {
	woodService wood.WoodService
}

func (f woodHandler) Update(c *gin.Context) {
	var a *wood.Wood
	_ = c.BindJSON(&a)
	id, _ := strconv.Atoi(c.Param("id"))
	a.ID = id
	_, err := f.woodService.Update(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {

		c.JSON(http.StatusOK, a)
	}
}

func (f woodHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := f.woodService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}

func (f woodHandler) Get(c *gin.Context) {
	Woods, err := f.woodService.FindAllWoods()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, Woods)

}

func (f woodHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Wood, err := f.woodService.FindWoodById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, Wood)

	}
}

func (f woodHandler) Create(c *gin.Context) {
	var wd *wood.Wood
	_ = c.BindJSON(&wd)
	fmt.Println(wd)
	_, err := f.woodService.CreateWood(wd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusCreated, wd)
	}
}


func NewWoodHandler(woodService wood.WoodService) WoodHandler {
	return &woodHandler{
		woodService,
	}
}

