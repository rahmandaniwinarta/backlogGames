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

	// Auth Routes
	router.POST("/api/register/buyer", controllers.RegisterUserBuyer)
	router.POST("/api/register/admin", controllers.RegisterUserAdmin)
	router.POST("/api/login", controllers.LoginUser)

	// Games Routes
	gamesGroup := router.Group("/api/games", middlewares.AuthMiddleware())
	{
		gamesGroup.POST("/", middlewares.AdminMiddleware(), controllers.InsertGames)      // Insert Game
		gamesGroup.GET("/", controllers.GetAllGames)                                      // Get All Games
		gamesGroup.GET("/:id", controllers.GetGameByID)                                   // Get Game by ID
		gamesGroup.GET("/genre/:genre_id", controllers.GetGamesByGenre)                   // Get Games by Genre
		gamesGroup.PUT("/:id", middlewares.AdminMiddleware(), controllers.UpdateGame)     // Update Game
		gamesGroup.DELETE("/:id", middlewares.AdminMiddleware(), controllers.DeleteGames) // Delete Game
	}

	// Genre Routes
	genreGroup := router.Group("/api/genre", middlewares.AuthMiddleware())
	{
		genreGroup.POST("/", middlewares.AdminMiddleware(), controllers.InsertGenre)               // Insert Genre
		genreGroup.GET("/", controllers.GetAllGenres)                                              // Get All Genres
		genreGroup.GET("/:genreId", controllers.GetGenreByID)                                      // Get Genre by ID
		genreGroup.PUT("/:genreId", middlewares.AdminMiddleware(), controllers.UpdateGenre)        // Update Genre
		genreGroup.DELETE("/:genreId", middlewares.AdminMiddleware(), controllers.SoftDeleteGenre) // Delete Genre
	}

	// Cart Routes
	cartGroup := router.Group("/api/cart", middlewares.AuthMiddleware())
	{
		cartGroup.POST("/", controllers.CreateCart)              // Create Cart
		cartGroup.POST("/item", controllers.AddToCart)           // Add Item to Cart
		cartGroup.GET("/:cartId/item", controllers.GetCartItems) // Get Cart Items
		cartGroup.PUT("/:itemId", controllers.UpdateCartItem)    // Update Cart Item
		cartGroup.DELETE("/:itemId", controllers.DeleteCartItem) // Delete Cart Item
	}

	// Transaction Routes
	transactionGroup := router.Group("/api/transaction", middlewares.AuthMiddleware())
	{
		transactionGroup.POST("/", controllers.CreateTransaction)                    // Create Transaction
		transactionGroup.PUT("/:transactionId", controllers.UpdateTransactionStatus) // Update Transaction Status
		transactionGroup.GET("/:transactionId", controllers.GetTransactionByID)      // Get Transaction by ID
	}

	// Start Server
	router.Run(":8080")
}
