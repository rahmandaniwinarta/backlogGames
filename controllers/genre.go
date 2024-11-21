package controllers

import (
	"backlogGames/database"
	"backlogGames/repository"
	"backlogGames/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertGenre(c *gin.Context) {
	var genre structs.Genre

	err := c.BindJSON(&genre)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (??: )" + err.Error(),
		})
		return
	}

	if genre.Name == "" {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs () : Title Cannot Be Null",
		})
		return
	}

	fmt.Println(genre)

	username, _ := c.Get("username")
	genre.CreatedBy = username.(string)
	genre.UpdatedBy = username.(string)

	err = repository.InsertGenre(database.DbConnection, genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (54): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successfully created " + genre.Name,
	})
}
