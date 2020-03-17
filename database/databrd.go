package database

import (
	"calculator/board"
	"database/sql"
	"errors"
	"fmt"
)

type boardRepo struct {
	db *sql.DB
}

func (b boardRepo) Create(brd *board.Board) (*board.Board, error) {
	err1 := db.QueryRow("SELECT * FROM board WHERE board_size = ?", brd.Size).Scan(&brd.ID, &brd.Size)
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO board (ID, board_size)  values (?, ?)")
		defer statement.Close()
		_, err := statement.Exec(nil, brd.Size)
		if err != nil {
			fmt.Printf("error: %s", err)
			return nil, errors.New(err.Error())
		}
		return brd, nil
	}
	return nil, errors.New("duplicate value")
}
func (b boardRepo) FindById(id int) (*board.Board, error) {
	brd := new(board.Board)
	err := db.QueryRow("SELECT * FROM board WHERE id = ?", id).Scan(&brd.ID, &brd.Size)
	fmt.Println(brd)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
		return nil, err
	}
	return brd, nil
}

func (b boardRepo) FindAll() ([]*board.Board, error) {
		rows, err := b.db.Query("SELECT * FROM board")
		defer rows.Close()
		brds :=make([]*board.Board, 0)
		for rows.Next() {
			board := new(board.Board)
			if err = rows.Scan(&board.ID, &board.Size);
				err != nil {
				return nil, err
			}
	brds = append(brds, board)
		}

		return brds, nil
}

func (b boardRepo) Update(brd *board.Board) (*board.Board, error) {
	err1 := db.QueryRow("SELECT * FROM board WHERE board_size = ?", brd.Size).Scan(&brd.ID, &brd.Size)
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("update board set board_size=? where id=?")
		defer statement.Close()
		_, err := statement.Exec(brd.Size, brd.ID)
		if err != nil {
		//fmt.Sprintf("error: %s", err)
			return nil, errors.New(err.Error())
	}
	//i, _ := result.RowsAffected()
	return brd, nil
}
return nil, errors.New("duplicate value")
}

func (b boardRepo) Delete(id int) (error) {
	statement, _ := db.Prepare("delete from board where id=?")
	defer statement.Close()
	result, err := statement.Exec(id)

	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)

	return nil
}
func NewBoardRepo(db *sql.DB) board.BoardRepository {
	return &boardRepo{
		db,
	}
}

