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
