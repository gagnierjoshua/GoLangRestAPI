package main

//commands ran: go mod init & go get github.com/gin-gonic/gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Task 1", Completed: false},
	{ID: "2", Item: "Task 2", Completed: false},
	{ID: "3", Item: "Task 3", Completed: false},
}

func getTodos(c *gin.Context) { //gets the tasks from server and displays them *gin is a web framework
	c.IndentedJSON(http.StatusOK, todos) //returns the tasks in json format with indentation
}

func addTodo(context *gin.Context) { //adds a new task to the server
	var newTodo todo //creates a new task

	if err := context.BindJSON(&newTodo); err != nil { //binds the new task to the server
		return //returns the new task
	}

	todos = append(todos, newTodo) //adds the new task to the server

	context.IndentedJSON(http.StatusCreated, newTodo) //returns the new task in json format with indentation
}

func main() {
	router := gin.Default()        //creates a router/server
	router.GET("/todos", getTodos) //gets the data from the server
	router.POST("/todos", addTodo) //adds the data to the server
	router.Run("localhost:9090")   //runs the server on port 9090
}

//go run main.go to make it run
