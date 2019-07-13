package helper

import (
	"testing"

	"github.com/chapterzero/sai_rover/rover"
)

func TestParseMaxCoord(t *testing.T) {
	testCases := []struct {
		name               string
		input              string
		expectedX          int
		expectedY          int
		expectedErrMessage string
	}{
		{
			name:               "Empty string",
			input:              "",
			expectedX:          0,
			expectedY:          0,
			expectedErrMessage: "Wrong format",
		},
		{
			name:               "Incomplete input",
			input:              "1",
			expectedX:          0,
			expectedY:          0,
			expectedErrMessage: "Wrong format",
		},
		{
			name:               "X Invalid integer",
			input:              "a 1",
			expectedX:          0,
			expectedY:          0,
			expectedErrMessage: "strconv.Atoi: parsing \"a\": invalid syntax",
		},
		{
			name:               "Y Invalid integer",
			input:              "1 b",
			expectedX:          0,
			expectedY:          0,
			expectedErrMessage: "strconv.Atoi: parsing \"b\": invalid syntax",
		},
		{
			name:               "X negative",
			input:              "-4 1",
			expectedX:          0,
			expectedY:          0,
			expectedErrMessage: "Input must be positive integer",
		},
		{
			name:               "Y negative",
			input:              "1 -2",
			expectedX:          0,
			expectedY:          0,
			expectedErrMessage: "Input must be positive integer",
		},
		{
			name:               "Valid input",
			input:              "3 2",
			expectedX:          3,
			expectedY:          2,
			expectedErrMessage: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act_x, act_y, err := ParseMaxCoord(tc.input)
			if act_x != tc.expectedX {
				t.Errorf("Error when asserting x, expected %d, got %d", tc.expectedX, act_x)
			}
			if act_y != tc.expectedY {
				t.Errorf("Error when asserting y, expected %d, got %d", tc.expectedY, act_y)
			}

			if tc.expectedErrMessage != "" {
				if err.Error() != tc.expectedErrMessage {
					t.Errorf("Error when asserting error, expected '%s' got '%s'", tc.expectedErrMessage, err.Error())
				}
			}
		})
	}
}

func TestParseInitialPos(t *testing.T) {
	testCases := []struct {
		name               string
		input              string
		expectedX          int
		expectedY          int
		expectedCardinal   rover.Cardinal
		expectedErrMessage string
	}{
		{
			name:               "Empty string",
			input:              "",
			expectedX:          0,
			expectedY:          0,
			expectedCardinal:   rover.Cardinal(-1),
			expectedErrMessage: "Wrong format",
		},
		{
			name:               "Incomplete input",
			input:              "1",
			expectedX:          0,
			expectedY:          0,
			expectedCardinal:   rover.Cardinal(-1),
			expectedErrMessage: "Wrong format",
		},
		{
			name:               "Invalid cardinal",
			input:              "1 1 X",
			expectedX:          0,
			expectedY:          0,
			expectedCardinal:   rover.Cardinal(-1),
			expectedErrMessage: "X is not valid cardinal",
		},
		{
			name:               "X Invalid integer",
			input:              "a 1 N",
			expectedX:          0,
			expectedY:          0,
			expectedCardinal:   rover.North,
			expectedErrMessage: "strconv.Atoi: parsing \"a\": invalid syntax",
		},
		{
			name:               "Y Invalid integer",
			input:              "1 b N",
			expectedX:          0,
			expectedY:          0,
			expectedCardinal:   rover.North,
			expectedErrMessage: "strconv.Atoi: parsing \"b\": invalid syntax",
		},
		{
			name:               "Valid input",
			input:              "3 2 E",
			expectedX:          3,
			expectedY:          2,
			expectedCardinal:   rover.East,
			expectedErrMessage: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act_x, act_y, act_c, err := ParseInitialPos(tc.input)
			if act_x != tc.expectedX {
				t.Errorf("Error when asserting x, expected %d, got %d", tc.expectedX, act_x)
			}
			if act_y != tc.expectedY {
				t.Errorf("Error when asserting y, expected %d, got %d", tc.expectedY, act_y)
			}
			if act_c != tc.expectedCardinal {
				t.Errorf(
					"Error when asserting cardinal, expected %s, got %s",
					tc.expectedCardinal.Str(),
					act_c.Str(),
				)
			}

			if tc.expectedErrMessage != "" {
				if err.Error() != tc.expectedErrMessage {
					t.Errorf("Error when asserting error, expected '%s' got '%s'", tc.expectedErrMessage, err.Error())
				}
			}
		})
	}
}
