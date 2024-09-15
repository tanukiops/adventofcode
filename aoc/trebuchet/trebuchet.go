package trebuchet

import "strconv"

var lengths map[string][]int = map[string][]int{
	"o": {3},
	"t": {3, 5},
	"f": {4},
	"x": {3},
	"r": {4},
	"s": {3, 5},
	"e": {5, 4, 3},
	"n": {4, 5},
}
var numbers map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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
		leftOk, leftResult := CheckCharForNumber(s, i, true)
		rightOk, rightResult := CheckCharForNumber(s, j, false)
		if left == "" && leftOk {
			left = leftResult
		}
		if right == "" && rightOk {
			right = rightResult
		}
		if left != "" && right != "" {
			break
		}
		j--
	}

	number, err := strconv.Atoi(left + right)
	if err != nil {
		return 0
	}
	return number
}

func CheckCharForNumber(s string, idx int, rightToLeft bool) (ok bool, result string) {
	//default return cases
	ok = false
	result = ""
	letter := string(s[idx])
	// if s[index] is a number
	if IsNumber(letter) {
		ok = true
		result = letter
		return
	}
	potentialLengths, ok := lengths[letter]
	if ok {
		for _, length := range potentialLengths {
			if rightToLeft && idx+length < len(s) {
				result, ok = numbers[s[idx:idx+length]]
			} else if idx-length+1 > -1 {
				result, ok = numbers[s[idx-length+1:idx+1]]
			}
			//if string is a key in the map we found a number.
			if ok {
				break
			}
		}
	}
	return
}

func IsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}
