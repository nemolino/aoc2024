package day08

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

type pos struct {
	x int
	y int
}

func Part1And2(lines []string) (int, int, error) {

	M := make(map[string][]pos)

	SIZE := len(lines)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if lines[i][j] != '.' {
				key := string(lines[i][j])
				_, ok := M[key]
				if !ok {
					M[key] = make([]pos, 0)
				}
				M[key] = append(M[key], pos{x: i, y: j})
			}
		}
	}

	antinodePositions := sets.New[pos]()
	antinodePositions2 := sets.New[pos]()
	var p pos
	for _, value := range M {
		for i := 0; i < len(value); i++ {
			for j := 0; j < len(value); j++ {
				if i == j {
					continue
				}
				Δ_x := value[j].x - value[i].x
				Δ_y := value[j].y - value[i].y

				// part1
				p = pos{x: value[j].x + Δ_x, y: value[j].y + Δ_y}
				if PosIsValid(p, SIZE) {
					antinodePositions.Insert(p)
				}
				p = pos{x: value[i].x - Δ_x, y: value[i].y - Δ_y}
				if PosIsValid(p, SIZE) {
					antinodePositions.Insert(p)
				}

				// part2
				p = pos{x: value[j].x, y: value[j].y}
				for PosIsValid(p, SIZE) {
					antinodePositions2.Insert(p)
					p.x += Δ_x
					p.y += Δ_y
				}
				p = pos{x: value[i].x, y: value[i].y}
				for PosIsValid(p, SIZE) {
					antinodePositions2.Insert(p)
					p.x -= Δ_x
					p.y -= Δ_y
				}
			}
		}
	}
	result1 := len(antinodePositions)
	result2 := len(antinodePositions2)

	return result1, result2, nil
}

func PosIsValid(p pos, size int) bool {
	return p.x >= 0 && p.x < size && p.y >= 0 && p.y < size
}
