package day10

import (
	"slices"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

type node struct {
	value int
	i     int
	j     int
}

func Part1And2(lines []string) (int, int, error) {

	SIZE := len(lines)
	M := make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		M[i] = utils.MapStringToInt(strings.Split(lines[i], ""))
	}

	result1 := 0
	result2 := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if M[i][j] == 0 {
				score, rating := exploreFromSource(node{0, i, j}, M, SIZE)
				result1 += score
				result2 += rating
			}
		}
	}
	return result1, result2, nil
}

func exploreFromSource(source node, M [][]int, SIZE int) (int, int) {

	reachableDestinations := make([]*node, 0)

	q := make([]*node, 0)
	q = append(q, &source)

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if x.value == 9 && !slices.Contains(reachableDestinations, x) {
			reachableDestinations = append(reachableDestinations, x)
			continue
		}
		if x.j+1 < SIZE && M[x.i][x.j+1] == x.value+1 {
			q = append(q, &node{x.value + 1, x.i, x.j + 1})
		}
		if x.j-1 >= 0 && M[x.i][x.j-1] == x.value+1 {
			q = append(q, &node{x.value + 1, x.i, x.j - 1})
		}
		if x.i-1 >= 0 && M[x.i-1][x.j] == x.value+1 {
			q = append(q, &node{x.value + 1, x.i - 1, x.j})
		}
		if x.i+1 < SIZE && M[x.i+1][x.j] == x.value+1 {
			q = append(q, &node{x.value + 1, x.i + 1, x.j})
		}
	}

	rating := len(reachableDestinations)

	distinctReachableDestinations := make([]node, 0)
	for _, x := range reachableDestinations {
		if !slices.Contains(distinctReachableDestinations, *x) {
			distinctReachableDestinations = append(distinctReachableDestinations, *x)
		}
	}
	score := len(distinctReachableDestinations)

	return score, rating
}
