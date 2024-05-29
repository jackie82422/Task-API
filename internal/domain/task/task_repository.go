package task

type Repository interface {
	GetAll() ([]Task, error)
	GetByID(id int) (Task, error)
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(id int) error
}
