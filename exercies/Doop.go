package piscine

import (
	"fmt"
)

func Doop(arg0 int, arg1 string, arg2 int) int {
	maxint := 9223372036854775807
	minint := -9223372036854775808
	result := 0

	if arg1 == "+" {
		result = arg0 + arg2
	}

	if arg1 == "-" {
		result = arg0 - arg2
	}

	if arg1 == "*" {
		result = arg0 * arg2

	}
	if arg1 == "/" {
		if arg2 == 0 {
			fmt.Println("No division by 0")
		}
		result = arg0 / arg2

	}
	if arg1 == "%" {
		if arg2 == 0 {
			fmt.Println("No Modulo by 0")
		}
		result = arg0 % arg2
	}

	if arg0+arg2 >= maxint || arg0+arg2 <= minint || arg0-arg2 >= maxint || arg0-arg2 <= minint || arg0*arg2 >= maxint || arg0*arg2 <= minint {
		result = 0
	}
	return result
}
