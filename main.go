package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tanukiops/adventofcode/aoc/trebuchet"
)

func main() {
	fmt.Println("Day 1")
	day1()
}
func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}
func day1() {
	inputFile := "assets/day1.txt"
	lines := []string{}
	file, err := os.Open(inputFile)
	checkErrors(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	result := trebuchet.TrebuchetCalibration(lines)
	fmt.Println(result)
}
