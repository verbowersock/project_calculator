package handlers

import (
	"calculator/products"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateProd(c *gin.Context)
	AddMat(c *gin.Context)
	DeleteMat(c *gin.Context)
	DeleteProd(c *gin.Context)
	UpdateMatQ(c *gin.Context)
}

type productHandler struct {
	productService products.ProductService
}

func (f productHandler) Get(c *gin.Context) {
	Products, err := f.productService.FindAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, Products)
}

func (f productHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Board, err := f.productService.FindProductById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, Board)
	}
}

func (f productHandler) Create(c *gin.Context) {
	var prd *products.Product
	_ = c.BindJSON(&prd)
	_, err := f.productService.CreateProduct(prd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusCreated, prd)
	}
}

func (f productHandler) UpdateProd(c *gin.Context) {
	var a *products.Product
	_ = c.BindJSON(&a)
	id, _ := strconv.Atoi(c.Param("id"))
	a.ID = id
	_, err := f.productService.UpdateProd(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {

		c.JSON(http.StatusOK, a)
	}
}

func (f productHandler) AddMat(c *gin.Context) {

	pId, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(pId)
	type a struct {
		MatId int `json:"mid" binding:"required"`
		Qty   int `json:"qty"`
	}
	var b *a
	err1 := c.BindJSON(&b);
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
	}

	//var prd *products.Product
	prd, err := f.productService.AddMat(pId, b.MatId, b.Qty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, prd)
	}
}

func (f productHandler) UpdateMatQ(c *gin.Context) {

	pId, _ := strconv.Atoi(c.Param("id"))
	mId, _ :=strconv.Atoi(c.Param("mid"))
	fmt.Println(pId)
	type a struct {
		Qty   int `json:"qty"`
	}
	var b *a
	err1 := c.BindJSON(&b);
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
	}

	//var prd *products.Product
	prd, err := f.productService.UpdateMatQ(pId, mId, b.Qty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, prd)
	}
}


func (f productHandler) DeleteMat(c *gin.Context) {
	mId, _ := strconv.Atoi(c.Param("mid"))
	pId, _ := strconv.Atoi(c.Param("id"))
	err := f.productService.DeleteMat(pId, mId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func (f productHandler) DeleteProd(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := f.productService.DeleteProd(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func NewProductHandler(productService products.ProductService) ProductHandler {
	return &productHandler{
		productService,
	}
}

