package sets

import "testing"

func getSimpleTest() (setA, setB map[int][]int, pairs map[int]int) {
	setA = map[int][]int{
		1: []int{4, 5, 6},
		2: []int{6, 5, 4},
		3: []int{5, 6, 4},
	}
	setB = map[int][]int{
		4: []int{1, 2, 3},
		5: []int{1, 2, 3},
		6: []int{1, 2, 3},
	}
	pairs = map[int]int{
		4: 1,
		5: 3,
		6: 2,
	}
	return
}

func okPairs(pairs, test map[int]int) bool {
	if len(pairs) != len(test) {
		return false
	}
	for k, v := range pairs {
		if test[k] != v {
			return false
		}
	}
	return true
}

func TestSimpleStableMarriage(t *testing.T) {
	setA, setB, test := getSimpleTest()
	pairs, err := StableMarriage(setA, setB)
	if err != nil {
		t.Fatal(err)
	}
	if !okPairs(pairs, test) {
		t.Error("Unexpected pairs result", pairs)
	}
}

func TestSimpleIsStable(t *testing.T) {
	setA, setB, test := getSimpleTest()
	ok, err := IsStable(setA, setB, test)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("Unexpected failure")
	}
}
