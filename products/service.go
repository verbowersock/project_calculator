package products

type ProductService interface {
	CreateProduct(product *Product) (*Product, error)
	FindProductById(id int) (*Product, error)
	FindAllProducts() ([]*Product, error)
	UpdateProd(prod *Product)(*Product, error)
	AddMat(pId int, mId int, qty int)(*Product, error)
	UpdateMatQ(pId int, mId int, qty int)(*Product, error)
	DeleteMat(pId int, mId int) error
	DeleteProd(int)error


}

type productService struct {
	repo ProductRepository
}

func (p productService) CreateProduct(product *Product) (*Product, error) {
	return p.repo.Create(product)
}

func (p productService) FindProductById(id int) (*Product, error) {
	return p.repo.FindById(id)
}

func (p productService) FindAllProducts() ([]*Product, error) {
	return p.repo.FindAll()
}

func (p productService) UpdateProd(prod *Product) (*Product, error) {
	return p.repo.Update(prod)
}

func (p productService) AddMat(pId int, mId int, qty int) (*Product, error) {
	return p.repo.AddMat(pId, mId, qty)
}

func (p productService) UpdateMatQ(pId int, mId int, qty int) (*Product, error) {
	return p.repo.UpdateMatQ(pId, mId, qty)
}

func (p productService) DeleteMat(pId int, mId int) error {
	return p.repo.DeleteMat(pId, mId)
}

func (p productService) DeleteProd(id int) error {
	return p.repo.Delete(id)
}


func NewProductService(repo ProductRepository) ProductService {
	return &productService{
		repo,
	}
}