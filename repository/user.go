package repository

import (
	"backlogGames/structs"
	"database/sql"
	"fmt"
)

func InsertUserBuyer(db *sql.DB, user structs.User) (err error) {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Transaction Begin Error:", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	sqlQuery := `
			INSERT INTO users 
			( username, email, password, role ) 
			 VALUES ($1,$2,$3,$4)
			 RETURNING id`

	err = tx.QueryRow(sqlQuery, user.Username, user.Email, user.Password, "buyer").Scan(&user.ID)
	if err != nil {
		fmt.Println("QueryRow Error:", err)
		tx.Rollback()
		return fmt.Errorf("error inserting buyer: %w", err)
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Transaction Commit Error:", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("User inserted with ID:", user.ID)
	return nil
}

func InsertUserAdmin(db *sql.DB, user structs.User) (err error) {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Transaction Begin Error:", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	sqlQuery := `
			INSERT INTO users 
			( username, email, password, role ) 
			 VALUES ($1,$2,$3,$4)
			 RETURNING id`

	err = tx.QueryRow(sqlQuery, user.Username, user.Email, user.Password, "admin").Scan(&user.ID)
	if err != nil {
		fmt.Println("QueryRow Error:", err)
		tx.Rollback()
		return fmt.Errorf("error inserting admin: %w", err)
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Transaction Commit Error:", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("User inserted with ID:", user.ID)
	return nil
}

func GetUser(db *sql.DB, user *structs.User, encryptedPass string) (err error) {

	fmt.Println(user)

	sqls := "SELECT id, username, password from users WHERE username = $1"

	errors := db.QueryRow(sqls, user.Username).Scan(&user.ID, &user.Username, &user.Password)

	if errors != nil {
		fmt.Println("Errors (13) :", errors)
		return fmt.Errorf("errors (13): wrong password or username")
	}

	if encryptedPass != user.Password {
		return fmt.Errorf("errors (14) : wrong password or username")
	}

	user.Password = ""

	return nil
}

func KeepLogin(db *sql.DB, user string) (err error) {

	var username string

	fmt.Println(user)

	sqls := "SELECT username from users WHERE username = $1"

	errors := db.QueryRow(sqls, user).Scan(&username)

	if errors != nil {
		fmt.Println("Errors (13)", errors)
		return fmt.Errorf("errors (13): wrong password or username")
	}

	return nil
}
