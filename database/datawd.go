package database

import (
	"calculator/wood"
	"database/sql"
	"errors"
	"log"
)

type woodRepository struct {
	db *sql.DB
}

func (w woodRepository) Update(wd *wood.Wood) (*wood.Wood, error) {
	statement, _ := db.Prepare("update wood set wood=? where id=?")
	defer statement.Close()
	_, err := statement.Exec(wd.Type, wd.ID)
	if err != nil {
		return nil, err
	}
	//i, _ := result.RowsAffected()
	return wd, nil
}

func (w woodRepository) Delete(id int) error {
	statement, _ := db.Prepare("delete from wood where id=?")
	defer statement.Close()
	_, err := statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (w woodRepository) Create(wd *wood.Wood) (*wood.Wood, error) {
	err1 := db.QueryRow("SELECT * FROM wood WHERE wood = ?", wd.Type).Scan(&wd.ID, &wd.Type)
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO wood (ID, wood)  values (?, ?)")
		defer statement.Close()
		_, err := statement.Exec(nil, wd.Type)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return wd, nil
	}
	return nil, errors.New("duplicate value")
}

func (w woodRepository) FindById(id int) (*wood.Wood, error) {
	wd := new(wood.Wood)
	err := db.QueryRow("SELECT * FROM wood WHERE id = ?", id).Scan(&wd.ID, &wd.Type)
	if err != nil {
		return nil, err
	}
	return wd, nil
}

func (w woodRepository) FindAll() ([]*wood.Wood, error) {
	rows, err := w.db.Query("SELECT * FROM wood")
	defer rows.Close()
	woods :=make([]*wood.Wood, 0)
	for rows.Next() {
		wood := new(wood.Wood)
		if err = rows.Scan(&wood.ID, &wood.Type); err != nil {
			log.Fatal()
			return nil, err
		}
		woods = append(woods, wood)
	}
	return woods, nil
}

func NewWoodRepository(db *sql.DB) wood.WoodRepository{
	return &woodRepository{
		db,
	}
}

