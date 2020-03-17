package finish

type FinishService interface {
	CreateFinish(finish *Finish) (*Finish, error)
	FindFinishById(id int) (*Finish, error)
	FindAllFinishes() ([]*Finish, error)
	Update(fin *Finish) (*Finish, error)
	Delete (id int) error
}

type finishService struct {
	repo FinishRepository
}


func NewFinishService(repo FinishRepository) FinishService {
	return &finishService{
		repo,
	}
}

func (s *finishService) CreateFinish(finish *Finish) (*Finish, error) {
	return s.repo.Create(finish)
}


func (s *finishService) FindFinishById(id int) (*Finish, error){
	return s.repo.FindById(id)
}

func (s *finishService) FindAllFinishes() ([]*Finish, error) {
	return s.repo.FindAll()
}

func (s *finishService) Update(fin *Finish) (*Finish, error) {
	return s.repo.Update(fin)
}

func (s *finishService) Delete(id int) error {
	return s.repo.Delete(id)
}