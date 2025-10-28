package validator

import (
	"regexp"
	"testing"
)

func TestNewValidator(t *testing.T) {
	t.Run("creating new Validator", func(t *testing.T) {
		v := NewValidator()
		if v.Errors == nil {
			t.Error("unexpected nil Errors")
		}
	})
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name      string
		errMap    map[string]string
		wantValid bool
	}{
		{
			name:      "empty errors map",
			errMap:    make(map[string]string),
			wantValid: true,
		},
		{
			name:   "non empty errors map",
			errMap: map[string]string{"key": "value"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewValidator()
			v.Errors = test.errMap
			isValid := v.IsValid()
			if isValid != test.wantValid {
				t.Errorf("expected isValid = %v, got isValid = %v", test.wantValid, isValid)
			}
		})
	}
}

func TestAddError(t *testing.T) {
	tests := []struct {
		name               string
		setupValidator     func(v *Validator)
		wantKey, wantValue string
		wantErrors         bool
	}{
		{
			name: "one error",
			setupValidator: func(v *Validator) {
				v.Errors = make(map[string]string) // clean the errors
				v.AddError("key1", "value1")
			},
			wantKey:    "key1",
			wantValue:  "value1",
			wantErrors: true,
		},
		{
			name: "key already present",
			setupValidator: func(v *Validator) {
				// v.Errors = make(map[string]string) // we don't clean the errors
				v.AddError("key1", "value2") // same key, different value
			},
			wantKey:    "key1", // should change if the key is present
			wantValue:  "value1",
			wantErrors: true,
		},
		{
			name: "empty key",
			setupValidator: func(v *Validator) {
				v.Errors = make(map[string]string)
				v.AddError("", "value")
			},
			wantKey:   "",
			wantValue: "value",
		},
		{
			name: "empty value",
			setupValidator: func(v *Validator) {
				v.Errors = make(map[string]string)
				v.AddError("key", "")
			},
			wantKey:   "key",
			wantValue: "",
		},
		{
			name: "empty key and value",
			setupValidator: func(v *Validator) {
				v.Errors = make(map[string]string)
				v.AddError("", "")
			},
			wantKey:   "",
			wantValue: "",
		},
	}

	v := NewValidator()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setupValidator(v)
			if test.wantErrors && len(v.Errors) == 0 {
				t.Fatalf("expected errors")
			}
			if !test.wantErrors && len(v.Errors) != 0 {
				t.Fatalf("expected no errors but got %v", v.Errors)
			}
			if !test.wantErrors && len(v.Errors) == 0 {
				return
			}
			if _, exists := v.Errors[test.wantKey]; !exists {
				t.Fatalf("not found wanted key = %s", test.wantKey)
			}
			if v.Errors[test.wantKey] != test.wantValue {
				t.Errorf("wanted value = %s for key = %s, got value = %s", test.wantValue, test.wantKey, v.Errors[test.wantKey])
			}
		})
	}
}

func TestCheckAddError(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		key, value string
		wantErrors bool
	}{
		{
			name:       "true condition",
			condition:  1+1 == 2,
			key:        "key",
			value:      "value",
			wantErrors: false,
		},
		{
			name:       "false condition",
			condition:  1+1 == 1,
			key:        "key",
			value:      "value",
			wantErrors: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewValidator()
			v.CheckAddError(test.condition, test.key, test.value)
			if test.wantErrors && len(v.Errors) == 0 {
				t.Fatal("expected errors")
			}
			if !test.wantErrors && len(v.Errors) != 0 {
				t.Fatalf("expected no errors, got %v", v.Errors)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		name        string
		rx          *regexp.Regexp
		value       string
		wantMatches bool
	}{
		{
			name:        "pattern matching",
			rx:          regexp.MustCompile("[a-z]{5}"), // five lowercase letters
			value:       "abcde",
			wantMatches: true,
		},
		{
			name:  "pattern not matching",
			rx:    regexp.MustCompile("[a-z]{3,}"), // at least three lowercase letters
			value: "a",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := Matches(test.rx, test.value)
			if matches != test.wantMatches {
				t.Errorf("wanted matches = %v, got matches = %v", test.wantMatches, matches)
			}
		})
	}
}
