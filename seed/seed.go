// Package seed overengineers a way to initialise the math/rand default Source from crypto/rand.
package seed

import (
	crand "crypto/rand"
	mrand "math/rand"
)

const (
	size = 8
)

// Set calls math/rand.Seed() with a value generated from crypto/rand.Read()
func Set() {
	bs := make([]byte, size)
	crand.Read(bs)
	var i uint64
	var s int64
	for i = 0; i < size; i++ {
		s = s | int64(bs[i])<<(i<<3)
	}
	mrand.Seed(s)
}
