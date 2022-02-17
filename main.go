package main

import (
	"github.com/emnaruto07/beginner-javascript/controllers"
	"github.com/emnaruto07/beginner-javascript/db_client"
	"github.com/gin-gonic/gin"
)

func main() {
	db_client.DBConnection()

	r := gin.Default()
	r.GET("/actors", controllers.FetchPost)
	r.GET("/actor/:id", controllers.GetPost)

	if err := r.Run(":5000"); err != nil {
		panic(err)
	}

}
