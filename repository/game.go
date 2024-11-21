package repository

import (
	"backlogGames/structs"
	"database/sql"
	"fmt"
)

func InsertGame(db *sql.DB, games structs.Games) (structs.Games, error) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Transaction Games Begin Error:", err)
		return games, fmt.Errorf("failed to begin Games transaction: %w", err)
	}

	count := 0
	query := `SELECT count(1) from games where title = $1`
	errors := db.QueryRow(query, games.Title).Scan(&count)
	if errors != sql.ErrNoRows && errors != nil {
		return games, fmt.Errorf("error during checking the games title : %w", errors)
	}

	if count > 0 {
		return games, fmt.Errorf("this game already existed")
	}

	sqlQuery := `INSERT INTO games (
				title, genre_id, price, stock, created_by, updated_by)
				VALUES(
				$1, $2, $3, $4, $5, $6)
				RETURNING id`

	err = tx.QueryRow(sqlQuery, games.Title, games.GenreID, games.Price, games.Stock, games.CreatedBy, games.UpdatedBy).Scan(&games.ID)
	if err != nil {
		fmt.Println("QueryRow Error:", err)
		tx.Rollback()
		return games, fmt.Errorf("error inserting games: %w", err)
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Transaction Commit Error:", err)
		return games, fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Game inserted with ID", games.ID)
	return games, nil
}

func GetAllGames(db *sql.DB) ([]structs.Games, error) {
	var games []structs.Games
	query := `SELECT id, title, genre_id, price, stock, created_by, updated_by FROM games`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching games: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var game structs.Games
		if err := rows.Scan(&game.ID, &game.Title, &game.GenreID, &game.Price, &game.Stock, &game.CreatedBy, &game.UpdatedBy); err != nil {
			return nil, fmt.Errorf("error scanning game: %w", err)
		}
		games = append(games, game)
	}

	return games, nil
}

func GetGameByID(db *sql.DB, id int) (*structs.Games, error) {
	var game structs.Games
	query := `SELECT id, title, genre_id, price, stock, created_by, updated_by FROM games WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&game.ID, &game.Title, &game.GenreID, &game.Price, &game.Stock, &game.CreatedBy, &game.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching game by ID: %w", err)
	}

	return &game, nil
}

func UpdateGame(db *sql.DB, games structs.Games) error {
	query := `
        UPDATE games
        SET title = $1, genre_id = $2, price = $3, stock = $4, updated_by = $5, updated_at = CURRENT_TIMESTAMP
        WHERE id = $6
    `
	_, err := db.Exec(query, games.Title, games.GenreID, games.Price, games.Stock, games.UpdatedBy, games.ID)
	if err != nil {
		return fmt.Errorf("error updating game: %w", err)
	}
	return nil
}

func DeleteGame(db *sql.DB, id int) error {
	query := `DELETE FROM games WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting game: %w", err)
	}
	return nil
}

func GetGamesByGenre(db *sql.DB, genreID int) ([]structs.Games, error) {
	var games []structs.Games
	query := `SELECT id, title, genre_id, price, stock, created_by, updated_by FROM games WHERE genre_id = $1`

	rows, err := db.Query(query, genreID)
	if err != nil {
		return nil, fmt.Errorf("error fetching games by genre: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var game structs.Games
		if err := rows.Scan(&game.ID, &game.Title, &game.GenreID, &game.Price, &game.Stock, &game.CreatedBy, &game.UpdatedBy); err != nil {
			return nil, fmt.Errorf("error scanning game: %w", err)
		}
		games = append(games, game)
	}

	return games, nil
}

func IsGenreIDValid(db *sql.DB, genreID int) (bool, error) {
	var count int
	query := `SELECT COUNT(1) FROM genres WHERE id = $1`

	err := db.QueryRow(query, genreID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error validating genre ID: %w", err)
	}

	return count > 0, nil
}
