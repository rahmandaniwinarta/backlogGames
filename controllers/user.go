package controllers

import (
	"backlogGames/database"
	"backlogGames/repository"
	"backlogGames/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RegisterUserBuyer(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	err = repository.InsertUserBuyer(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (1)" + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, message{
			Code:    http.StatusOK,
			Message: "User Created : " + user.Name,
		})
	}

}
