package day01

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_02/part12.input")
	if res1, res2, err := part1And2(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Println("Day 02 " + strings.Repeat("-", 73))
		fmt.Println()
		fmt.Println("part1 : ", res1)
		fmt.Println("part2 : ", res2)
		fmt.Println()
	}
}

func part1And2(lines []string) (int, int, error) {

	result1 := 0
	result2 := 0
	for _, line := range lines {
		report := utils.MapStringToInt(strings.Fields(line))
		if isSafe1(report) {
			result1++
			result2++
		} else if isSafe2(report) {
			result2++
		}
	}
	return result1, result2, nil
}

func isSafe1(report []int) bool {
	firstΔ := report[1] - report[0]
	if firstΔ >= 1 && firstΔ <= 3 {
		for i := 1; i < len(report)-1; i++ {
			Δ := report[i+1] - report[i]
			if Δ < 1 || Δ > 3 {
				return false
			}
		}
		return true
	} else if firstΔ >= -3 && firstΔ <= -1 {
		for i := 1; i < len(report)-1; i++ {
			Δ := report[i+1] - report[i]
			if Δ < -3 || Δ > -1 {
				return false
			}
		}
		return true
	}
	return false
}

func isSafe2(report []int) bool {
	r := make([]int, len(report))
	for i := 0; i < len(report); i++ {
		copy(r, report)
		if isSafe1(slices.Delete(r, i, i+1)) {
			return true
		}
	}
	return false
}
