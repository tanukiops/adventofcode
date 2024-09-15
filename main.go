package main

import (
	"bufio"
	"fmt"
	"os"

	cubeconundrum "github.com/tanukiops/adventofcode/aoc/cube_conundrum"
	"github.com/tanukiops/adventofcode/aoc/trebuchet"
)

func main() {
	fmt.Println("Day 1")
	day1()
	fmt.Println("Day 2")
	day2()
}
func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}
func day2() {
	inputFile := "assets/day2.txt"
	lines := fileToStrings(inputFile)
	games := cubeconundrum.ParseGames(lines)
	sum := cubeconundrum.PlayGames(games)
	power := cubeconundrum.SumOfPower(games)
	fmt.Println(sum)
	fmt.Println(power)
}
func day1() {
	inputFile := "assets/day1.txt"
	lines := fileToStrings(inputFile)
	result := trebuchet.TrebuchetCalibration(lines)
	fmt.Println(result)
}

func fileToStrings(filename string) (strings []string) {
	file, err := os.Open(filename)
	checkErrors(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	return
}
