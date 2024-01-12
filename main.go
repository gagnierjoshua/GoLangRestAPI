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
	c.IndentedJSON(http.StatusOK, todos) //returns the tasks in json format
}

func main() {
	router := gin.Default()        //creates a router/server
	router.GET("/todos", getTodos) //gets the data from the server
	router.Run("localhost:9090")   //runs the server on port 9090
}

//go run main.go to make it run
