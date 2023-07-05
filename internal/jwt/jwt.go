package jwt

import (
	"os"
	"time"

	"github.com/acemouty/disney-pass/internal/database"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *database.User) (string, error) {

	jwtSecret := os.Getenv("JWT_SECRET")

	// fallback for local dev
	if jwtSecret == "" {
		jwtSecret = "keep it secret keep it safe"
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims (e.g., user ID, expiration time, etc.)
	claims["user_id"] = user.UserID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration
	claims["iat"] = time.Now().Unix()

	// Sign the token with a secret key
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
