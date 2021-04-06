package database

import (
	"calculator/material"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type materialRepository struct {
	db *sql.DB
}

func (m materialRepository) Create(material *material.Material) (*material.Material, error) {
	err1 := db.QueryRow("SELECT * FROM materials WHERE wood = ? AND board = ? AND price = ?", material.Wood, material.Board, material.Price).Scan(&material.ID, &material.Wood, &material.Board, &material.Price)
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO materials (ID, wood, board, price)  values (?, ?, ?, ?)")
		defer statement.Close()
		_, err := statement.Exec(nil, material.Wood, material.Board, material.Price)
		if err != nil {
			fmt.Printf("error: %s", err)
			return nil, errors.New(err.Error())
		}
		return material, nil
	}
	return nil, errors.New("duplicate value")
}

func (m materialRepository) Delete(id int) error {
	statement, _ := db.Prepare("delete from materials where id=?")
	defer statement.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return nil
}

func (m materialRepository) UpdatePrice(id int, price float32) (*material.PriceUpdateMaterial, error) {
	newMat := new(material.PriceUpdateMaterial)
	stmt, err := db.Prepare("UPDATE materials SET price=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(price, id)
	if err != nil {
		log.Fatal(err)
	}
	return newMat, nil
}

func (m materialRepository) FindById(id int) (*material.Material, error) {
	mat := new(material.Material)
	err := db.QueryRow("SELECT * FROM materials WHERE id = ?", id).Scan(&mat.ID, &mat.Wood, &mat.Board, &mat.Price)
	if err != nil {
		//fmt.Println(err)
		//log.Fatal(err)
		return nil, err
	}
	return mat, nil
}

func (m materialRepository) FindAll() ([]*material.Material, error) {
	rows, err := m.db.Query("SELECT materials.ID, wood.wood AS wood, board.board_size AS board, materials.price AS price FROM materials INNER JOIN board ON  materials.board=board.ID INNER JOIN wood ON materials.wood=wood.ID")
	defer rows.Close()
	materials := make([]*material.Material, 0)
	for rows.Next() {
		material := new(material.Material)
		if err = rows.Scan(&material.ID, &material.Wood, &material.Board, &material.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}

func NewMaterialRepository(db *sql.DB) material.MaterialRepository {
	return &materialRepository{
		db,
	}
}
