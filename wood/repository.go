package wood

type WoodRepository interface {
	Create(wd *Wood) (*Wood, error)
	FindById(id int) (*Wood, error)
	FindAll() ([]*Wood, error)
	Update(wd *Wood) (*Wood, error)
	Delete (id int) error
}
