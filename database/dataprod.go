package database

import (
	"calculator/products"
	"database/sql"
	"errors"
	"fmt"
)

type productRepository struct {
	db *sql.DB
}

func (p productRepository) Create(prod *products.Product) (*products.Product, error) {
	err1 := db.QueryRow("SELECT name, finish, hours FROM product WHERE  name=? AND finish = ?", prod.Name, prod.Finish).Scan(&prod.Name, &prod.Finish, prod.Hours)
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO product(ID, name, finish, hours)  values (?, ?, ?, ?)")
		defer statement.Close()
		_, err := statement.Exec(nil, prod.Name, prod.Finish, prod.Hours)
		if err != nil {
			fmt.Printf("error: %s", err)
			return nil, errors.New(err.Error())
		}
		return prod, nil
	}
	return nil, errors.New("duplicate value")
}

func (p productRepository) FindById(id int) (*products.Product, error) {
	prod := new(products.Product)
	err := db.QueryRow("SELECT * FROM product WHERE id = ?", id).Scan(&prod.ID, &prod.Name, &prod.Finish, &prod.Hours)
	prod.Materials, _ = FindMats(prod.ID)
	fees := totals(prod)[0]
	total := totals(prod)[1]
	prod.Fees = fmt.Sprintf("%0.2f", fees)
	prod.Total = fmt.Sprintf("%0.2f", total)
	if err != nil {
		//fmt.Println(err)
		//log.Fatal(err)
		return nil, err
	}
	return prod, nil
}

func (p productRepository) FindAll() ([]*products.Product, error) {
	rows, err := p.db.Query("select product.id, product.name, product.finish, product.hours from product")
	defer rows.Close()
	prods := make([]*products.Product, 0)
	for rows.Next() {
		prod := new(products.Product)
		if err = rows.Scan(&prod.ID, &prod.Name, &prod.Finish, &prod.Hours); err != nil {
			fmt.Printf("error: %s", err)
			return nil, err
		}
		prod.Materials, _ = FindMats(prod.ID)
		prods = append(prods, prod)
		fees := totals(prod)[0]
		total := totals(prod)[1]
		prod.Fees = fmt.Sprintf("%0.2f", fees)
		prod.Total = fmt.Sprintf("%0.2f", total)
	}
	return prods, nil
}

func FindMats(id int) ([]*products.ProdMat, error) {
	rows, err := db.Query("SELECT materials.ID, board.board_size AS board, wood.wood AS wood, materials.price, product_composition.quantity from product left join product_composition on product.ID=product_composition.prodID left join materials on materials.ID=product_composition.material left join board on board.ID=materials.board left join wood on wood.ID=materials.wood where product.ID=? and materials.ID=product_composition.material", id)
	defer rows.Close()
	prodMats := make([]*products.ProdMat, 0)
	for rows.Next() {
		pm := new(products.ProdMat)
		if err = rows.Scan(&pm.ID, &pm.Board, &pm.Wood, &pm.Price, &pm.Quantity); err != nil {
			fmt.Printf("error: %s", err)
			return nil, err
		}
		prodMats = append(prodMats, pm)
		fmt.Println(prodMats)
	}
	return prodMats, nil
}

func totals(prod *products.Product) []float32 {
	var fees float32 = 0
	var matCost float32 = 0
	var finCost float32 = 0
	var total float32 = 0
	var totalNums []float32
	err := db.QueryRow("SELECT price FROM finishes left join product on product.finish = finishes.ID WHERE product.ID= ?", prod.ID).Scan(&finCost)
	if err != nil {
		return nil
	}
	for _, i := range prod.Materials {
		matCost += i.Price * float32(i.Quantity)
	}
	fees = (matCost + finCost + (prod.Hours * 25)) * 0.1
	totalNums = append(totalNums, fees)
	total = matCost + finCost + (prod.Hours * 25) + fees
	totalNums = append(totalNums, total)
	return totalNums
}

func (p productRepository) Update(prod *products.Product) (*products.Product, error) {
	statement, _ := db.Prepare("update product set name=?, finish=?, hours=? where id=?")
	defer statement.Close()
	_, err := statement.Exec(prod.Name, prod.Finish, prod.Hours, prod.ID)
	if err != nil {
		return nil, err
	}
	//i, _ := result.RowsAffected()
	return prod, nil
}

func (p productRepository) AddMat(pId int, mId int, qty int) (*products.Product, error) {
	err1 := db.QueryRow("SELECT ID, material, quantity FROM product_composition WHERE  material=?", mId).Scan()
	if err1 == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO product_composition(ID, prodID, material, quantity)  values (?, ?, ?, ?)")
		defer statement.Close()
		_, err := statement.Exec(nil, pId, mId, qty)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		prod, err := p.FindById(pId)
		return prod, nil
	}
	return nil, errors.New("duplicate value")
}

func (p productRepository) UpdateMatQ(pId int, mId int, qty int) (*products.Product, error) {
	statement, _ := db.Prepare("UPDATE product_composition SET quantity=? WHERE material=? AND prodID=?")
	defer statement.Close()
	_, err := statement.Exec(qty, mId, pId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	prod, err := p.FindById(pId)
	return prod, nil
}

func (p productRepository) DeleteMat(pId int, mId int) error {
	statement1, _ := db.Prepare("delete from product_composition where material=? and prodID=?")
	defer statement1.Close()
	_, err := statement1.Exec(mId, pId)
	if err != nil {
		return err
	}
	return nil
}

func (p productRepository) Delete(id int) error {
	statement1, _ := db.Prepare("delete from product_composition where prodID=?")
	defer statement1.Close()
	_, err1 := statement1.Exec(id)
	if err1 != nil {
		return err1
	}
	statement2, _ := db.Prepare("delete from product where id=?")
	defer statement2.Close()
	_, err := statement2.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func NewProductRepository(db *sql.DB) products.ProductRepository {
	return &productRepository{
		db,
	}
}
