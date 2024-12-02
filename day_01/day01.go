package day01

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_01/part12.input")
	if res1, res2, err := part1And2(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Println("Day 01 " + strings.Repeat("-", 73))
		fmt.Println()
		fmt.Println("part1 : ", res1)
		fmt.Println("part2 : ", res2)
		fmt.Println()
	}
}

func part1And2(lines []string) (int, int, error) {

	l1 := make([]int, len(lines))
	l2 := make([]int, len(lines))

	for i, line := range lines {
		pair := strings.Fields(line)
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
		result1 += utils.Abs(l1[i] - l2[i])
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
