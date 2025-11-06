package token

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Repo interface {
	InsertToken(token *Token) error
	GetByTokenString(tokenString string) (*Token, error)
	DeleteByTokenString(tokenString string) error
}

type TokenService struct {
	Repo Repo
}

func (ts *TokenService) NewToken(tokenType TokenType, tokenUse TokenUse, userID int) (string, error) {
	var ttl time.Duration
	var err error
	switch tokenUse {
	case REFRESH:
		ttl, err = time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRATION_TIME"))
	case ACCESS:
		ttl, err = time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRATION_TIME"))
	default:
		err = errors.New("invalid tokenType")
	}
	if err != nil {
		return "", err
	}

	var token string
	switch tokenType {
	case JWT:
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		token, err = createJWT(jwtSecret, ttl, userID)
	case RANDOMSTRING:
		token, err = ts.generateRandomToken(ttl, userID)
	default:
		err = errors.New("invalid token type")
	}

	if err != nil {
		return "", err
	}

	return token, nil
}

func (ts *TokenService) DeleteToken(tokenString string) error {
	return ts.Repo.DeleteByTokenString(tokenString)
}

func createJWT(jwtSecret []byte, ttl time.Duration, userID int) (string, error) {
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		return "", errors.New("JWT_ISSUER variable missing")
	}

	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		Subject:   strconv.Itoa(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if len(jwtSecret) == 0 {
		return "", errors.New("JWT_SECRET variable is not set")
	}

	if len(jwtSecret) < 32 {
		return "", errors.New("JWT_SECRET variable must be at least 32 bytes for HS256")
	}

	return token.SignedString(jwtSecret)
}

func (ts *TokenService) generateRandomToken(ttl time.Duration, userID int) (string, error) {
	token := &Token{
		UserID:      userID,
		Expires:     time.Now().Add(ttl),
		TokenString: uuid.New().String(),
	}
	err := ts.Repo.InsertToken(token)
	return token.TokenString, err
}

func ValidateJWT(tokenString string, jwtSecret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// ensure the token was signed with HMAC, not something else
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		// custom_errors.InvalidAuthenticationTokenErrorResponse(w)
		return nil, errors.New("invalid token")
	}

	return token, nil
}
