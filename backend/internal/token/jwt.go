package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int `json:"sub"`
	jwt.RegisteredClaims
}

func CreateJWT(userID int) (string, error) {
	expirationTime, err := time.ParseDuration(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		return "", err
	}

	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		return "", errors.New("JWT_ISSUER variable missing")
	}

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
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
