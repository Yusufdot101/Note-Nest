package validator

import (
	"errors"
	"regexp"
)

var ErrFailedValidation = errors.New("failed validation")

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, value string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = value
	}
}

func (v *Validator) CheckAddError(condition bool, key, value string) {
	if !condition {
		v.AddError(key, value)
	}
}

func Matches(rx *regexp.Regexp, value string) bool {
	return rx.MatchString(value)
}
