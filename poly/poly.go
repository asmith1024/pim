// Package poly follows chapter 2 of A Programmer's Introduction to Mathematics (J Kun): Polynomials.
// A polynomial is represented by a slice of float64. The values represent the coefficients with the
// index indicating the degree. For example, {1,-2,3,-4} = 1 - 2x + 3x^2 - 4x^3.
package poly

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Y calculates the value of a polynomial given a point on the x-axis and a slice of coefficients.
// The degree of the polynomial is the highest index in the slice with a non-zero value.
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

// Point describes a spot on a 2D cartesian plane. There isn't one defined in the Go math package.
type Point struct {
	X float64
	Y float64
}

// Interpolate calculates the unique polynomial described by points. An empty slice is returned if points is empty or contains a duplicate X value.
func Interpolate(points []Point) []float64 {
	if len(points) == 0 {
		return []float64{}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
	if hasDup(points) {
		return []float64{}
	}
	return pcalc(points)
}

func hasDup(points []Point) bool {
	for i := 0; i < len(points)-1; i++ {
		if points[i].X == points[i+1].X {
			return true
		}
	}
	return false
}

// Keep in mind that any time you see a slice of floats it's a polynomial.
func pcalc(points []Point) []float64 {
	numerator := make([][]float64, len(points))
	for i := range points {
		factors := make([][]float64, len(points)-1)
		denominator := 1.0
		fidx := 0
		for j, p := range points {
			if j == i { // this is why we need fidx
				continue
			}
			denominator *= (points[i].X - p.X)
			factors[fidx] = []float64{-1 * p.X, 1}
			fidx++
		}
		expansion := factors[0]
		for k := 1; k < len(factors); k++ {
			expansion = factorAndReduce(expansion, factors[k])
		}
		yfactor := []float64{points[i].Y / denominator}
		expansion = factorAndReduce(expansion, yfactor)
		numerator[i] = expansion
	}
	return sumPolys(numerator)
}

// pa, pb are assumed to follow the representation of polynomials described in the package notes.
func factorAndReduce(pa, pb []float64) []float64 {
	if len(pa) == 0 {
		return pb
	}
	if len(pb) == 0 {
		return pa
	}
	result := make([]float64, len(pa)+len(pb)-1)
	for i, a := range pa {
		for j, b := range pb {
			result[i+j] += a * b
		}
	}
	return result
}

// per comments for factorAndReduce()
func sumPolys(polys [][]float64) []float64 {
	switch len(polys) {
	case 0:
		return []float64{}
	case 1:
		return polys[0]
	}
	l := len(polys[0])
	for i := 1; i < len(polys); i++ {
		l1 := len(polys[i])
		if l1 > l {
			l = l1
		}
	}
	result := make([]float64, l)
	for _, p := range polys {
		for i, t := range p {
			result[i] += t
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
