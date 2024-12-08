package day05

import (
	"log"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Part1And2(lines []string) (int, int, error) {

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
