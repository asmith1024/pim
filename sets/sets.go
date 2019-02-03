// Package sets follows chapter 4 of A Programmer's Introduction to Mathematics (J Kun): Sets.
// A set is modeled as a map. Since Go doesn't have generics (yet), specifics are context-dependent.
package sets

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/asmith1024/pim/seed"
)

// StableMarriage implements the "stable marriages" algorithm for sets where
// items from each set are identified by int and preferences into the other
// set are represented by a slice of int identifiers. The lower the index
// the greater the preference. There are (at least) two problems with this
// representation:
// 1. You have to separately enforce element uniqueness.
// 2. Checking preferences is not efficicient with large sets.
//
// By convention, "proposals" flow from setA to setB.
// The return value is the final set of pairs keyed by setB.
//
// setA and setB must be the same size.
// An identifier 0 means uninitialized.
func StableMarriage(setA, setB map[int][]int) (pairsBA map[int]int, err error) {
	if len(setA) != len(setB) {
		err = errors.New("sets A and B do not form a bijection")
		return
	}
	pairsBA = make(map[int]int, len(setB))
	for len(pairsBA) < len(setB) {
		for a, aprefs := range setA {
			for _, b := range aprefs {
				current, ok := pairsBA[b]
				if !ok {
					pairsBA[b] = a
					break
				}
				bprefs, ok := setB[b]
				if !ok {
					err = fmt.Errorf("%d was not found in set B but is referenced by %d in set A", b, a)
					return
				}
				if preferred(current, a, bprefs) {
					pairsBA[b] = a
					break
				}
			}
		}
		// there are dodgy parameter values that will cause this never to return
	}
	return
}

// FnPrefs is a preference allocation function used by InitSets. Implementions
// must expect both set member keys to be a number between 1 and size.
// A mapping is implied between these indices and whatever the source sets represent.
// isA is true if key is a member of setA, false if key is a member of setB.
type FnPrefs func(isA bool, key, size int) []int

// RandomPrefs (more-or-less) randomly assigns preferences.
func RandomPrefs(isA bool, key, size int) []int {
	result := make([]int, size)
	rnd := rand.New(rand.NewSource(seed.MACSeed()))
	for i := 1; i <= size; i++ {
		v := rnd.Intn(size-1) + 1
		for j := 0; j < i; j++ {
			if result[j] == v {
				break
			}
		}

	}
	return result
}

// InitSets takes a size and a function that allocates preferences and returns
// maps suitable as setA, setB parameters to StableMarriage and IsStable.
func InitSets(size int, fn FnPrefs) (setA, setB map[int][]int) {
	setA = make(map[int][]int, size)
	setB = make(map[int][]int, size)
	for i := 1; i <= size; i++ {
		setA[i] = fn(true, i, size)
		setB[i] = fn(false, i, size)
	}
	return
}

func preferred(idCurrent, idNew int, prefs []int) bool {
	for _, v := range prefs {
		switch {
		case v == idNew:
			return true
		case v == idCurrent:
			return false
		}
	}
	return true
}

// IsStable returns true if the set pairs represents a stable distribution of
// mutual preferences defined in sets A and B.
// If IsStable returns false, err identifies the first pair from pairs that
// is not stable. This method is like O(n!) so I suck.
func IsStable(setA, setB map[int][]int, pairsBA map[int]int) (bool, error) {
	pairsAB := make(map[int]int, len(pairsBA))
	for b, a := range pairsBA {
		pairsAB[a] = b
	}
	for b, a := range pairsBA {
		bprefs, ok := setB[b]
		if !ok {
			return ok, fmt.Errorf("identifier %d not found in set B", b)
		}
		for _, v := range bprefs {
			if v == a {
				break
			}
			currentB, ok := pairsAB[v]
			if !ok {
				return ok, fmt.Errorf("set A identifier %d is not paired with an identifier from set B", v)
			}
			aprefs, ok := setA[v]
			if preferred(currentB, b, aprefs) {
				return false, fmt.Errorf("set B identifier %d paired with set A identifier %d is more stable with set A identifier %d", b, a, v)
			}
		}
	}
	return true, nil
}
