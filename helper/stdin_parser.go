package helper

import (
	"errors"
	"strconv"
	"strings"

	"github.com/chapterzero/sai_rover/rover"
)

func ParseMaxCoord(in string) (int, int, error) {
	max_s := strings.Split(in, " ")
	if len(max_s) != 2 {
		return 0, 0, errors.New("Wrong format")
	}

	max_x, err := strconv.Atoi(max_s[0])
	if err != nil {
		return 0, 0, err
	}

	max_y, err := strconv.Atoi(max_s[1])
	if err != nil {
		return 0, 0, err
	}

	if max_x < 0 || max_y < 0 {
		return 0, 0, errors.New("Input must be positive integer")
	}

	return max_x, max_y, nil
}

func ParseInitialPos(in string) (int, int, rover.Cardinal, error) {
	pos := strings.Split(in, " ")
	if len(pos) != 3 {
		return 0, 0, rover.Cardinal(-1), errors.New("Wrong format")
	}

	c, err := rover.NewCardinal(pos[2])
	if err != nil {
		return 0, 0, c, err
	}

	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return 0, 0, c, err
	}

	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return 0, 0, c, err
	}

	return x, y, c, nil
}
