package task

type Repository interface {
	GetAll() ([]Info, error)
	GetByID(id int) (Info, error)
	Create(taskInfo Info) (Info, error)
	Update(taskInfo Info) (Info, error)
	Delete(id int) error
}
