package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddAll(t *testing.T) {
	setOne := &IntSet{}
	setTwo := &IntSet{}
	setThree := &IntSet{}
	setOne.Add(1)
	setOne.Add(6)
	setOne.AddAll(7, 8, 99)
	setTwo.Add(8)
	setTwo.AddAll(24, 54)
	setThree.Add(1)
	setThree.AddAll(2, 76)

	var tests = []struct {
		set  *IntSet
		want []int
	}{
		{setOne, []int{1, 6, 7, 8, 99}},
		{setTwo, []int{8, 24, 54}},
		{setThree, []int{1, 2, 76}},
	}
	var gotSlice []int
	for _, test := range tests {
		t.Run(fmt.Sprintln("AddAllPass:", test.want), func(t *testing.T) {
			gotSlice = test.set.toSlice()
			if !equal(test.want, gotSlice) {
				t.Error("fail")
			}
		})
	}

}

func TestAddAllWithDupes(t *testing.T) {
	setOne := &IntSet{}
	setTwo := &IntSet{}

	setOne.Add(1)
	setOne.Add(6)
	setOne.AddAll(1, 2, 2000)
	setTwo.Add(8)
	setTwo.AddAll(24, 8)

	var tests = []struct {
		set  *IntSet
		want []int
	}{
		{setOne, []int{1, 2, 6, 2000}},
		{setTwo, []int{8, 24}},
	}
	var gotSlice []int

	for _, test := range tests {
		t.Run(fmt.Sprintln("AddAllNoDupesPass:", test.want), func(t *testing.T) {

			gotSlice = test.set.toSlice()
			if !equal(test.want, gotSlice) {
				t.Error("fail")
			}
		})
	}

}

func TestToSlice(t *testing.T) {
	setOne := &IntSet{}
	setOne.Add(1)
	setOne.Add(6)
	setOne.AddAll(1, 2, 2000)
	setOneSlice := setOne.toSlice()
	expected := []int{1, 2, 6, 2000}
	if reflect.TypeOf(setOneSlice) != reflect.TypeOf(expected) {
		t.Error("toSlice() did not convert to []int ")
	}

}

func TestEqual(t *testing.T) {
	a1 := []int{1, 3, 4}
	a2 := []int{1, 3, 4}
	a3 := []int{1, 5}

	if !equal(a1, a2) {
		t.Error("Test Equal failure: returned false, expected true")
	}

	if equal(a1, a3) {
		t.Error("Test Equal Failure: returned true, expected false")
	}

}

func equal(want []int, got []int) bool {
	if len(want) != len(got) {
		return false
	}

	for i, word := range got {
		if word != want[i] {
			return false
		}
	}
	return true
}

func (intset *IntSet) toSlice() []int {
	var newSlice []int
	for i, word := range intset.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				val := 64*i + j
				newSlice = append(newSlice, val)
			}
		}
	}
	return newSlice
}
