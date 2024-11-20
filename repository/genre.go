package repository

import (
	"backlogGames/structs"
	"database/sql"
	"fmt"
)

func InsertGenre(db *sql.DB, genre structs.Genre) (err error) {
	tx, err := db.Begin()

	if err != nil {
		fmt.Println("Transaction Genre Begin Error:", err)
		return fmt.Errorf("failed to begin genre transaction: %w", err)
	}

	count := 0

	fmt.Println(genre)

	query := `SELECT count(1) from genres where name = $1`

	errors := db.QueryRow(query, genre.Name).Scan(&count)

	if errors != sql.ErrNoRows && errors != nil {
		return fmt.Errorf("error during cheking the genre name : %w", errors)
	}

	if count > 0 {
		return fmt.Errorf("this genre already existed")
	}

	sqlQuery := `INSERT INTO genres (
				name, created_by, updated_by)
				VALUES(
				$1,$2,$3)
				RETURNING id`

	err = tx.QueryRow(sqlQuery, genre.Name, genre.CreatedBy, genre.UpdatedBy).Scan(&genre.ID)
	if err != nil {
		fmt.Println("QueryRow Error:", err)
		tx.Rollback()
		return fmt.Errorf("error inserting genre: %w", err)
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Transaction Commit Error:", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Genre inserted with ID", genre.ID)

	return nil
}
