package poly

import (
	"math"
	"math/rand"
	"testing"
)

func TestEmptyCoefficients(t *testing.T) {
	y := Y(1, []float64{})
	if y != 0 {
		t.Error("Expected 0 got", y)
	}
}

func TestZeroX(t *testing.T) {
	y := Y(0, []float64{25.4, 1, 2, 3, 4})
	if y != 25.4 {
		t.Error("Expected 25.4 got", y)
	}
}

func TestPositiveCoeffients(t *testing.T) {
	y := Y(2, []float64{1, 2, 3, 4})
	if y != 49 {
		t.Error("Expected 49 got", y)
	}
}

func TestNegativeCoeffients(t *testing.T) {
	y := Y(2, []float64{-1, -2, -3, -4})
	if y != -49 {
		t.Error("Expected -49 got", y)
	}
}

func TestMixedSignCoefficients(t *testing.T) {
	y := Y(2, []float64{-1, 2, -3, 4})
	if y != 23 {
		t.Error("Expected 23 got", y)
	}
}

func TestNegativeX(t *testing.T) {
	y := Y(-2, []float64{-1, 2, -3, 4})
	if y != -49 {
		t.Error("Expected -49 got", y)
	}
}

func TestFractionalX(t *testing.T) {
	y := Y(0.1, []float64{1, 2, 3})
	if y != 1.23 {
		t.Error("Expected 1.23 got", y)
	}
}

func TestZeroCoefficients(t *testing.T) {
	y := Y(2, []float64{1, 0, 0, 0, 5})
	if y != 81 {
		t.Error("Expected 81 got", y)
	}
}

func TestRandomDegree5(t *testing.T) {
	for i := 0; i < 1000; i++ {
		x := randFloat()
		cs := randCs()
		tval := Y(x, cs[:])
		cval := dumbEval(x, cs)
		if tval != cval {
			t.Error("Expected", cval, "got", tval)
		}
	}
}

func randFloat() float64 {
	abs := rand.Float64() * rand.Float64() * 1000
	if rand.Float32() < 0.5 {
		return abs * -1
	}
	return abs
}

func randCs() [5]float64 {
	return [5]float64{
		randFloat(),
		randFloat(),
		randFloat(),
		randFloat(),
		randFloat(),
	}
}

func dumbEval(x float64, cs [5]float64) float64 {
	return cs[0] + x*cs[1] + math.Pow(x, 2)*cs[2] + math.Pow(x, 3)*cs[3] + math.Pow(x, 4)*cs[4]
}

func TestInterpolate1(t *testing.T) {
	pts := []Point{Point{X: 1, Y: 6}, Point{X: 2, Y: 17}, Point{X: 3, Y: 34}}
	test := Interpolate(pts)
	expected := []float64{1, 2, 3}
	if !testEq(test, expected) {
		t.Error("Unexpected polynomial", test)
	}
}

func testEq(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, c := range a {
		if c != b[i] {
			return false
		}
	}
	return true
}

func TestPrint1(t *testing.T) {
	s := Print([]float64{1, -2, 3, -4})
	e := "f(x) = -4x^3 + 3x^2 - 2x + 1"
	if s != e {
		t.Error("Unexpected rendering", s)
	}
}

func TestPrint2(t *testing.T) {
	s := Print([]float64{1, 0, 0, 0})
	e := "f(x) = 1"
	if s != e {
		t.Error("Unexpected rendering", s)
	}
}

func TestPrint3(t *testing.T) {
	s := Print([]float64{0, 0, 0, -8.123, 0, 0, 0, 0})
	e := "f(x) = -8.123x^3"
	if s != e {
		t.Error("Unexpected rendering", s)
	}
}

func TestPrint4(t *testing.T) {
	s := Print([]float64{-6.444, 0, 0, -8.123, 0, 0, 0, 0})
	e := "f(x) = -8.123x^3 - 6.444"
	if s != e {
		t.Error("Unexpected rendering", s)
	}
}

func TestPrintEmpty(t *testing.T) {
	s := Print([]float64{})
	if s != "" {
		t.Error("Unexpected rendering", s)
	}
}

func TestPrintAllZeroes(t *testing.T) {
	s := Print([]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	if s != "" {
		t.Error("Unexpected rendering", s)
	}
}
