package user

import (
	"os"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func TestMain(m *testing.M) {
	// setup
	utilities.LoadEnv(".env.test")

	// run the tests
	code := m.Run()

	// exit
	os.Exit(code)
}
