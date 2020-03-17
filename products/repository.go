package products

type ProductRepository interface {
	Create(*Product) (*Product, error)
	FindById(int) (*Product, error)
	FindAll() ([]*Product, error)
	Update(prod *Product)(*Product, error)
	AddMat(pId int, mId int, qty int)(*Product, error)
	UpdateMatQ(pId int, mId int, qty int)(*Product, error)
	DeleteMat(pId int, mId int) error
	Delete(int)error

}
