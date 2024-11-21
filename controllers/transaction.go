package controllers

import (
	"backlogGames/database"
	"backlogGames/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var req struct {
		CartID int `json:"cart_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validasi keberadaan cart
	if !repository.CartExists(database.DbConnection, req.CartID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Buat transaksi baru
	transactionID, err := repository.CreateTransaction(database.DbConnection, req.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction_id": transactionID,
		"status":         "pending",
	})
}

func UpdateTransactionStatus(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validasi status
	allowedStatuses := map[string]bool{
		"pending":   true,
		"completed": true,
		"cancelled": true,
	}
	if !allowedStatuses[req.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction status"})
		return
	}

	// Update status transaksi
	err = repository.UpdateTransactionStatus(database.DbConnection, transactionID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction status updated"})
}

func GetTransactionByID(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := repository.GetTransactionByID(database.DbConnection, transactionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
