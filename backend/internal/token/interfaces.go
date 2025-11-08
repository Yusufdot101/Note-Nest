package token

import "time"

type (
	TokenType string
	TokenUse  string
)

const (
	JWT          TokenType = "jwt"
	RANDOMSTRING TokenType = "random string"

	ACCESS  TokenUse = "access"
	REFRESH TokenUse = "refresh"
)

type Token struct {
	TokenString string
	CreatedAt   time.Time
	UserID      int
	Expires     time.Time
}

type Repo interface {
	InsertToken(token *Token) error
	GetByTokenString(tokenString string) (*Token, error)
	DeleteByTokenString(tokenString string) error
}

type TokenService struct {
	Repo Repo
}
