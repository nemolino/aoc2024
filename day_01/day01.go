package day01

import (
	"sort"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Part1And2(lines []string) (int, int, error) {

	l1 := make([]int, len(lines))
	l2 := make([]int, len(lines))

	for i, line := range lines {
		pair := strings.Fields(line)
		l1[i] = utils.StringToInt(&pair[0])
		l2[i] = utils.StringToInt(&pair[1])
	}
	sort.Ints(l1)
	sort.Ints(l2)

	// part1
	result1 := 0
	for i := 0; i < len(l1); i++ {
		result1 += utils.Abs(l1[i] - l2[i])
	}

	// part2
	result2 := 0
	j := 0
	for i := 0; i < len(l1); i++ {
		count := 0
		inc := 0
		for j < len(l2) {
			if l2[j] < l1[i] {
				j++
				inc++
			} else if l2[j] == l1[i] {
				count++
				j++
				inc++
			} else {
				break
			}
		}
		if i+1 < len(l1) && l1[i] == l1[i+1] {
			j -= inc
		}
		result2 += l1[i] * count
	}
	return result1, result2, nil
}
