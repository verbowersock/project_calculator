package board

type BoardRepository interface {
	Create(brd *Board) (*Board,error)
	FindById(id int) (*Board, error)
	FindAll() ([]*Board, error)
	Update( brd *Board) (*Board, error)
	Delete (id int) error
}


