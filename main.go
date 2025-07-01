package main

import (
	//"fmt"
	//"golang.org/x/net/html"
	//"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID uint `json:"id"`
	Item string `json:"item"`
	Description string `json:"desc"`
	Completed bool `json:"status"`
}
var todo_list = []todo{
	{ID: 1, Item: "Pasutit EVAK karti", Description: "Ieej eveseliba.gov.lv, vajag lidz Portugalei paspet", Completed: false},
	{ID: 2, Item: "MethodsX paper", Description: "Izej cauri tam kas ir izdarits, velreiz parlasi example, sac rakstit", Completed: false},
	{ID: 3, Item: "Facebook marketplace salikt mebeles", Description: "Var ari pielikt taja wpp grupa, pajauta Tomas", Completed: false},
}

func getTodoList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todo_list)
}

// Goal: make a website API that works as a to-do list
func main() {
	router := gin.Default()
	router.GET("/todo_list", getTodoList)
	router.Run("localhost:9090")
	
}