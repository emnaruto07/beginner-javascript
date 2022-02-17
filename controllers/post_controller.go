package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/emnaruto07/beginner-javascript/db_client"
	"github.com/gin-gonic/gin"
)

type Actor struct {
	Actor_id    int       `json:"id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Last_update time.Time `json:"last_update"`
	// last_modified
}

func FetchRecord(c *gin.Context) {
	var posts []Actor

	rows, err := db_client.DBClient.Query("SELECT actor_id, first_name,last_name,last_update FROM actor")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			// "message": "pong",
		})
		return
	}

	for rows.Next() {
		var singleRecord Actor
		if err := rows.Scan(&singleRecord.Actor_id, &singleRecord.First_name, &singleRecord.Last_name, &singleRecord.Last_update); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}
		posts = append(posts, singleRecord)
	}
	c.JSON(http.StatusOK, posts)
}

func GetRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad error. Use only integer.",
		})
		return
	}
	row := db_client.DBClient.QueryRow("select actor_id, first_name,last_name,last_update from actor where actor_id = $1;", id)
	var post Actor
	if err := row.Scan(&post.Actor_id, &post.First_name, &post.Last_name, &post.Last_update); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "id not found.",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}
