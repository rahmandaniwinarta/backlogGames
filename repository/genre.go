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

func GetGenreByID(db *sql.DB, id int) (structs.Genre, error) {
	query := `SELECT id, name, created_at, created_by, updated_at, updated_by FROM genres WHERE id = $1 AND deleted_at IS NULL`
	var genre structs.Genre
	err := db.QueryRow(query, id).Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.CreatedBy, &genre.UpdatedAt, &genre.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return genre, fmt.Errorf("genre not found")
		}
		return genre, fmt.Errorf("failed to fetch genre by ID: %w", err)
	}
	return genre, nil
}

func GetAllGenres(db *sql.DB) ([]structs.Genre, error) {
	query := `SELECT id, name, created_at, created_by, updated_at, updated_by FROM genres WHERE deleted_at IS NULL`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch genres: %w", err)
	}
	defer rows.Close()

	var genres []structs.Genre
	for rows.Next() {
		var genre structs.Genre
		if err := rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.CreatedBy, &genre.UpdatedAt, &genre.UpdatedBy); err != nil {
			return nil, fmt.Errorf("failed to scan genre: %w", err)
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func UpdateGenre(db *sql.DB, genre structs.Genre) error {
	query := `
    UPDATE genres
    SET name = $1, updated_by = $2, updated_at = CURRENT_TIMESTAMP
    WHERE id = $3 AND deleted_at IS NULL`
	result, err := db.Exec(query, genre.Name, genre.UpdatedBy, genre.ID)
	if err != nil {
		return fmt.Errorf("failed to update genre: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("genre not found or already deleted")
	}
	return nil
}

func SoftDeleteGenre(db *sql.DB, id int, deletedBy string) error {
	query := `
    UPDATE genres
    SET deleted_at = CURRENT_TIMESTAMP, updated_by = $1
    WHERE id = $2 AND deleted_at IS NULL`
	result, err := db.Exec(query, deletedBy, id)
	if err != nil {
		return fmt.Errorf("failed to delete genre: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("genre not found or already deleted")
	}
	return nil
}
