package day01

import (
	"regexp"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Part1And2(lines []string) (int, int, error) {

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
