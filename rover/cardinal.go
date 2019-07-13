package rover

const (
	North Cardinal = iota
	East
	South
	West
)

type Cardinal int

// calculate x and y increments
// based on cardinality
func (c Cardinal) TranslateMoveToXY() (int, int) {
	x, y := 0, 0
	switch c {
	case North:
		y = 1
	case East:
		x = 1
	case South:
		y = -1
	case West:
		x = -1
	}

	return x, y
}

func (c Cardinal) RotateLeft() Cardinal {
	switch c {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	}

	return c
}

func (c Cardinal) RotateRight() Cardinal {
	switch c {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}

	return c
}

func (c Cardinal) Str() string {
	switch c {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	}

	return ""
}
