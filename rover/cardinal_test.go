package rover

import (
	"testing"
)

func TestTranslateMoveToXY(t *testing.T) {
	testCases := []struct {
		name       string
		c          Cardinal
		expected_x int
		expected_y int
	}{
		{
			name:       "From north",
			c:          North,
			expected_x: 0,
			expected_y: 1,
		},
		{
			name:       "From east",
			c:          East,
			expected_x: 1,
			expected_y: 0,
		},
		{
			name:       "From South",
			c:          South,
			expected_x: 0,
			expected_y: -1,
		},
		{
			name:       "From West",
			c:          West,
			expected_x: -1,
			expected_y: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act_x, act_y := tc.c.TranslateMoveToXY()
			if act_x != tc.expected_x {
				t.Errorf("Error when asserting x, expected %d, got %d", tc.expected_x, act_x)
			}

			if act_y != tc.expected_y {
				t.Errorf("Error when asserting y, expected %d, got %d", tc.expected_y, act_y)
			}
		})
	}
}

func TestCardinalRotateLeft(t *testing.T) {
	testCases := []struct {
		name     string
		from     Cardinal
		expected Cardinal
	}{
		{
			name:     "Rotate left from north",
			from:     North,
			expected: West,
		},
		{
			name:     "Rotate left from east",
			from:     East,
			expected: North,
		},
		{
			name:     "Rotate left from south",
			from:     South,
			expected: East,
		},
		{
			name:     "Rotate left from west",
			from:     West,
			expected: South,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.from.RotateLeft()
			if actual != tc.expected {
				t.Errorf("Expected %d got %d", tc.expected, actual)
			}
		})
	}
}

func TestCardinalRotateRight(t *testing.T) {
	testCases := []struct {
		name     string
		from     Cardinal
		expected Cardinal
	}{
		{
			name:     "Rotate right from north",
			from:     North,
			expected: East,
		},
		{
			name:     "Rotate right from east",
			from:     East,
			expected: South,
		},
		{
			name:     "Rotate right from south",
			from:     South,
			expected: West,
		},
		{
			name:     "Rotate right from west",
			from:     West,
			expected: North,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.from.RotateRight()
			if actual != tc.expected {
				t.Errorf("Expected %d got %d", tc.expected, actual)
			}
		})
	}
}

func TestCardinalStr(t *testing.T) {
	testCases := []struct {
		name     string
		c        Cardinal
		expected string
	}{
		{
			name:     "North",
			c:        North,
			expected: "N",
		},
		{
			name:     "East",
			c:        East,
			expected: "E",
		},
		{
			name:     "South",
			c:        South,
			expected: "S",
		},
		{
			name:     "West",
			c:        West,
			expected: "W",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.c.Str() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, tc.c.Str())
			}
		})
	}
}
