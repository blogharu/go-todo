package sqlite

import (
	"go-todo/domain/todo"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB_FILE")), &gorm.Config{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&todo.Todo{})
	if err != nil {
		return
	}
	Todo = TodoRepository{DB: DB}
	return
}
