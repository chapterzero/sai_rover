package rover

const (
	CMD_ROTATE_LEFT  = 'L'
	CMD_ROTATE_RIGHT = 'R'
	CMD_MOVE         = 'M'
)

func isValidCmd(cmd string) bool {
	for _, c := range cmd {
		if c != CMD_MOVE && c != CMD_ROTATE_LEFT && c != CMD_ROTATE_RIGHT {
			return false
		}
	}
	return true
}
