package controllers

import (
	"backlogGames/database"
	"backlogGames/functions"
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

	Encrypted, err := functions.PasswordGenerator(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (1): " + err.Error(),
		})
		return
	}

	user.Password = Encrypted

	err = repository.InsertUserBuyer(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (2)" + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, message{
			Code:    http.StatusOK,
			Message: "User Created : " + user.Username,
		})
	}

}

func RegisterUserAdmin(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	Encrypted, err := functions.PasswordGenerator(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (3): " + err.Error(),
		})
		return
	}

	user.Password = Encrypted

	err = repository.InsertUserAdmin(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (4)" + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, message{
			Code:    http.StatusOK,
			Message: "Admin Created : " + user.Username,
		})
	}
}

func LoginUser(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)

	if err != nil {
		fmt.Println(err)
	}

	EncryptedPass, err := functions.PasswordGenerator(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (15): " + err.Error(),
		})
		return
	}

	err = repository.GetUser(database.DbConnection, &user, EncryptedPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (16): " + err.Error(),
		})
		return
	}

	res, errors := functions.EncodeJWT(map[string]interface{}{
		"data":    user,
		"isLogin": true,
	})

	if errors != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (17): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":  http.StatusOK,
		"token": res,
	})
}
