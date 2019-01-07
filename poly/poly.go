// Package poly follows chapter 2 of A Programmer's Introduction to Mathematics (J Kun): Polynomials
package poly

import "math"

// Eval calculates the value of a polynomial given a slice of coefficients of degree equal to their index values and an x value.
func Eval(coeffs []float64, x float64) float64 {
	if len(coeffs) == 0 {
		return 0
	}
	if x == 0 {
		return coeffs[0]
	}
	result := 0.0
	for d, c := range coeffs {
		result += c * math.Pow(x, float64(d))
	}
	return result
}
