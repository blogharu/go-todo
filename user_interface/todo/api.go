package todo

import (
	"fmt"
	"go-todo/domain/todo"
	"go-todo/infrastructure/sqlite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddApis(engine *gin.Engine) {
	todo := engine.Group("/todo")
	{
		todo.GET("/", getTodoList)
		todo.POST("/", createTodo)
		todo.GET("/:todoID", getTodo)
		todo.POST("/:todoID", updateTodo)
		todo.POST("/delete/:todoID", deleteTodo)
	}
}

func getTodoList(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Todos": sqlite.Todo.GetList(),
	})
}

func createTodo(c *gin.Context) {
	var createRequest CreateRequest
	if c.ShouldBind(&createRequest) == nil && len(createRequest.Title) > 0 {
		sqlite.Todo.Create(
			&todo.Todo{
				Title:       createRequest.Title,
				Category:    createRequest.Category,
				Description: createRequest.Description,
			},
		)
		c.Redirect(http.StatusFound, "/todo")
	} else {
		c.HTML(http.StatusBadRequest, "bad_request.html", nil)
	}
}

func getTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "bad_request.html", nil)
	}
	if todoEntity := sqlite.Todo.Get(todoID); todoEntity.ID == 0 {
		c.HTML(http.StatusNoContent, "no_content.html", nil)
	} else {
		c.HTML(
			http.StatusOK,
			"detail.html",
			gin.H{"TodoEntity": todoEntity},
		)
	}
}

func updateTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "bad_request.html", nil)
	}
	if todoEntity := sqlite.Todo.Get(todoID); todoEntity.ID == 0 {
		c.HTML(http.StatusNoContent, "no_content.html", nil)
	} else {
		var updateRequest UpdateRequest
		if c.ShouldBind(&updateRequest) == nil && len(updateRequest.Title) > 0 {
			todoEntity.Title = updateRequest.Title
			todoEntity.Category = updateRequest.Category
			todoEntity.Description = updateRequest.Description
			sqlite.Todo.Update(&todoEntity)
			c.Redirect(http.StatusFound, fmt.Sprintf("/todo/%d", todoID))
		} else {
			c.HTML(http.StatusBadRequest, "bad_request.html", nil)
		}
	}
}

func deleteTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "bad_request.html", nil)
	}
	if sqlite.Todo.Get(todoID).ID == 0 {
		c.HTML(http.StatusNoContent, "no_content.html", nil)
	}
	sqlite.Todo.Delete(todoID)
	c.Redirect(http.StatusFound, "/todo")
}
