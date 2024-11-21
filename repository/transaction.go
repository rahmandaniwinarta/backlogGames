package repository

import (
	"backlogGames/structs"
	"database/sql"
	"fmt"
)

func CreateTransaction(db *sql.DB, cartID int) (int, error) {
	// Hitung total transaksi dari cart
	queryTotal := `SELECT COALESCE(SUM(total_price), 0) FROM cart_items WHERE cart_id = $1`
	var totalTransaction float64
	err := db.QueryRow(queryTotal, cartID).Scan(&totalTransaction)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total transaction: %v", err)
	}

	// Buat transaksi baru
	query := `INSERT INTO transactions (cart_id, total_transaction, status) VALUES ($1, $2, 'pending') RETURNING id`
	var transactionID int
	err = db.QueryRow(query, cartID, totalTransaction).Scan(&transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %v", err)
	}

	return transactionID, nil
}

func UpdateTransactionStatus(db *sql.DB, transactionID int, status string) error {
	query := `
    UPDATE transactions
    SET status = $1, updated_at = CURRENT_TIMESTAMP
    WHERE id = $2`
	_, err := db.Exec(query, status, transactionID)
	if err != nil {
		return fmt.Errorf("failed to update transaction status: %v", err)
	}
	return nil
}

func GetTransactionByID(db *sql.DB, transactionID int) (*structs.Transaction, error) {
	query := `
    SELECT id, cart_id, total_transaction, status, created_at, updated_at
    FROM transactions
    WHERE id = $1`
	var transaction structs.Transaction
	err := db.QueryRow(query, transactionID).Scan(
		&transaction.ID,
		&transaction.CartID,
		&transaction.TotalTransaction,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction: %v", err)
	}
	return &transaction, nil
}
