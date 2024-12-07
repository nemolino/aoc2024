package day01

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nemolino/aoc2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_06/part12.input")
	if res1, res2, err := part1And2(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Println("Day 06 " + strings.Repeat("-", 73))
		fmt.Println()
		fmt.Println("part1 : ", res1)
		fmt.Println("part2 : ", res2)
		fmt.Println()
	}
}

type pos struct {
	x int
	y int
}

type direction struct {
	xIncrement int
	yIncrement int
}

func part1And2(lines []string) (int, int, error) {

	var curPos pos

	OBSTACLE_POSITIONS := sets.New[pos]()
	for i, row := range lines {
		for j, el := range row {
			if el == '#' {
				OBSTACLE_POSITIONS.Insert(pos{x: i, y: j})
			} else if el == '^' {
				curPos.x = i
				curPos.y = j
			}
		}
	}

	SIZE := len(lines)

	DIRECTIONS := [4]direction{
		{xIncrement: -1, yIncrement: 0}, // up
		{xIncrement: 0, yIncrement: 1},  // right
		{xIncrement: 1, yIncrement: 0},  // down
		{xIncrement: 0, yIncrement: -1}, // left
	}

	result1 := Part1(curPos, SIZE, &DIRECTIONS, OBSTACLE_POSITIONS)

	result2 := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			newObstaclePos := pos{x: i, y: j}
			if newObstaclePos == curPos || OBSTACLE_POSITIONS.Has(newObstaclePos) {
				continue
			}
			OBSTACLE_POSITIONS.Insert(newObstaclePos)
			if Part2(curPos, SIZE, &DIRECTIONS, OBSTACLE_POSITIONS) {
				result2++
			}
			OBSTACLE_POSITIONS.Delete(newObstaclePos)
		}
	}

	return result1, result2, nil
}

func Part1(curPos pos, SIZE int, DIRECTIONS *[4]direction, OBSTACLE_POSITIONS sets.Set[pos]) int {
	curDirectionIndex := 0
	visitedPos := sets.New[pos]()
	for PosIsValid(curPos, SIZE) {
		visitedPos.Insert(curPos)
		nextPos := GetNextPos(curPos, (*DIRECTIONS)[curDirectionIndex])
		if OBSTACLE_POSITIONS.Has(nextPos) {
			curDirectionIndex = (curDirectionIndex + 1) % 4
			nextPos = GetNextPos(curPos, (*DIRECTIONS)[curDirectionIndex])
		}
		curPos = nextPos
	}
	return visitedPos.Len()
}

func PosIsValid(p pos, size int) bool {
	return p.x >= 0 && p.x < size && p.y >= 0 && p.y < size
}

func GetNextPos(curPos pos, d direction) pos {
	return pos{
		curPos.x + d.xIncrement,
		curPos.y + d.yIncrement,
	}
}

func Part2(curPos pos, SIZE int, DIRECTIONS *[4]direction, OBSTACLE_POSITIONS sets.Set[pos]) bool {

	type state struct {
		p        pos
		dirIndex int
	}
	curDirectionIndex := 0
	visitedStates := sets.New[state]()
	for PosIsValid(curPos, SIZE) {

		curState := state{curPos, curDirectionIndex}
		if visitedStates.Has(curState) {
			return true
		}
		visitedStates.Insert(curState)

		nextPos := GetNextPos(curPos, (*DIRECTIONS)[curDirectionIndex])
		if OBSTACLE_POSITIONS.Has(nextPos) {
			curDirectionIndex = (curDirectionIndex + 1) % 4
			nextPos = GetNextPos(curPos, (*DIRECTIONS)[curDirectionIndex])
		}
		curPos = nextPos
	}
	return false
}
