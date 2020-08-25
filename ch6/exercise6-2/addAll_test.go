package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	setOne := &IntSet{}
	setTwo := &IntSet{}
	setThree := &IntSet{}
	setOne.Add(1)
	setOne.Add(6)
	setOne.AddAll(7, 8, 9)
	setTwo.Add(8)
	setTwo.AddAll(24, 54)
	setThree.Add(1)
	setThree.AddAll(2, 3)

	var tests = []struct {
		set  *IntSet
		want []int
	}{
		{setOne, []int{1, 6, 7, 8, 9}},
		{setTwo, []int{8, 24, 54}},
		{setThree, []int{1, 2, 3}},
	}

	for _, test := range tests {
		for i, word := range test.set.words {
			if word == 0 {
				continue
			}
			testWantLength := len(test.want)
			k := 0
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j)) != 0 {
					val := 64*i + j
					if val != test.want[k] {
						t.Error("Expected:", test.want)
						t.Error("Actual:", test.set.String())
					}
					if k < testWantLength {
						k++
					}
				}
			}
		}
	}

}
