package fizzBuzz

import (
	"strconv"
	"strings"
)

func fizzBuzz(start, end, firstDivisor, secondDivisor int) string {
	var result strings.Builder

	for i := start; i <= end; i++ {
		if i % (firstDivisor * secondDivisor) == 0 {
			result.WriteString("FizzBuzz ")
		} else if i % firstDivisor == 0 {
			result.WriteString("Fizz ")
		} else if i % secondDivisor == 0 {
			result.WriteString("Buzz ")
		} else {
			result.WriteString(strconv.Itoa(i))
			result.WriteString(" ") 
		}
	}

	return strings.Trim(result.String(), " ")
}
