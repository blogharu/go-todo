package main

import (
	"go-todo/infrastructure/sqlite"
	"go-todo/user_interface/todo"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = sqlite.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	todo.AddApis(r)
	r.Run(":8000")
}
