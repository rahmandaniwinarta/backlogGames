package structs

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Games struct {
	ID         int       `json:"id"`
	GenreID    int       `json:"genre_id"`
	Title      string    `json:"title"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	GameID     int       `json:"game_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

type Genre struct {
	ID         int       `json:"id"`
	Genre      string    `json:"genre"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}
