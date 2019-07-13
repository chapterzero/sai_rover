package rover

import (
	"testing"
)

type testCase struct {
	name       string
	r          *Rover
	command    string
	expected_x int
	expected_y int
	expected_c Cardinal
}

func TestRoverSingleMoveCmd(t *testing.T) {
	testCases := []testCase{
		testCase{
			name: "Test single move facing north",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     North,
			},
			command:    "M",
			expected_x: 1,
			expected_y: 2,
			expected_c: North,
		},
		testCase{
			name: "Test single move facing east",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     East,
			},
			command:    "M",
			expected_x: 2,
			expected_y: 1,
			expected_c: East,
		},
		testCase{
			name: "Test single move facing south",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     South,
			},
			command:    "M",
			expected_x: 1,
			expected_y: 0,
			expected_c: South,
		},
		testCase{
			name: "Test single move facing west",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     West,
			},
			command:    "M",
			expected_x: 0,
			expected_y: 1,
			expected_c: West,
		},
	}

	runTest(t, testCases)
}

func TestRoverRotateAndMoveCmd(t *testing.T) {
	testCases := []testCase{
		testCase{
			name: "From north rotate left and move",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     North,
			},
			command:    "LM",
			expected_x: 0,
			expected_y: 1,
			expected_c: West,
		},
		testCase{
			name: "From East rotate right and move",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     East,
			},
			command:    "RM",
			expected_x: 1,
			expected_y: 0,
			expected_c: South,
		},
		testCase{
			name: "From South rotate left twice and move",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     South,
			},
			command:    "LLM",
			expected_x: 1,
			expected_y: 2,
			expected_c: North,
		},
		testCase{
			name: "From West, rotate left, rotate right x2 and move",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     West,
			},
			command:    "LRRM",
			expected_x: 1,
			expected_y: 2,
			expected_c: North,
		},
		testCase{
			name: "From east, rotate right, move, rotate right",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     East,
			},
			command:    "RMR",
			expected_x: 1,
			expected_y: 0,
			expected_c: West,
		},
		testCase{
			name: "From North, rotate left, move, rotate right move",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 3,
				max_y: 3,
				c:     North,
			},
			command:    "LMRM",
			expected_x: 0,
			expected_y: 2,
			expected_c: North,
		},
	}

	runTest(t, testCases)
}

func TestRoverMoveOutsideBoundary(t *testing.T) {
	testCases := []testCase{
		testCase{
			name: "Hit South Boundary",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 2,
				max_y: 2,
				c:     South,
			},
			command:    "MM",
			expected_x: 1,
			expected_y: 0,
			expected_c: South,
		},
		testCase{
			name: "Rotate, move then Hit East Boundary",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 2,
				max_y: 2,
				c:     North,
			},
			command:    "RMMM",
			expected_x: 2,
			expected_y: 1,
			expected_c: East,
		},
		testCase{
			name: "Rotate and move after hit North boundary",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 2,
				max_y: 2,
				c:     East,
			},
			command:    "LMMMLM",
			expected_x: 0,
			expected_y: 2,
			expected_c: West,
		},
		testCase{
			name: "Reverse after hit west boundary",
			r: &Rover{
				x:     1,
				y:     1,
				max_x: 2,
				max_y: 2,
				c:     North,
			},
			command:    "LMMMMRRM",
			expected_x: 1,
			expected_y: 1,
			expected_c: East,
		},
	}

	runTest(t, testCases)
}

func runTest(t *testing.T, testCases []testCase) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.r.Move(tc.command)
			act_x, act_y, act_c := tc.r.Pos()

			if act_x != tc.expected_x {
				t.Errorf("Error when asserting x pos, expected %d, got %d", tc.expected_x, act_x)
			}

			if act_y != tc.expected_y {
				t.Errorf("Error when asserting y pos, expected %d, got %d", tc.expected_y, act_y)
			}

			if act_c != tc.expected_c {
				t.Errorf("Error when asserting cardinality, expected %s, got %s", tc.expected_c.Str(), act_c.Str())
			}
		})
	}
}

func TestNewRoverInitialOutsideOfBoundary(t *testing.T) {
	testCases := []struct {
		name    string
		start_x int
		start_y int
		max_x   int
		max_y   int
	}{
		{
			name:    "X Outside of max x",
			start_x: 4,
			start_y: 3,
			max_x:   3,
			max_y:   3,
		},
		{
			name:    "Negative x",
			start_x: -1,
			start_y: 3,
			max_x:   3,
			max_y:   3,
		},
		{
			name:    "Y Outside of max y",
			start_x: 3,
			start_y: 4,
			max_x:   3,
			max_y:   3,
		},
		{
			name:    "Negative y",
			start_x: 3,
			start_y: -1,
			max_x:   3,
			max_y:   3,
		},
		{
			name:    "X and Y outside of max x and y",
			start_x: 4,
			start_y: 4,
			max_x:   3,
			max_y:   3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := New(tc.start_x, tc.start_y, tc.max_x, tc.max_y, North)
			if r != nil {
				t.Errorf("Expected rover is not created, got non nil")
			}

			if err == nil {
				t.Errorf("Expected error is not nil, got nil")
				return
			}

			if err.Error() != "Invalid rover parameter" {
				t.Errorf("Expected error message '%s' got '%s'", "Invalid rover parameter", err.Error())
			}
		})
	}
}

func TestRoverInvalidMove(t *testing.T) {
	r, _ := New(1, 1, 3, 3, North)
	err := r.Move("MMXX")
	if err == nil {
		t.Errorf("Expected error but got nil")
		return
	}
	if err.Error() != "MMXX is invalid command" {
		t.Errorf("Expected error message '%s', got '%s'", "MMXX is invalid command", err.Error())
	}
}
