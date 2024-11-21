package controllers

import (
	"backlogGames/database"
	"backlogGames/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Panggil repository untuk membuat cart baru
	cartID, err := repository.CreateCart(database.DbConnection, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	// Kirim response dengan cart ID
	c.JSON(http.StatusCreated, gin.H{"cart_id": cartID})
}

func AddToCart(c *gin.Context) {
	var req struct {
		CartID   int `json:"cart_id"`
		GameID   int `json:"game_id"`
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validasi keberadaan cart dan game
	if !repository.CartExists(database.DbConnection, req.CartID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	if !repository.GameExists(database.DbConnection, req.GameID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	// Ambil harga game dari database
	price, err := repository.GetGamePrice(database.DbConnection, req.GameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch game price"})
		return
	}

	// Hitung total price
	totalPrice := price * float64(req.Quantity)

	// Tambahkan item ke cart
	err = repository.InsertCartItem(database.DbConnection, req.CartID, req.GameID, req.Quantity, totalPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart", "total_price": totalPrice})
}

func GetCartItems(c *gin.Context) {
	// Ambil parameter dari URL
	cartID, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	// Panggil repository untuk mendapatkan item di cart
	items, err := repository.GetCartItemsByCartID(database.DbConnection, cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func UpdateCartItem(c *gin.Context) {
	// Ambil parameter dari URL
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Ambil data dari body request
	var req struct {
		Quantity   int     `json:"quantity"`
		TotalPrice float64 `json:"total_price"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Panggil repository untuk memperbarui item di cart
	err = repository.UpdateCartItem(database.DbConnection, itemID, req.Quantity, req.TotalPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item updated"})
}

func DeleteCartItem(c *gin.Context) {
	// Ambil parameter dari URL
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Panggil repository untuk menghapus item dari cart
	err = repository.DeleteCartItem(database.DbConnection, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted"})
}
