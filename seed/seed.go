// Package seed contains methods that generate seeds for random number generators that hopefully don't completely suck.
// Note very carefully: this first cut is super naive and not well thought out. Just chucking Mac Addresses into the mix.
package seed

import (
	"crypto/sha256"
	"net"
	"time"
)

// MACSeed generates a seed based on the first located MAC address .
func MACSeed() int64 {
	var result int64
	hash := sha256.Sum256([]byte(getFirstMacAddress()))
	var i uint
	for i < 8 {
		result = or(result, i, hash[i:i+4])
		i++
	}
	return result % time.Now().UnixNano()
}

func or(v int64, shift uint, bs []byte) int64 {
	var result int64
	for _, b := range bs {
		result = result ^ int64(b)
	}
	return v | result<<shift
}

func getFirstMacAddress() string {
	ni, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, a := range ni {
		r := a.HardwareAddr.String()
		if r != "" {
			return r
		}
	}
	return ""
}
