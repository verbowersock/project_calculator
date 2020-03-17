package database

import (
	"calculator/finish"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type finishRepository struct {
	db *sql.DB
}

func (f finishRepository) Update(fin *finish.Finish) (*finish.Finish, error) {
	statement, _ := db.Prepare("update finishes set fintype=?, stain=?, price=? where id=?")
	defer statement.Close()
	_, err := statement.Exec(fin.Finish, fin.Stain, fin.Price, fin.ID)
	if err != nil {
		fmt.Sprintf("error: %s", err)
		return nil, err
	}
	//i, _ := result.RowsAffected()
	return fin, nil
}

func (f finishRepository) Delete(id int) error {
	statement, _ := db.Prepare("delete from finishes where id=?")
	defer statement.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return nil
}

func (f finishRepository) Create(fin *finish.Finish) (*finish.Finish, error) {
	err1 := db.QueryRow("SELECT * FROM finishes WHERE fintype = ? AND stain = ?", fin.Finish, fin.Stain).Scan(&fin.ID, &fin.Finish, &fin.Stain)
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO finishes (ID, fintype, stain)  values (?, ?, ?)")
		defer statement.Close()
		_, err := statement.Exec(nil, fin.Finish, fin.Stain)
		if err != nil {
			fmt.Printf("error: %s", err)
			return nil, errors.New(err.Error())
		}
		return fin, nil
	}
	return nil, errors.New("duplicate value")
}

func (f finishRepository) FindById(id int) (*finish.Finish, error) {
	fin := new(finish.Finish)
	err := db.QueryRow("SELECT * FROM finishes WHERE id = ?", id).Scan(&fin.ID, &fin.Finish, &fin.Stain)
	//fmt.Println(fin)
	if err != nil {
		//fmt.Println(err)
		//log.Fatal(err)
		return nil, err
	}
	return fin, nil
}

func (f finishRepository) FindAll() ([]*finish.Finish, error) {

	rows, err := f.db.Query("SELECT * FROM finishes")

	defer rows.Close()
	finishes :=make([]*finish.Finish, 0)
	for rows.Next() {
		finish := new(finish.Finish)
		if err = rows.Scan(&finish.ID, &finish.Finish, &finish.Stain); err != nil {
			log.Fatal()
			return nil, err
		}
		finishes = append(finishes, finish)
	}
	return finishes, nil
}


func NewFinishRepository(db *sql.DB) finish.FinishRepository{
	return &finishRepository{
		db,
	}
}

