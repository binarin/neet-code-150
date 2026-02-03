package main

import (
	"math"
	"strconv"
)

func main() {
	reverse(123)
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	var sign int64 = 1

	if x < 0 {
		sign = -1
		x = -x
	}

	digits := strconv.Itoa(x)
	result := int64(0)
	last_idx := len(digits) - 1
	for i := last_idx; i >= 0; i-- {
		result = result*10 + int64(digits[i]) - int64('0')
		if result > math.MaxInt32 {
			// overflow
			return 0
		}
	}

	result *= sign // negative range is bigger, so no overflow here
	return int(result)
}
