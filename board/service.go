package board

import "fmt"

type BoardService interface {
	CreateBoard(brd *Board) (*Board, error)
	FindBoardById(id int) (*Board, error)
	FindAllBoards() ([]*Board, error)
	Update( brd *Board) (*Board, error)
	Delete (id int) error
}

type boardService struct {
	repo BoardRepository
}

func NewBoardService(repo BoardRepository) BoardService {
	return &boardService{
		repo,
	}
}

func (s *boardService) CreateBoard(board *Board) (*Board, error) {
	fmt.Println("going into service")
	return s.repo.Create(board)
}


func (s *boardService) FindBoardById(id int) (*Board, error){
	return s.repo.FindById(id)
}

func (s *boardService) FindAllBoards() ([]*Board, error) {
	return s.repo.FindAll()
}

func (s *boardService) Update(brd *Board) (*Board, error) {
	return s.repo.Update(brd)
}

func (s *boardService) Delete(id int)  error {
	return s.repo.Delete(id)
}