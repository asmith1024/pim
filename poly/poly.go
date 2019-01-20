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
	return fromPoints(points)
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
func fromPoints(points []Point) []float64 {
	pterms := make([][]float64, len(points))
	for i := range points {
		term := termAtIndex(points, i)
		expansion := term[0]
		for k := 1; k < len(term); k++ {
			expansion = product(expansion, term[k])
		}
		pterms[i] = expansion
	}
	return sum(pterms)
}

func termAtIndex(points []Point, i int) [][]float64 {
	factors := make([][]float64, len(points))
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
	factors[fidx] = []float64{points[i].Y / denominator}
	return factors
}

// pa, pb are assumed to follow the representation of polynomials described in the package notes.
// Keep in mind that we are multiplying ax^m by bx^n so we end up with abx^(n+m).
func product(pa, pb []float64) []float64 {
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

// see comments for multPoly()
func sum(polys [][]float64) []float64 {
	switch len(polys) {
	case 0:
		return []float64{}
	case 1:
		return polys[0]
	}
	result := makePolyForSum(polys)
	for _, p := range polys {
		for i, t := range p {
			result[i] += t
		}
	}
	return result
}

func makePolyForSum(polys [][]float64) []float64 {
	l := len(polys[0])
	for i := 1; i < len(polys); i++ {
		l1 := len(polys[i])
		if l1 > l {
			l = l1
		}
	}
	return make([]float64, l)
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
