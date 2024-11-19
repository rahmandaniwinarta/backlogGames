package repository

import (
	"backlogGames/structs"
	"database/sql"
	"fmt"
)

func InsertUserBuyer(db *sql.DB, user structs.User) (err error) {

	sqlQuery := `
			INSERT INTO users 
			( id , name, email, password, role , created_at) 
			 VALUES ($1,$2,$3,$4,$5,$6)`

	_, err = db.Exec(sqlQuery, user.ID, user.Name, user.Email, user.Password, "buyer", user.CreatedAt)
	if err != nil {
		fmt.Println("Error inserting buyer", err)
		return fmt.Errorf("error : %w", err)
	}

	return nil

}
