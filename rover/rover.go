package rover

func New(
	start_x int,
	start_y int,
	max_x int,
	max_y int,
	start_cardinal Cardinal,
) *Rover {
	return &Rover{
		x:     start_x,
		y:     start_y,
		max_x: max_x,
		max_y: max_y,
		c:     start_cardinal,
	}
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
// command will be validated as whole
// return error if any command is not recognized
func (r *Rover) Move(cmd string) error {
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

func (r *Rover) Pos() (int, int) {
	return r.x, r.y
}
