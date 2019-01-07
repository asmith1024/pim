package poly

import "testing"

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
