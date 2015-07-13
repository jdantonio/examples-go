package mymath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqrt(t *testing.T) {
	assert.InDelta(t, 1.414213562373095, Sqrt(2), 0.0000000001)
}

func TestMaxFirstIsLower(t *testing.T) {
	assert.Equal(t, 100, Max(1, 100))
}

func TestMaxFirstIsHigher(t *testing.T) {
	assert.Equal(t, 100, Max(100, 1))
}

func TestSumAndProduct(t *testing.T) {
	expected_sum := 7
	expected_product := 12

	sum, product := SumAndProduct(3, 4)

	assert.Equal(t, expected_sum, sum)
	assert.Equal(t, expected_product, product)
}

func TestSum(t *testing.T) {
	assert.Equal(t, 1, Sum(1))
	assert.Equal(t, 55, Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
