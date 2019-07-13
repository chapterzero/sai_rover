package helper

import (
	"errors"
	"strconv"
	"strings"

	"github.com/chapterzero/sai_rover/rover"
)

func ParseMaxCoord(in string) (int, int, error) {
	maxS := strings.Split(in, " ")
	if len(maxS) != 2 {
		return 0, 0, errors.New("Wrong format")
	}

	maxX, err := strconv.Atoi(maxS[0])
	if err != nil {
		return 0, 0, err
	}

	maxY, err := strconv.Atoi(maxS[1])
	if err != nil {
		return 0, 0, err
	}

	if maxX < 0 || maxY < 0 {
		return 0, 0, errors.New("Input must be positive integer")
	}

	return maxX, maxY, nil
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
