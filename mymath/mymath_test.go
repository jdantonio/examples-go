package mymath

import "testing"

func TestSqrt(t *testing.T) {

	lower := 1.414213562373091
	upper := 1.414213562373099

	if subject := Sqrt(2); subject < lower || subject > upper {
		t.Error("Expected 1.414213562373095, got", subject)
	}
}

func TestMaxFirstIsLower(t *testing.T) {
	if subject := Max(1, 100); subject != 100 {
		t.Error("Expected 100, got", subject)
	}
}

func TestMaxFirstIsHigher(t *testing.T) {
	if subject := Max(100, 1); subject != 100 {
		t.Error("Expected 100, got", subject)
	}
}

func TestSumAndProduct(t *testing.T) {
	sum := 7
	product := 12

	if subject, _ := SumAndProduct(3, 4); subject != sum {
		t.Error("Expected sum 7, got", subject)
	}

	if _, subject := SumAndProduct(3, 4); subject != product {
		t.Error("Expected product 12, got", subject)
	}
}

func TestSum(t *testing.T) {
	if subject := Sum(1); subject != 1 {
		t.Error("Expected sum of 1 to be 1, got", subject)
	}

	if subject := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10); subject != 55 {
		t.Error("Expected sum of 1-10 to be 55, got", subject)
	}
}
