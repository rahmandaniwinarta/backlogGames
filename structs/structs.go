package structs

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Games struct {
	ID        int       `json:"id"`
	GenreID   int       `json:"genre_id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"Updated_at"`
	UpdatedBy string    `json:"Updated_by"`
}

type Transaction struct {
	ID        int       `json:"id"`
	CartID    int       `json:"cart_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type Cart struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	GameID     int     `json:"game_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type Genre struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"Updated_at"`
	UpdatedBy string    `json:"Updated_by"`
}
