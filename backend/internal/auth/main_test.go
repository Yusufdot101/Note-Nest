package auth

import (
	"os"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

type mockTokenRepo struct {
	InsertTokenCalled               bool
	DeleteByTokenStringAndUseCalled bool
}

func (mtr *mockTokenRepo) InsertToken(token *token.Token) error {
	mtr.InsertTokenCalled = true
	return nil
}

func (mtr *mockTokenRepo) GetByTokenStringAndUse(tokenString string, tokenUse token.TokenUse) (*token.Token, error) {
	return nil, nil
}

func (mtr *mockTokenRepo) DeleteByTokenStringAndUse(tokenString string, tokUse token.TokenUse) error {
	mtr.DeleteByTokenStringAndUseCalled = true
	return nil
}

type mockUserRepo struct {
	InsertUserCalled     bool
	InsertedUser         *user.User
	GetUserByEmailCalled bool
	GotUser              *user.User
}

func (mur *mockUserRepo) InsertUser(u *user.User) error {
	mur.InsertUserCalled = true
	mur.InsertedUser = u
	return nil
}

func (mur *mockUserRepo) GetUserByEmail(email string) (*user.User, error) {
	mur.GetUserByEmailCalled = true
	mur.GotUser = &user.User{
		ID:    0,
		Name:  "yusuf",
		Email: "ym@gmail.com",
	}
	err := mur.GotUser.Password.Set("12345678")
	if err != nil {
		return nil, err
	}
	return mur.GotUser, nil
}

func (mur *mockUserRepo) UpdatePasswordUsingToken(token, newPassword string) error {
	return nil
}

func TestMain(m *testing.M) {
	// setup
	utilities.LoadEnv("test.env")

	// run the tests
	code := m.Run()

	// exit
	os.Exit(code)
}
