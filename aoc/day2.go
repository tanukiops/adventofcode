package aoc

var lengths map[string][]int = map[string][]int{
	"o": {3},
	"t": {3, 5},
	"f": {4},
	"s": {3, 5},
	"e": {5},
	"n": {4},
}
var numbers map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func CheckCharForNumber(s string, idx int) (next int, ok bool, result string) {
	//default return cases
	ok = false
	next = idx + 1
	result = ""
	letter := string(s[idx])
	// if s[index] is a number
	if IsNumber(letter) {
		ok = true
		result = letter
		return
	}
	return
}
