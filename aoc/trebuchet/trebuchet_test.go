package trebuchet

import "testing"

type TestCase2 struct {
	name        string
	input       string
	want        string
	wanti       int
	ok          bool
	index       int
	leftToRight bool
}

func TestConcatOuterNumbers(t *testing.T) {
	testcases := []TestCase2{
		{
			name:  "Testcase 1",
			input: "two1nine",
			wanti: 29,
		},
		{
			name:  "Testcase 2",
			input: "eightwothree",
			wanti: 83,
		},
		{
			name:  "Testcase 3",
			input: "abcone2threexyz",
			wanti: 13,
		},
		{
			name:  "Testcase 4",
			input: "xtwone3four",
			wanti: 24,
		},
		{
			name:  "Testcase 5",
			input: "4nineeightseven2",
			wanti: 42,
		},
		{
			name:  "Testcase 6",
			input: "zoneight234",
			wanti: 14,
		},
		{
			name:  "Testcase 5",
			input: "7pqrstsixteen",
			wanti: 76,
		},
		{
			name:  "Testcase 6",
			input: "twojrfgzjstjvlngnpqcj8grtnvxmzbxkfrvqmvnmvmbdr",
			wanti: 28,
		},
	}

	for _, testcase := range testcases {
		t.Log(testcase.name)
		got := ConcatOuterNumbers(testcase.input)
		if got != testcase.wanti {
			t.Errorf("wanted %d but got %d", testcase.wanti, got)
		}

	}
}
func TestCheckCharForNumer(t *testing.T) {
	testcases := []TestCase2{
		{
			name:        "Testcase 1",
			input:       "two1nine",
			index:       0,
			ok:          true,
			want:        "2",
			leftToRight: true,
		},
		{
			name:        "Testcase 2",
			input:       "two1nine",
			index:       4,
			ok:          true,
			want:        "9",
			leftToRight: true,
		},
		{
			name:        "Testcase 3",
			input:       "two1nine",
			index:       1,
			ok:          false,
			want:        "",
			leftToRight: true,
		},
		{
			name:        "Testcase 4",
			input:       "two1nine",
			index:       3,
			ok:          true,
			want:        "1",
			leftToRight: true,
		},
		{
			name:        "Testcase 4",
			input:       "two1nine",
			index:       7,
			ok:          true,
			want:        "9",
			leftToRight: false,
		},
	}
	for _, testcase := range testcases {
		t.Log(testcase.name)
		ok, got := CheckCharForNumber(testcase.input, testcase.index, testcase.leftToRight)
		if testcase.ok != ok {
			t.Errorf("ok result should be %t but it's %t", testcase.ok, ok)
		}
		if testcase.want != got {
			t.Errorf("result should be %q but it's %q", testcase.want, got)
		}
	}
}
