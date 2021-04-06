package handlers

import (
	"calculator/board"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardHandler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type boardHandler struct {
	boardService board.BoardService
}

func (b boardHandler) Update(c *gin.Context) {
	var a *board.Board
	_ = c.BindJSON(&a)
	id, _ := strconv.Atoi(c.Param("id"))
	a.ID = id
	_, err := b.boardService.Update(a)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {

		c.JSON(http.StatusOK, a)
	}
}

func (b boardHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := b.boardService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}

func (b boardHandler) Get(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	Boards, err := b.boardService.FindAllBoards()
	fmt.Println(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, Boards)
	}
}
func (b boardHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Board, err := b.boardService.FindBoardById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, Board)

	}
}

func (b boardHandler) Create(c *gin.Context) {
	fmt.Println("going into handler")
	var brd *board.Board
	if err := c.BindJSON(&brd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)

	} else {
		brd, err := b.boardService.CreateBoard(brd)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, brd)
		}
	}
}

func NewBoardHandler(boardService board.BoardService) BoardHandler {
	return &boardHandler{
		boardService,
	}
}
