package day01

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_05/part12.input")
	if res1, res2, err := part1And2(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Println("Day 05 " + strings.Repeat("-", 73))
		fmt.Println()
		fmt.Println("part1 : ", res1)
		fmt.Println("part2 : ", res2)
		fmt.Println()
	}
}

func part1And2(lines []string) (int, int, error) {

	var cannotAppearBefore [100][100]bool
	// cannotAppearBefore[i][j] = True if i cannot apper before j in the updates

	result1 := 0
	result2 := 0
	k := 0
	for ; lines[k] != ""; k++ {
		s := strings.Split(lines[k], "|")
		j := utils.StringToInt(&s[0])
		i := utils.StringToInt(&s[1])
		cannotAppearBefore[i][j] = true
	}
	for k = k + 1; k < len(lines); k++ {
		update := utils.MapStringToInt(strings.Split(lines[k], ","))
		if r := checkUpdate(update, &cannotAppearBefore); r > 0 {
			result1 += r
		} else {
			result2 += correctUpdate(update, &cannotAppearBefore)
		}
	}
	return result1, result2, nil
}

func checkUpdate(u []int, cannotAppearBefore *[100][100]bool) int {
	for m := 1; m < len(u); m++ {
		for n := 0; n < m; n++ {
			if cannotAppearBefore[u[n]][u[m]] {
				return 0
			}
		}
	}
	return u[len(u)/2]
}

func correctUpdate(u []int, cannotAppearBefore *[100][100]bool) int {
	targetIndex := len(u) / 2
	for _, m := range u {
		mNewIndex := 0
		for _, n := range u {
			if cannotAppearBefore[m][n] {
				mNewIndex++
			}
		}
		if mNewIndex == targetIndex {
			return m
		}
	}
	log.Fatal("unreachable")
	return -1
}
