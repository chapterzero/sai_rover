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
	_ = []testCase{
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
			expected_x: 1,
			expected_y: 2,
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
}
