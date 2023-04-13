package fizzBuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		start         int
		end           int
		firstDivisor  int
		secondDivisor int
		expected      string
	}{
		{1, 15, 3, 5, "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz "},
		{15, 30, 2, 7, "15 Fizz 17 Fizz 19 Fizz Buzz Fizz 23 Fizz 25 Fizz 27 FizzBuzz 29 Fizz "},
		{10, 20, 2, 4, "Fizz 11 Fizz 13 Fizz 15 FizzBuzz 17 Fizz 19 Fizz "},
	}

	for _, tc := range testCases {
		result := fizzBuzz(tc.start, tc.end, tc.firstDivisor, tc.secondDivisor)
		if result != tc.expected {
			t.Errorf("fizzBuzz(%d, %d, %d, %d) = %s, expected %s", tc.start, tc.end, tc.firstDivisor, tc.secondDivisor, result, tc.expected)
		}
	}
}
