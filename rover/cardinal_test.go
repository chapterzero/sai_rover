package rover

import (
	"testing"
)

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
			actual := RotateLeft(tc.from)
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
			actual := RotateRight(tc.from)
			if actual != tc.expected {
				t.Errorf("Expected %d got %d", tc.expected, actual)
			}
		})
	}
}
