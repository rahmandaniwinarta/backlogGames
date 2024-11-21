package functions

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	secretKey []byte
)

type Claims struct {
	Username string `json:"username"` // Username pengguna
	Role     string `json:"role"`     // Role pengguna (admin/buyer)
	jwt.StandardClaims
}

// Init function to load the secret key from .env
func init() {
	wd, _ := os.Getwd()

	curDir := fmt.Sprint(wd, "/.env")

	err := godotenv.Load(curDir)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the secret key from environment variables
	secretKey = []byte(os.Getenv("JWTKEY"))
	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set in the environment variables")
	}
}

// EncodeJWT creates a new JWT token
func EncodeJWT(username, role string) (string, error) {
	fmt.Println("EncodeJWT - Username:", username, "Role:", role) // Debugging

	claims := Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "your_app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// DecodeJWT parses and validates the JWT token
func DecodeJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate token and claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
