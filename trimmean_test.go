package trimmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 5.5 // With 10% trimming
	result := TrimmedMean(data, 0.1)
	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestRound(t *testing.T) {
	result := Round(3.14159265, 4)
	expected := 3.1416
	if result != expected {
		t.Errorf("Expected %.4f, got %.4f", expected, result)
	}
}
