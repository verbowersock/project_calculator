package wood

type WoodService interface {
	CreateWood(wood *Wood) (*Wood, error)
	FindWoodById(id int) (*Wood, error)
	FindAllWoods() ([]*Wood, error)
	Update(wd *Wood) (*Wood, error)
	Delete (id int) error
}

type woodService struct {
	repo WoodRepository
}


func NewWoodService(repo WoodRepository) WoodService {
	return &woodService{
		repo,
	}
}

func (s *woodService) CreateWood(wood *Wood) (*Wood, error) {
	return s.repo.Create(wood)
}


func (s *woodService) FindWoodById(id int) (*Wood, error){
	return s.repo.FindById(id)
}

func (s *woodService) FindAllWoods() ([]*Wood, error) {
	return s.repo.FindAll()
}

func (s *woodService) Update(wd *Wood) (*Wood, error) {
	return s.repo.Update(wd)
}

func (s *woodService) Delete(id int) error {
	return s.repo.Delete(id)
}
