package gearratios

import (
	"fmt"
	"strconv"

	"github.com/tanukiops/adventofcode/aoc/trebuchet"
)

type DiscoveredNumber struct {
	Indexes []int
	Line    int
	Part    bool
	Number  int
}

func SumOfParts(lines []string) (sum int) {
	numbers := []DiscoveredNumber{}
	length := len(lines)
	for i, line := range lines {
		numbers = discoverNumbers(numbers, line, i)
	}
	for _, number := range numbers {
		if checkIfPart(number, lines, length) {
			sum += number.Number
		}
	}
	fmt.Println(numbers)
	return
}

func checkIfPart(number DiscoveredNumber, lines []string, length int) bool {
	fmt.Println(length)
	fmt.Println(number.Line)
	var left int
	var right int
	indexesLenght := len(number.Indexes)
	if number.Indexes[0] > 0 {
		left = number.Indexes[0] - 1
	} else {
		left = number.Indexes[0]
	}
	if number.Indexes[indexesLenght-1] < length-1 {
		right = number.Indexes[indexesLenght-1] + 1
	} else {
		right = number.Indexes[indexesLenght-1]
	}
	for i := left; i <= right; i++ {
		if isSymbol(string(lines[number.Line][i])) {
			return true
		}
		if number.Line > 1 {
			if isSymbol(string(lines[number.Line-1][i])) {
				return true
			}
		}
		if number.Line < length-1 {
			if isSymbol(string(lines[number.Line+1][i])) {
				return true
			}
		}
	}
	return false
}

func isSymbol(symbol string) bool {
	return (symbol != "." && !trebuchet.IsNumber(symbol))
}

func discoverNumbers(numbers []DiscoveredNumber, line string, lineIndex int) []DiscoveredNumber {
	potentialNumber := ""
	indexes := []int{}
	processingNumber := false
	for i, char := range line {
		if trebuchet.IsNumber(string(char)) {
			processingNumber = true
			potentialNumber += string(char)
			indexes = append(indexes, i)
		}
		// if char is not a number, it means it's either a symbol or a dot, meaning that the contents of potentialNumber is a number
		if (!trebuchet.IsNumber(string(char)) || i == len(line)-1) && processingNumber {
			// we're not processing a number anymore
			processingNumber = false
			number := DiscoveredNumber{}
			number.Number, _ = strconv.Atoi(potentialNumber)
			number.Indexes = indexes
			number.Line = lineIndex
			numbers = append(numbers, number)
			potentialNumber = ""
			indexes = []int{}
		}
	}
	return numbers
}
