package aoc

import "testing"

type TestCase2 struct {
	input     string
	nextIndex int
	want      string
	ok        bool
	index     int
}

func TestCheckCharForNumer(t *testing.T) {
	testcases := []TestCase2{
		{
			input:     "two1nine",
			index:     0,
			ok:        true,
			want:      "2",
			nextIndex: 4,
		},
		{
			input:     "two1nine",
			index:     4,
			ok:        true,
			want:      "9",
			nextIndex: -1,
		},
		{
			input:     "two1nine",
			index:     1,
			ok:        false,
			want:      "",
			nextIndex: 2,
		},
	}
	for _, testcase := range testcases {
		nextInt, ok, got := CheckCharForNumber(testcase.input, testcase.index)
		if testcase.nextIndex != nextInt {
			t.Errorf("nextInt should be %d but it's %d", testcase.nextIndex, nextInt)
		}
		if testcase.ok != ok {
			t.Errorf("ok result should be %t but it's %t", testcase.ok, ok)
		}
		if testcase.want != got {
			t.Errorf("result should be %q but it's %q", testcase.want, got)
		}
	}
}
