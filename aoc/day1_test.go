package aoc

import "testing"

type TestCase struct {
	input string
	want  int
	index int
}

func TestConcatOuterNumbers(t *testing.T) {
	testcases := []TestCase{
		{
			input: "1abc2",
			want:  12,
		},
		{
			input: "pqr3stu8vwx",
			want:  38,
		},
		{
			input: "a1b2c3d4e5f",
			want:  15,
		},
		{
			input: "treb7uchet",
			want:  77,
		},
	}
	for _, testcase := range testcases {
		got := ConcatOuterNumbers(testcase.input)
		if got != testcase.want {
			t.Errorf("got %d but wanted %d for input %q", got, testcase.want, testcase.input)
		}

	}

}
