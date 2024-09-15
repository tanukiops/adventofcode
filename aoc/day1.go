package aoc

import "strconv"

func TrebuchetCalibration(lines []string) (sum int) {
	sum = 0
	for _, line := range lines {
		sum += ConcatOuterNumbers(line)
	}
	return
}

func ConcatOuterNumbers(s string) (number int) {
	length := len(s)
	left := ""
	right := ""
	for i, j := 0, length-1; i < length; i++ {
		if left == "" && IsNumber(string(s[i])) {
			left = string(s[i])
		}
		if right == "" && IsNumber(string(s[j])) {
			right = string(s[j])
		}
		j--
	}

	number, err := strconv.Atoi(left + right)
	if err != nil {
		return 0
	}
	return number
}

func IsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}
