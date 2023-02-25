package server

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{n: 11, expected: true},
		{n: 15, expected: false},
		{n: 2, expected: true},
		{n: 1, expected: false},
		{n: 29, expected: true},
		{n: -5, expected: false},
		{n: 0, expected: false},
		{n: -1, expected: false},
		{n: -7, expected: false},
		{n: 19 * 41, expected: false},
	}

	for i, testCase := range tests {
		t.Run(fmt.Sprintf("Test_№%d", i+1), func(t *testing.T) {
			require.Equal(t, testCase.expected, isPrime(testCase.n))
		})
	}
}

func TestGetPrimes(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected Response
	}{
		{
			numbers:  []int{1, 5, 7, 11, 15},
			expected: Response{false, true, true, true, false},
		},
		{
			numbers:  []int{-7, -16542, -66, 0, -3},
			expected: Response{false, false, false, false, false},
		},
		{
			numbers:  []int{-11, math.MaxInt64, 8, 41, 37},
			expected: Response{false, false, false, true, true},
		},
	}

	for i, testCase := range tests {
		t.Run(fmt.Sprintf("Test_№%d", i+1), func(t *testing.T) {
			require.Equal(t, testCase.expected, getPrimes(testCase.numbers))
		})
	}
}
