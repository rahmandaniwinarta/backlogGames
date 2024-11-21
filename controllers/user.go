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

	// Bind JSON body ke struct User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Enkripsi password untuk mencocokkan dengan database
	encryptedPass, err := functions.PasswordGenerator(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error encrypting password: " + err.Error(),
		})
		return
	}

	// Validasi user di database
	err = repository.GetUser(database.DbConnection, &user, encryptedPass)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Invalid username or password",
		})
		return
	}

	// Log role sebelum memanggil EncodeJWT
	fmt.Println("LoginUser - Username:", user.Username, "Role:", user.Role)

	// Generate token JWT
	token, err := functions.EncodeJWT(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error generating token: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"token": token,
	})
}
