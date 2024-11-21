package main

import (
	"backlogGames/controllers"
	"backlogGames/database"
	"backlogGames/middlewares"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASS"),
		os.Getenv("PGNAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer func() {
		if err := DB.Close(); err != nil {
			fmt.Println("Error : ", err)
		}
	}()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB, "up")

	fmt.Println("Successfully connected!")

	router := gin.Default()

	router.POST("/api/register/buyer", controllers.RegisterUserBuyer)
	router.POST("/api/register/admin", controllers.RegisterUserAdmin)
	router.POST("/api/login", controllers.LoginUser)
	router.POST("/api/genre", controllers.InsertGenre)
	router.POST("/api/games", middlewares.AuthMiddleware(), middlewares.AdminMiddleware(), controllers.InsertGames)
	router.GET("/api/games", controllers.GetAllGames)
	router.GET("/api/games/:id", controllers.GetGameByID)
	router.PUT("/api/games/:id", controllers.UpdateGame)
	router.GET("/api/games/genre/:genre_id", controllers.GetGamesByGenre)

	router.Run(":8080")
}
