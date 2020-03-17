package finish

type FinishRepository interface {
	Create(*Finish) (*Finish, error)
	FindById(id int) (*Finish, error)
	FindAll() ([]*Finish, error)
	Update( fin *Finish) (*Finish, error)
	Delete (id int) error
}