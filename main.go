package main

import (
	"bufio"
	"log"
	"os"

	"github.com/chapterzero/sai_rover/helper"
	"github.com/chapterzero/sai_rover/rover"
)

func main() {
	log.Println("MARS ROVER PROGRAM v 0.1 press CTRL-C to exit")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// BOUNDARY settings
		log.Println("Enter upper right coordinates of plateau(x y), example: 5 5")
		scanner.Scan()
		max_x, max_y, err := helper.ParseMaxCoord(scanner.Text())
		if err != nil {
			print_error(err)
			continue
		}

		for {
			// ROVER creation
			log.Println("Create new ROVER. Enter initial position(start_x start_y cardinal_direction), example: 1 2 N")
			scanner.Scan()
			x, y, c, err := helper.ParseInitialPos(scanner.Text())
			if err != nil {
				print_error(err)
				continue
			}
			r, err := rover.New(x, y, max_x, max_y, c)
			if err != nil {
				print_error(err)
				continue
			}

			for {
				// ROVER move
				log.Println("Enter command to move the rover(L|M|R), example: LMMRM")
				scanner.Scan()
				err := r.Move(scanner.Text())
				if err != nil {
					print_error(err)
					continue
				}

				x, y, c := r.Pos()
				log.Printf("ROVER position after move: %d %d %s\n", x, y, c.Str())
				break
			}
		}
	}
}

func print_error(err error) {
	log.Println("ERR:", err)
}
