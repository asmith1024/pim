// Package poly follows chapter 2 of A Programmer's Introduction to Mathematics (J Kun): Polynomials
package poly

import (
	"fmt"
	"math"
	"strings"
)

// Y calculates the value of a polynomial given a point on the x-axis and a slice of coefficients. The degree of the polynomial is the slice high index.
func Y(x float64, coeffs []float64) float64 {
	if len(coeffs) == 0 {
		return 0
	}
	if x == 0 {
		return coeffs[0]
	}
	result := 0.0
	for e, c := range coeffs {
		if c != 0 {
			result += c * math.Pow(x, float64(e))
		}
	}
	return result
}

// Print writes a polynomial in the format f(x) = ax^n +|- bx^(n-1) ... +|- cx + d.
func Print(coeffs []float64) string {
	if zero(coeffs) {
		return ""
	}
	var b strings.Builder
	signs := signs(coeffs)
	b.WriteString("f(x) = ")
	wroteTerm := false
	for i := len(coeffs) - 1; i >= 0; i-- {
		wroteTerm = writeTerm(coeffs[i], i, signs[i], wroteTerm, &b)
	}
	return b.String()
}

func zero(coeffs []float64) bool {
	for _, c := range coeffs {
		if c != 0 {
			return false
		}
	}
	return true
}

func signs(coeffs []float64) []byte {
	s := make([]byte, len(coeffs))
	for i, c := range coeffs {
		if c < 0 {
			s[i] = '-'
		} else {
			s[i] = '+'
		}
	}
	return s
}

func writeTerm(c float64, i int, s byte, begun bool, b *strings.Builder) bool {
	if c == 0 {
		return false || begun
	}
	if begun {
		c = math.Abs(c)
		b.WriteString(fmt.Sprintf(" %c ", s))
	}
	switch i {
	case 0:
		b.WriteString(fmt.Sprintf("%g", c))
	case 1:
		b.WriteString(fmt.Sprintf("%gx", c))
	default:
		b.WriteString(fmt.Sprintf("%gx^%d", c, i))
	}
	return true
}
