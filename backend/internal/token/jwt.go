package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(userID int) (string, error) {
	expirationTime, err := time.ParseDuration(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userID,
		"iss": "note-nest",
		"exp": time.Now().Add(expirationTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT_SECRET variable is not set")
	}

	if len(jwtSecret) < 32 {
		return "", errors.New("JWT_SECRET variable must be at least 32 bytes for HS256")
	}

	return token.SignedString(jwtSecret)
}
