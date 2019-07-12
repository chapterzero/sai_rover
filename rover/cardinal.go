package rover

const (
	North Cardinal = iota
	East
	South
	West
)

type Cardinal int

func (c *Cardinal) RotateLeft() {
	switch c {
	case North:
		c = West
	case East:
		c = North
	case South:
		c = East
	case West:
		c = South
	}
}

func (c *Cardinal) RotateRight() {
}
