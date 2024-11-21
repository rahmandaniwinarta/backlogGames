package controllers

import (
	"backlogGames/database"
	"backlogGames/repository"
	"backlogGames/structs"
	"fmt"
	"net/http"
	"strconv"

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

func GetAllGenres(c *gin.Context) {
	genres, err := repository.GetAllGenres(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genres)
}

func GetGenreByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("genreId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	genre, err := repository.GetGenreByID(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genre)
}

func UpdateGenre(c *gin.Context) {
	genreID, err := strconv.Atoi(c.Param("genreId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	var req struct {
		Name      string `json:"name"`
		UpdatedBy string `json:"updated_by"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre name cannot be empty"})
		return
	}

	// Set updated_by dari context jika tidak disediakan
	if req.UpdatedBy == "" {
		username, _ := c.Get("username") // Ambil dari context (jika ada middleware)
		req.UpdatedBy = username.(string)
	}

	genre := structs.Genre{
		ID:        genreID,
		Name:      req.Name,
		UpdatedBy: req.UpdatedBy,
	}

	err = repository.UpdateGenre(database.DbConnection, genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre updated successfully"})
}

func SoftDeleteGenre(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("genreId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	username, _ := c.Get("username")

	err = repository.SoftDeleteGenre(database.DbConnection, id, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}
