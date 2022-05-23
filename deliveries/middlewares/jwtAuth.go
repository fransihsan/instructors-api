package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Generate jwt token
func GenerateToken(ID uint) (string, error) {
	if ID < 1 {
		return "", errors.New("user_id tidak valid")
	}
	data := jwt.MapClaims{}
	data["id"] = ID
	data["exp"] = time.Now().Add(time.Minute * 120).Unix()
	data["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// Get user id
func ExtractTokenUserID(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		id := uint(data["id"].(float64))
		return id
	}
	return 0
}
