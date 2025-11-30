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
	RESET   TokenUse = "reset"
)

type Token struct {
	TokenString string
	CreatedAt   time.Time
	UserID      int
	Expires     time.Time
	Use         TokenUse
}

type Repo interface {
	InsertToken(token *Token) error
	GetByTokenStringAndUse(tokenString string, tokenUse TokenUse) (*Token, error)
	DeleteByTokenStringAndUse(tokenString string, tokenUse TokenUse) error
}

type TokenService struct {
	Repo Repo
}
