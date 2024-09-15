package cubeconundrum

import (
	"fmt"
	"strconv"
	"strings"
)

var cubes map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// tests if a game is possible
func TestGame(grab map[string]int) bool {
	for k, v := range grab {
		if cubes[k]-v < 0 {
			return false
		}
	}
	return true
}

// ParseGame takes in a string structures as "Game ID: color: amount, ...; ..."
// and returns a list of maps with their respective colors and amounts
func ParseGame(s string) (grabs []map[string]int) {
	withoutPrefix := s[strings.Index(s, ":")+2:]
	rawGrabs := strings.Split(withoutPrefix, ";")
	for _, s := range rawGrabs {
		rawColors := strings.Split(s, ",")
		grab := make(map[string]int)
		for _, color := range rawColors {
			colorValue := strings.Split(strings.Trim(color, " "), " ")
			grab[colorValue[1]], _ = strconv.Atoi(colorValue[0])

		}
		grabs = append(grabs, grab)
	}
	fmt.Println(grabs)
	return
}

// Iterates over all the grabs of a game and test if those are possible
// returns false when one of the grabs would be impossible.
func PlayGame(grabs []map[string]int) (possible bool) {
	possible = true
	for _, grab := range grabs {
		possible = TestGame(grab)
		if !possible {
			break
		}
	}
	return
}

// Iterate over all the games
// If the game is possible, add gameid to the sum
func PlayGames(games [][]map[string]int) (sum int) {
	for i, game := range games {
		id := i + 1
		if PlayGame(game) {
			sum += id
		}
	}
	return
}
func SumOfPower(games [][]map[string]int) (sum int) {
	for _, game := range games {
		sum += PowerOfGame(game)
	}
	return
}
func PowerOfGame(grabs []map[string]int) (power int) {
	colors := make(map[string]int)
	colors["red"] = 0
	colors["blue"] = 0
	colors["green"] = 0
	for _, grab := range grabs {
		if grab["red"] > colors["red"] {
			colors["red"] = grab["red"]
		}
		if grab["green"] > colors["green"] {
			colors["green"] = grab["green"]
		}
		if grab["blue"] > colors["blue"] {
			colors["blue"] = grab["blue"]
		}
	}
	return colors["red"] * colors["green"] * colors["blue"]
}

func ParseGames(rawGames []string) (games [][]map[string]int) {
	for _, rawGame := range rawGames {
		games = append(games, ParseGame(rawGame))
	}
	return
}
