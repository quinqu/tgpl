package minMax

import "testing"

func TestValidMinInput(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{([]int{1, 2, 3, 4}), 1},
		{([]int{100, 2, 3, 4}), 2},
		{([]int{1, 2, -33, 4}), -33},
	}

	for _, test := range tests {
		output, _ := min(test.input...)

		if output != test.want {
			t.Error(test.input, ": ", output)
		}
	}
}
func TestInvalidMinInput(t *testing.T) {

	_, err := min()

	if err == nil {
		t.Error("error was not returned for invalid input")
	}

}

func TestValidMaxInput(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{([]int{1, 2, 3, 4}), 4},
		{([]int{100, 2, 3, 4}), 100},
		{([]int{1, 22, -33, 4}), 22},
	}

	for _, test := range tests {
		output, _ := max(test.input...)

		if output != test.want {
			t.Error(test.input, ": ", output)
		}
	}

}

func TestInvalidMaxInput(t *testing.T) {
	_, err := max()

	if err == nil {
		t.Error("error was not returned for invalid input")
	}

}
