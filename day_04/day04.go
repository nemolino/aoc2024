package day01

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_04/part12.input")
	res1, err := part1(lines)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	res2, err := part2(lines)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	fmt.Println("Day 04 " + strings.Repeat("-", 73))
	fmt.Println()
	fmt.Println("part1 : ", res1)
	fmt.Println("part2 : ", res2)
	fmt.Println()
}

func part1(lines []string) (int, error) {

	result := 0
	for i, line := range lines {
		for j := 0; j < len(line)-3; j++ {
			horizontalChunk := line[j : j+4]
			checkChunk(horizontalChunk, &result)
		}
		if i < len(line)-3 {
			for j := 0; j < len(line); j++ {
				verticalChunk := string(lines[i][j]) + string(lines[i+1][j]) + string(lines[i+2][j]) + string(lines[i+3][j])
				checkChunk(verticalChunk, &result)
			}
			for j := 0; j < len(line)-3; j++ {
				diagonalChunk1 := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2]) + string(lines[i+3][j+3])
				checkChunk(diagonalChunk1, &result)
			}
			for j := 3; j < len(line); j++ {
				diagonalChunk2 := string(lines[i][j]) + string(lines[i+1][j-1]) + string(lines[i+2][j-2]) + string(lines[i+3][j-3])
				checkChunk(diagonalChunk2, &result)
			}
		}
	}
	return result, nil
}

func checkChunk(chunk string, result *int) {
	if chunk == "XMAS" || chunk == "SAMX" {
		(*result)++
	}
}

func part2(lines []string) (int, error) {

	result := 0
	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < len(lines[i])-2; j++ {
			diagonalChunk1 := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2])
			if diagonalChunk1 == "MAS" || diagonalChunk1 == "SAM" {
				diagonalChunk2 := string(lines[i][j+2]) + string(lines[i+1][j+1]) + string(lines[i+2][j])
				if diagonalChunk2 == "MAS" || diagonalChunk2 == "SAM" {
					result++
				}
			}
		}
	}
	return result, nil
}
