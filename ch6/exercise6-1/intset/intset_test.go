package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	x := &IntSet{}
	x.Add(127)
	x.Add(4091)
	if x.Len() != 2 {
		t.Error("x length is:", x.Len())
	}
}

func TestLenZero(t *testing.T) {
	x := &IntSet{}

	if x.Len() != 0 {
		t.Error("x length should be 0")
	}
}

func TestRemove(t *testing.T) {
	x := &IntSet{}
	x.Add(127)
	x.Add(4091)
	x.Remove(4091)

	if x.Len() != 1 {
		t.Error("x length is:", x.Len(), "nothing was removed")
	}
	result := "{127}"

	if x.String() != result {
		t.Error("removed wrong int")
	}
}

func TestClear(t *testing.T) {
	x := &IntSet{}
	x.Add(127)
	x.Add(4091)
	x.Remove(4091)
	x.Add(48)
	x.Add(41)
	x.Clear()

	if x.Len() != 0 {
		t.Error("IntSet was not cleared")

	}
}

func TestCopy(t *testing.T) {
	x := &IntSet{}
	x.Add(127)
	x.Add(4091)

	newCopy := x.Copy()
	for i, _ := range x.words {

		if newCopy.words[i] != x.words[i] {
			t.Error("values not copied over")
		}

	}

}
