package main

import (
	"github.com/emnaruto07/beginner-javascript/controllers"
	"github.com/emnaruto07/beginner-javascript/db_client"
	"github.com/gin-gonic/gin"
)

func main() {
	db_client.DBConnection()
	defer db_client.DBClient.Close()

	r := gin.Default()
	r.GET("/actors", controllers.FetchRecord)
	r.GET("/actor/:id", controllers.GetRecord)

	if err := r.Run(":5000"); err != nil {
		panic(err)
	}

}
