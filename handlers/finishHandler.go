package handlers

import (
	"calculator/finish"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FinishHandler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type finishHandler struct {
	finishService finish.FinishService
}

func (f finishHandler) Update(c *gin.Context) {
	var a *finish.Finish
	_ = c.BindJSON(&a)
	id, _ := strconv.Atoi(c.Param("id"))
	a.ID = id
	_, err := f.finishService.Update(a)
	if err != nil {
		fmt.Sprintf("error: %s", err )
		c.JSON(http.StatusBadRequest, err)
	} else {

		c.JSON(http.StatusOK, a)
	}
}

func (f finishHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := f.finishService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}

func (f finishHandler) Get(c *gin.Context){

		Finishes, err := f.finishService.FindAllFinishes()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, Finishes)

}


func (f finishHandler) GetById(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))
		Finish, err := f.finishService.FindFinishById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, Finish)


	}
}

func (f finishHandler) Create(c *gin.Context){

	var fin *finish.Finish
	_ = c.BindJSON(&fin)
	fmt.Println(fin)
	_, err := f.finishService.CreateFinish(fin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusCreated, fin)

	}
	}


func NewFinishHandler(finishService finish.FinishService) FinishHandler {
	return &finishHandler{
		finishService,
	}
}


