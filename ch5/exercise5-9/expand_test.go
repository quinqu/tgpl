package expand 

import "testing"


func TestExpand(t *testing.T) {
	var tests = []struct {
		input string
		output string 
	}{
		{"$foo $bar $hello$hi", "foofoo barbar hello$hihello$hi"}, 
		{"testing $123 testing$", "testing 123123 testing$"}, 
		{"go $go $go asd $fghjkl", "go gogo gogo asd fghjklfghjkl"}, 
		{"$ $ $", "$$ $$ $$"}, 
		{"", ""}, 
		{"$helloworld", "helloworldhelloworld"},
	}

	for _, test := range tests {
		if got := expand(test.input, f); got != test.output {
			t.Errorf("Expected: %q = %v", test.input, test.output)
			t.Errorf("Actual: %q = %v", test.input, got)
		}
	}
}