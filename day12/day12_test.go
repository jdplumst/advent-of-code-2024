package day12

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Want        int
	}{
		{"given", "input_test.txt", 140},
		{"given second", "input_test_second.txt", 772},
		{"given large", "input_test_large.txt", 1930},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := part1(test.Input)
			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}