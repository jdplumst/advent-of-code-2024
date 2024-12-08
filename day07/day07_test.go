package day07

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Want        int
	}{
		{"given", "input_test.txt", 3749},
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
