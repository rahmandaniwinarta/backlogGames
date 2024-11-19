package structs

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Games struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
	Genre string `json:"genre"`
	Price string `json:"price"`
	Stock string `json:"stock"`
}
