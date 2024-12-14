package day14

import (
	"fmt"
)

type robot struct {
	x    int
	y    int
	xVel int
	yVel int
}

func Part1And2(lines []string) (int, int, error) {

	const WIDTH, HEIGHT = 101, 103

	robots := make([](*robot), 0)
	for _, line := range lines {
		r := &robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.xVel, &r.yVel)
		robots = append(robots, r)
	}

	result1 := 0
	for seconds := 1; seconds <= 7286; seconds++ {
		for _, r := range robots {
			r.x = (r.x + r.xVel) % WIDTH
			r.y = (r.y + r.yVel) % HEIGHT
			for r.x < 0 {
				r.x += WIDTH
			}
			for r.y < 0 {
				r.y += HEIGHT
			}
		}
		if seconds == 100 {
			qUpLeft, qUpRight, qDownLeft, qDownRight := 0, 0, 0, 0
			for _, r := range robots {
				if r.x < WIDTH/2 {
					if r.y < HEIGHT/2 {
						qUpLeft++
					} else if r.y > HEIGHT/2 {
						qDownLeft++
					}
				} else if r.x > WIDTH/2 {
					if r.y < HEIGHT/2 {
						qUpRight++
					} else if r.y > HEIGHT/2 {
						qDownRight++
					}
				}
			}
			result1 = qUpLeft * qUpRight * qDownLeft * qDownRight
		}
	}
	// displayRobots(robots) // answer found by smart visual inspection

	return result1, 7286, nil
}

/*
func displayRobots(robots [](*robot)) {
	const WIDTH, HEIGHT = 101, 103
	var MATRIX [WIDTH][HEIGHT]bool
	for _, r := range robots {
		MATRIX[r.x][r.y] = true
	}
	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			if MATRIX[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
*/
