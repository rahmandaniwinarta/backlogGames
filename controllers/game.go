package controllers

import (
	"backlogGames/database"
	"backlogGames/repository"
	"backlogGames/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertGames(c *gin.Context) {
	var game structs.Games

	if err := c.BindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	username, _ := c.Get("username")
	game.CreatedBy = username.(string)
	game.UpdatedBy = username.(string)

	if game.Title == "" || game.GenreID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Error occurs () : Title Cannot Be Null",
		})
		return
	}

	isValidGenre, err := repository.IsGenreIDValid(database.DbConnection, game.GenreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to validate genre ID: " + err.Error(),
		})
		return
	}

	if !isValidGenre {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid genre_id: no such genre exists",
		})
		return
	}

	game, err = repository.InsertGame(database.DbConnection, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Game successfully inserted",
		"data":    game,
	})
}

func GetAllGames(c *gin.Context) {
	games, err := repository.GetAllGames(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to fetch games: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Games retreived successfully",
		"data":    games,
	})
}

func GetGameByID(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID parameter",
		})
		return
	}

	game, err := repository.GetGameByID(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to fetch game: " + err.Error(),
		})
		return
	}

	if game == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Game not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Game retrieved successfully",
		"data":    game,
	})
}

func GetGamesByGenre(c *gin.Context) {
	genreIDParam := c.Param("genre_id")
	genreID, err := strconv.Atoi(genreIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid genre_id parameter",
		})
		return
	}

	games, err := repository.GetGamesByGenre(database.DbConnection, genreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to fetch games by genre: " + err.Error(),
		})
		return
	}

	if len(games) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "No games found for the given genre",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Games retrieved successfully",
		"data":    games,
	})
}

func UpdateGame(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID parameter",
		})
		return
	}

	var game structs.Games
	if err := c.BindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid JSON body: " + err.Error(),
		})
		return
	}

	isValidGenre, err := repository.IsGenreIDValid(database.DbConnection, game.GenreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to validate genre ID: " + err.Error(),
		})
		return
	}

	if !isValidGenre {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid genre_id: no such genre exists",
		})
		return
	}

	game.ID = id

	err = repository.UpdateGame(database.DbConnection, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to update game: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Game successfully updated",
	})

}

func DeleteGames(c *gin.Context) {

}
