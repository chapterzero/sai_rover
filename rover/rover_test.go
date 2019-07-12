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
			name:       "Test single move facing north",
			r:          New(1, 1, 3, 3, North),
			command:    "M",
			expected_x: 1,
			expected_y: 2,
			expected_c: North,
		},
		testCase{
			name:       "Test single move facing east",
			r:          New(1, 1, 3, 3, East),
			command:    "M",
			expected_x: 2,
			expected_y: 1,
			expected_c: East,
		},
		testCase{
			name:       "Test single move facing south",
			r:          New(1, 1, 3, 3, South),
			command:    "M",
			expected_x: 1,
			expected_y: 0,
			expected_c: South,
		},
		testCase{
			name:       "Test single move facing west",
			r:          New(1, 1, 3, 3, West),
			command:    "M",
			expected_x: 0,
			expected_y: 1,
			expected_c: West,
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

func TestRoverInvalidMove(t *testing.T) {
	r := New(1, 1, 3, 3, North)
	err := r.Move("MMXX")
	if err == nil {
		t.Errorf("Expected error but got nil")
		return
	}
	if err.Error() != "MMXX is invalid command" {
		t.Errorf("Expected error message '%s', got '%s'", "MMXX is invalid command", err.Error())
	}
}
