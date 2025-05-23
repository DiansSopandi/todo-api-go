package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{}
var nextID = 1

// GetTodo godoc
// @Summary Get all todos
// @Description Get a list of all todo
// @Tags todo
// @Produce json
// @Success 200 {array} todo.Todo
// @Router /todo [get]
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "fetch all todo", "data": todos})
}

// @Summary Get todo by id
// @Tags todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} todo.Todo
// @Router /todo/{id} [get]
func Get(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error(), "data": nil})
		return
	}

	for _, eachTodo := range todos {
		if eachTodo.ID == id {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "fetch todo", "data": eachTodo})
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Todo not found", "data": nil})
}

// @Summary Create todo
// @Tags todo
// @Produce json
// @Param todo body todo.Todo true "Todo object"
// @Success 201 {object} todo.Todo
// @Router /todo [post]
func Create(c *gin.Context) {
	var input struct {
		Task string `json:"task"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error(), "data": nil})
		return
	}

	newTodo := Todo{ID: nextID, Task: input.Task}
	nextID++
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Create todo", "data": newTodo})
}

// @Summary Update todo
// @Tags todo
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body todo.Todo true "Todo object"
// @Success 200 {object} todo.Todo
// @Router /todo/{id} [put]
func Update(c *gin.Context) {
	var input struct {
		Task string `json:"task"`
	}
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid ID"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error(), "data": nil})
		return
	}

	for i, eachTodo := range todos {
		if eachTodo.ID == id {
			todos[i].Task = input.Task
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Update todo", "data": todos})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Todo not found", "data": nil})
}

// @Summary Delete todo
// @Tags todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} todo.Todo
// @Router /todo/{id} [delete]
func Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error(), "data": nil})
		return
	}

	for i, eachTodo := range todos {
		if eachTodo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete todo", "data": todos})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Todo not found", "data": nil})

}
