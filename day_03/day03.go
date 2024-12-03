package day01

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Solve() {
	var lines []string = utils.ReadInputFromFile("inputs/day_03/part12.input")
	if res1, res2, err := part1And2(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Println("Day 03 " + strings.Repeat("-", 73))
		fmt.Println()
		fmt.Println("part1 : ", res1)
		fmt.Println("part2 : ", res2)
		fmt.Println()
	}
}

func part1And2(lines []string) (int, int, error) {

	result1 := 0
	result2 := 0
	enabled := true
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	for _, line := range lines {
		for _, match := range r.FindAllString(line, -1) {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else {
				s := strings.Split(match[4:len(match)-1], ",")
				x := utils.StringToInt(&s[0])
				y := utils.StringToInt(&s[1])
				result1 += x * y
				if enabled {
					result2 += x * y
				}
			}
		}
	}
	return result1, result2, nil
}
