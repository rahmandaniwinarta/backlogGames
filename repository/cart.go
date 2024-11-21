package repository

import (
	"backlogGames/structs"
	"database/sql"
	"fmt"
)

func CreateCart(db *sql.DB, userID int) (int, error) {
	query := `INSERT INTO carts (user_id) VALUES ($1) RETURNING id`
	var cartID int
	err := db.QueryRow(query, userID).Scan(&cartID)
	if err != nil {
		return 0, fmt.Errorf("failed to create cart: %v", err)
	}
	return cartID, nil
}

// InsertCart: Membuat entri baru di tabel carts
func InsertCart(db *sql.DB, userID int) (int, error) {
	query := `INSERT INTO carts (user_id) VALUES ($1) RETURNING id`
	var cartID int
	err := db.QueryRow(query, userID).Scan(&cartID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert cart: %v", err)
	}
	return cartID, nil
}

// InsertCartItem: Menambahkan item ke dalam cart_items
func InsertCartItem(db *sql.DB, cartID, gameID, quantity int, totalPrice float64) error {
	// Validasi keberadaan cart dan game
	if !CartExists(db, cartID) {
		return fmt.Errorf("cart with ID %d does not exist", cartID)
	}
	if !GameExists(db, gameID) {
		return fmt.Errorf("game with ID %d does not exist", gameID)
	}

	// Insert item ke dalam cart_items
	query := `
    INSERT INTO cart_items (cart_id, game_id, quantity, total_price)
    VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, cartID, gameID, quantity, totalPrice)
	if err != nil {
		return fmt.Errorf("failed to insert cart item: %v", err)
	}
	return nil
}

// GetCartItemsByCartID: Mendapatkan semua item dalam cart berdasarkan cartID
func GetCartItemsByCartID(db *sql.DB, cartID int) ([]structs.CartItem, error) {
	// Validasi keberadaan cart
	if !CartExists(db, cartID) {
		return nil, fmt.Errorf("cart with ID %d does not exist", cartID)
	}

	// Query untuk mendapatkan item dalam cart
	query := `SELECT id, game_id, quantity, total_price FROM cart_items WHERE cart_id = $1`
	rows, err := db.Query(query, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}
	defer rows.Close()

	// Parsing hasil query ke dalam struct CartItem
	var items []structs.CartItem
	for rows.Next() {
		var item structs.CartItem
		if err := rows.Scan(&item.ID, &item.GameID, &item.Quantity, &item.TotalPrice); err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %v", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// UpdateCartItem: Memperbarui jumlah dan total harga item dalam cart
func UpdateCartItem(db *sql.DB, itemID, quantity int, totalPrice float64) error {
	// Validasi keberadaan item
	if !CartItemExists(db, itemID) {
		return fmt.Errorf("cart item with ID %d does not exist", itemID)
	}

	// Update item dalam cart_items
	query := `
    UPDATE cart_items
    SET quantity = $2, total_price = $3
    WHERE id = $1`
	_, err := db.Exec(query, itemID, quantity, totalPrice)
	if err != nil {
		return fmt.Errorf("failed to update cart item: %v", err)
	}
	return nil
}

// DeleteCartItem: Menghapus item dari cart
func DeleteCartItem(db *sql.DB, itemID int) error {
	// Validasi keberadaan item
	if !CartItemExists(db, itemID) {
		return fmt.Errorf("cart item with ID %d does not exist", itemID)
	}

	// Hapus item dari cart_items
	query := `DELETE FROM cart_items WHERE id = $1`
	_, err := db.Exec(query, itemID)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %v", err)
	}
	return nil
}

// CartExists: Mengecek apakah cart dengan ID tertentu ada
func CartExists(db *sql.DB, cartID int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM carts WHERE id = $1)`
	err := db.QueryRow(query, cartID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// GameExists: Mengecek apakah game dengan ID tertentu ada
func GameExists(db *sql.DB, gameID int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM games WHERE id = $1)`
	err := db.QueryRow(query, gameID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// CartItemExists: Mengecek apakah item dalam cart dengan ID tertentu ada
func CartItemExists(db *sql.DB, itemID int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM cart_items WHERE id = $1)`
	err := db.QueryRow(query, itemID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func GetGamePrice(db *sql.DB, gameID int) (float64, error) {
	query := `SELECT price FROM games WHERE id = $1`
	var price float64
	err := db.QueryRow(query, gameID).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch game price: %v", err)
	}
	return price, nil
}
