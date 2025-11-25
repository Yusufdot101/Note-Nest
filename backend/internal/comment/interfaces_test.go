package comment

import (
	"testing"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

func TestValidateComment(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantValid bool
	}{
		{
			name:      "valid comment",
			content:   "this is a valid comment",
			wantValid: true,
		},
		{
			name:      "invalid comment: missing content",
			content:   "",
			wantValid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := validator.NewValidator()
			c := &comment{
				Content: test.content,
			}
			validateComment(v, c)

			if v.IsValid() != test.wantValid {
				t.Errorf("expected IsValid = %v, got IsValid = %v", test.wantValid, v.IsValid())
			}
		})
	}
}
