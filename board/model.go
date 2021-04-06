package board

type Board struct {
	ID   int    `json:"id"`
	Size string `json:"board_size" binding:"required"`
}
