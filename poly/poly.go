// Package poly follows chapter 2 of A Programmer's Introduction to Mathematics (J Kun): Polynomials
package poly

import "math"

// Y calculates the value of a polynomial given a point on the x-axis and a slice of coefficients. The degree of the polynomial is the slice high index.
func Y(x float64, coeffs []float64) float64 {
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
