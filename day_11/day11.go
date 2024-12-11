package day11

import (
	"math"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Part1And2(lines []string) (int, int, error) {

	type args struct {
		s int
		n int
	}
	MEMO := make(map[args]int)

	var countStonesAfterNBlinks func(s int, n int) int
	countStonesAfterNBlinks = func(s int, n int) int {
		key := args{s, n}
		value, ok := MEMO[key]
		if ok {
			return value
		}

		for n > 0 {
			if s == 0 {
				s = 1
				n--
			} else if utils.CountDigits(s)%2 == 0 {
				x := int(math.Pow(10, float64(utils.CountDigits(s)/2)))
				s1 := s / x
				s2 := s % x
				n--

				var s1Res int
				key1 := args{s1, n}
				value, ok = MEMO[key1]
				if ok {
					s1Res = value
				} else {
					s1Res = countStonesAfterNBlinks(s1, n)
					MEMO[key1] = s1Res
				}

				var s2Res int
				key2 := args{s2, n}
				value, ok = MEMO[key2]
				if ok {
					s2Res = value
				} else {
					s2Res = countStonesAfterNBlinks(s2, n)
					MEMO[key2] = s2Res
				}

				MEMO[key] = s1Res + s2Res
				return s1Res + s2Res
			} else {
				s *= 2024
				n--
			}
		}
		MEMO[key] = 1
		return 1
	}

	stones := utils.MapStringToInt(strings.Fields(lines[0]))
	result1 := 0
	result2 := 0
	for _, s := range stones {
		result1 += countStonesAfterNBlinks(s, 25)
		result2 += countStonesAfterNBlinks(s, 75)
	}
	return result1, result2, nil
}
