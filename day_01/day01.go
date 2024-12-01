package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part12(lines []string) (int, int, error) {

	l1 := make([]int, len(lines))
	l2 := make([]int, len(lines))

	for i, line := range lines {
		pair := strings.Split(line, "   ")
		if x, err := strconv.Atoi(pair[0]); err != nil {
			return -1, -1, err
		} else {
			l1[i] = x
		}
		if x, err := strconv.Atoi(pair[1]); err != nil {
			return -1, -1, err
		} else {
			l2[i] = x
		}
	}
	sort.Ints(l1)
	sort.Ints(l2)

	// part1
	result1 := 0
	for i := 0; i < len(l1); i++ {
		result1 += int(math.Abs(float64(l1[i] - l2[i])))
	}

	// part2
	result2 := 0
	j := 0
	for i := 0; i < len(l1); i++ {
		count := 0
		for j < len(l2) {
			if l2[j] < l1[i] {
				j++
			} else if l2[j] == l1[i] {
				count++
				j++
			} else {
				break
			}
		}
		result2 += l1[i] * count
	}

	return result1, result2, nil
}
