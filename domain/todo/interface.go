package todo

type TodoInterface interface {
	GetList() ([]Todo, error)
	Get(todoID int) Todo
	Update(todo *Todo)
	Create(todo *Todo)
	Delete(todoID int)
}
