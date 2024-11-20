package controllers

import (
	"backlogGames/functions"
	"backlogGames/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func insertGames(c *gin.Context) {
	var game structs.Games

	err := c.BindJSON(&game)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (??: )" + err.Error(),
		})
		return
	}

	if game.Title == "" {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs () : Title Cannot Be Null",
		})
		return
	}

	fmt.Println(game)

	key := c.GetHeader("Auth")

	err, user := functions.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (??)" + err.Error(),
		})
		return
	}

	game.CreatedBy = user
	game.UpdatedBy = user
}
