package utilities

import (
	"net/url"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

func ReadInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be integer")
		return defaultValue
	}

	return i
}

func ReadStr(qs url.Values, key, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}
