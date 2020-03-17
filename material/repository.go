package material


type MaterialRepository interface {
	Create(material *Material) (*Material, error)
	UpdatePrice(id int, price float32)(*PriceUpdateMaterial, error)
	FindById(id int) (*Material, error)
	FindAll() ([]*Material, error)
	Delete (id int) error
}