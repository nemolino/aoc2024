package day01

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_07/part12.test_input")
	if res1, res2, err := part1And2(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Println("Day 07 " + strings.Repeat("-", 73))
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
		s := strings.Split(line, ": ")
		target := utils.StringToInt(&s[0])
		valuesAsStrings := strings.Fields(s[1])
		values := utils.MapStringToInt(valuesAsStrings)
		if checkPart1(target, values) {
			result1 += target
			result2 += target
		} else if checkPart2(target, values) {
			result2 += target
		}
	}
	return result1, result2, nil
}

func checkPart1(target int, values []int) bool {
	for op := 0; op < (1 << (len(values) - 1)); op++ {
		if evalPart1(values[1:], op, values[0]) == target {
			return true
		}
	}
	return false
}

func evalPart1(values []int, op int, acc int) int {
	if len(values) == 0 {
		return acc
	}
	if op%2 == 0 {
		return evalPart1(values[1:], op/2, acc+values[0])
	} else {
		return evalPart1(values[1:], op/2, acc*values[0])
	}
}

func checkPart2(target int, values []int) bool {
	n := int(math.Pow(3.0, float64(len(values)-1)))
	for op := 0; op < n; op++ {
		if evalPart2(values[1:], op, values[0]) == target {
			return true
		}
	}
	return false
}

func evalPart2(values []int, op int, acc int) int {
	if len(values) == 0 {
		return acc
	}
	if op%3 == 0 {
		return evalPart2(values[1:], op/3, acc+values[0])
	} else if op%3 == 1 {
		return evalPart2(values[1:], op/3, acc*values[0])
	} else {
		digitsCount := math.Floor(math.Log10(float64(values[0])) + 1.0)
		shift := int(math.Pow(10.0, digitsCount))
		return evalPart2(values[1:], op/3, (acc*shift)+values[0])
	}
}
