package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tanukiops/adventofcode/aoc"
)

func main() {
	printResult()
}
func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}
func printResult() {
	inputFile := "assets/day1.txt"
	lines := []string{}
	file, err := os.Open(inputFile)
	checkErrors(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	result := aoc.TrebuchetCalibration(lines)
	fmt.Println(result)
}
