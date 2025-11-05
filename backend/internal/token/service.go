package token

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(tokenType string, userID int) (string, error) {
	var expirationTime time.Duration
	var err error

	switch tokenType {
	case "REFRESH":
		expirationTime, err = time.ParseDuration(os.Getenv("REFRESH_JWT_EXPIRATION_TIME"))
	case "ACCESS":
		expirationTime, err = time.ParseDuration(os.Getenv("ACCESS_JWT_EXPIRATION_TIME"))
	default:
		err = errors.New("invalid tokenType")
	}

	if err != nil {
		return "", err
	}

	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		return "", errors.New("JWT_ISSUER variable missing")
	}

	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		Subject:   strconv.Itoa(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtSecret []byte
	switch tokenType {
	case "REFRESH":
		jwtSecret = []byte(os.Getenv("REFRESH_JWT_SECRET"))
	case "ACCESS":
		jwtSecret = []byte(os.Getenv("ACCESS_JWT_SECRET"))
	default:
		jwtSecret = []byte(os.Getenv(""))
	}

	if len(jwtSecret) == 0 {
		return "", errors.New("JWT_SECRET variable is not set")
	}

	if len(jwtSecret) < 32 {
		return "", errors.New("JWT_SECRET variable must be at least 32 bytes for HS256")
	}

	return token.SignedString(jwtSecret)
}
