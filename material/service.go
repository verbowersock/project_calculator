package material

type MaterialService interface {
	CreateMaterial(material *Material) (*Material, error)
	FindMaterialById(id int) (*Material, error)
	FindAllMaterials() ([]*Material, error)
	UpdatePrice(id int, price float32) (*PriceUpdateMaterial, error)
	Delete (id int)  error

}

type materialService struct {
	repo MaterialRepository
}

func NewMaterialService(repo MaterialRepository) MaterialService {
	return &materialService{
		repo,
	}
}

func (s *materialService) CreateMaterial(material *Material) (*Material, error) {
	return s.repo.Create(material)
}

func (s *materialService) FindMaterialById(id int) (*Material, error){
	return s.repo.FindById(id)
}

func (s *materialService) FindAllMaterials() ([]*Material, error) {
	return s.repo.FindAll()
}

func (s *materialService) UpdatePrice(id int, price float32) (*PriceUpdateMaterial, error) {
	return s.repo.UpdatePrice(id, price)
}
func (s *materialService) Delete(id int) error {
	return s.repo.Delete(id)
}