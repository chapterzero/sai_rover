package rover

const (
	North Cardinal = iota
	East
	South
	West
)

type Cardinal int

func RotateLeft(c Cardinal) Cardinal {
	switch c {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	default:
		return South
	}
}

func RotateRight(c Cardinal) Cardinal {
	switch c {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	default:
		return North
	}
}
