package rover

import (
	"fmt"
)

// Create new rover based on initial settings
// initial position must be valid: inside boundary
func New(
	start_x int,
	start_y int,
	max_x int,
	max_y int,
	start_cardinal Cardinal,
) (*Rover, error) {
	if start_x > max_x || start_y > max_y || start_x < 0 || start_y < 0 {
		return nil, fmt.Errorf("Invalid rover parameter")
	}

	return &Rover{
		x:     start_x,
		y:     start_y,
		max_x: max_x,
		max_y: max_y,
		c:     start_cardinal,
	}, nil
}

type Rover struct {
	x     int
	y     int
	max_x int
	max_y int
	c     Cardinal
}

// move the rover using command
// ex: LMLLMRMLM
// command will be validated as whole (not executing partially)
// return error if command is not recognized
func (r *Rover) Move(cmd string) error {
	if !isValidCmd(cmd) {
		return fmt.Errorf("%s is invalid command", cmd)
	}

	for _, c := range cmd {
		switch c {
		case 'M':
			r.move(r.c.TranslateMoveToXY())
		case 'L':
			r.c = r.c.RotateLeft()
		case 'R':
			r.c = r.c.RotateRight()
		}
	}

	return nil
}

// actual move
func (r *Rover) move(x, y int) {
	r.x += x
	r.y += y

	if r.x > r.max_x {
		r.x = r.max_x
	}
	if r.x < 0 {
		r.x = 0
	}

	if r.y > r.max_y {
		r.y = r.max_y
	}
	if r.y < 0 {
		r.y = 0
	}
}

func (r *Rover) Pos() (int, int, Cardinal) {
	return r.x, r.y, r.c
}
