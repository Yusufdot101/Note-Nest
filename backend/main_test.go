package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name       string
		a, b, want int
	}{
		{
			name: "1+2",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "225+75",
			a:    225,
			b:    75,
			want: 300,
		},
		{
			name: "77+33",
			a:    77,
			b:    33,
			want: 100,
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			got := addTwoNumbers(test.a, test.b)
			if got != test.want {
				t.Errorf("wanted %d, got %d", test.want, got)
			}
		})
	}
}
