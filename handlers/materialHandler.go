package handlers

import (
	"calculator/material"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MaterialHandler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type materialHandler struct {
	materialsService material.MaterialService
}

func (m materialHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := m.materialsService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}

func (m materialHandler) Update(c *gin.Context) {
	var a *material.PriceUpdateMaterial
	_ = c.BindJSON(&a)
	id, _ := strconv.Atoi(c.Param("id"))
	a.ID = id
	price, _ := strconv.ParseFloat(c.Param("price"), 32)
	a.Price = float32(price)
	_, err := m.materialsService.UpdatePrice(id, float32(price))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {

		c.JSON(http.StatusOK, a)
	}
}


func (m materialHandler) Get(c *gin.Context) {
	Materials, err := m.materialsService.FindAllMaterials()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, Materials)
}



func (m materialHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Material, err := m.materialsService.FindMaterialById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, Material)

	}
}

func (m materialHandler) Create(c *gin.Context) {
		var mat *material.Material
		error:= c.BindJSON(&mat)
		fmt.Println(error)
		_, err := m.materialsService.CreateMaterial(mat)
		if err != nil {
			fmt.Printf("Error: %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, mat)
		}
	}

func NewMaterialHandler(materialService material.MaterialService) MaterialHandler {
	return &materialHandler{
		materialService,
	}
}



