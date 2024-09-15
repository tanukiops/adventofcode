package cubeconundrum

import "testing"

type TestGameTestCase struct {
	grab map[string]int
	want bool
}

type PlayGameTestCase struct {
	grabs []map[string]int
	want  bool
}

type PlayGamesTestCase struct {
	games [][]map[string]int
	want  int
}

var possibleGrab map[string]int = map[string]int{
	"red":   5,
	"green": 4,
	"blue":  13,
}
var impossibleGrab map[string]int = map[string]int{
	"red":   51,
	"green": 4,
	"blue":  13,
}

func TestPlayGames(t *testing.T) {
	testcases := []PlayGamesTestCase{
		{
			want: 3,
			games: [][]map[string]int{
				{
					possibleGrab,
					possibleGrab,
				},
				{
					possibleGrab,
					possibleGrab,
				},
			},
		},
		{
			want: 5,
			games: [][]map[string]int{
				{
					possibleGrab,
					possibleGrab,
				},
				{
					impossibleGrab,
					possibleGrab,
				},
				{
					impossibleGrab,
					possibleGrab,
				},
				{
					possibleGrab,
					possibleGrab,
				},
			},
		},
	}
	for i, testcase := range testcases {
		t.Logf("TestCase %d", i)
		got := PlayGames(testcase.games)
		if testcase.want != got {
			t.Errorf("Wanted %d but got %d", testcase.want, got)
		}
	}
}
func TestPlayGame(t *testing.T) {
	testcases := []PlayGameTestCase{
		{
			want: true,
			grabs: []map[string]int{
				possibleGrab,
				possibleGrab,
			},
		},
		{
			want: false,
			grabs: []map[string]int{
				impossibleGrab,
				possibleGrab,
			},
		},
	}

	for i, testcase := range testcases {
		t.Logf("TestCase %d", i)
		got := PlayGame(testcase.grabs)
		if testcase.want != got {
			t.Errorf("Wanted %t but got %t", testcase.want, got)
		}
	}
}
func TestTestGame(t *testing.T) {
	testcases := []TestGameTestCase{
		{
			want: true,
			grab: map[string]int{
				"red":   5,
				"green": 4,
				"blue":  13,
			},
		},
		{
			want: false,
			grab: map[string]int{
				"red":   15,
				"green": 4,
				"blue":  1,
			},
		},
	}

	for i, testcase := range testcases {
		t.Logf("TestCase %d", i)
		got := TestGame(testcase.grab)
		if testcase.want != got {
			t.Errorf("Wanted %t but got %t", testcase.want, got)
		}
	}
}
