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
}
