package sqlite

import (
	"go-todo/domain/todo"
	"time"

	"gorm.io/gorm"
)

var Todo TodoRepository

type TodoRepository struct {
	DB *gorm.DB
}

func (t *TodoRepository) GetList() (todos []todo.Todo) {
	rows, err := t.DB.Raw("select * from todos").Rows()
	defer rows.Close()
	if err != nil {
		return
	}
	var todo_entity todo.Todo
	for rows.Next() {
		t.DB.ScanRows(rows, &todo_entity)
		todos = append(todos, todo_entity)
	}
	return
}

func (t *TodoRepository) Get(todoID int) (todo todo.Todo) {
	t.DB.Raw("select * from todos where id = ?", todoID).Scan(&todo)
	return
}

func (t *TodoRepository) Update(todo *todo.Todo) {
	now := time.Now()
	t.DB.Exec(
		"update todos set title = ?, category = ?, description = ?, updated_at = ? where id = ?",
		todo.Title,
		todo.Category,
		todo.Description,
		now,
		todo.ID,
	)
}

func (t *TodoRepository) Create(todo *todo.Todo) {
	now := time.Now()
	t.DB.Exec(
		"insert into todos (title, category, description, created_at, updated_at) values (?, ?, ?, ?, ?)",
		todo.Title,
		todo.Category,
		todo.Description,
		now,
		now,
	)
}

func (t *TodoRepository) Delete(todoID int) {
	t.DB.Exec(
		"delete from todos where id = ?",
		todoID,
	)
}
