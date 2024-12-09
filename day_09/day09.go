package day09

import (
	"slices"
	"strings"

	"github.com/nemolino/aoc2024/utils"
)

func Part1And2(lines []string) (int, int, error) {

	input := utils.MapStringToInt(strings.Split(lines[0], ""))
	disk := make([]int, 0)
	for i := 0; i < len(input); {
		for n := 0; n < input[i]; n++ {
			disk = append(disk, i/2)
		}
		i++
		if i >= len(input) {
			break
		}
		for n := 0; n < input[i]; n++ {
			disk = append(disk, -1)
		}
		i++
	}

	diskCopy := slices.Clone(disk)
	part1(diskCopy)
	result1 := checksum(diskCopy)

	part2(disk)
	result2 := checksum(disk)

	return result1, result2, nil
}

func part1(disk []int) {
	for start, end := 0, len(disk)-1; ; {
		for disk[start] != -1 && start < end {
			start++
		}
		if start >= end {
			break
		}
		for disk[end] == -1 && start < end {
			end--
		}
		if start >= end {
			break
		}
		disk[start], disk[end] = disk[end], disk[start]
	}
}

func checksum(disk []int) int {
	c := 0
	for i, v := range disk {
		if v != -1 {
			c += i * v
		}
	}
	return c
}

func part2(disk []int) {

	type spaceChunk struct {
		idx  int
		size int
	}

	sc := make([]spaceChunk, 0)
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			j := i
			for ; disk[j] == -1 && j < len(disk); j++ {
			}
			sc = append(sc, spaceChunk{idx: i, size: j - i})
			i = j
		}
	}

	expectedFileIDToMove := disk[len(disk)-1]
	for i := len(disk) - 1; i >= 0; i-- {

		if disk[i] != -1 {
			if disk[i] != expectedFileIDToMove {
				continue
			} else {
				expectedFileIDToMove--
			}
			id := disk[i]
			j := i
			for ; j >= 0; j-- {
				if disk[j] != id {
					break
				}
			}
			// file id is in disk[j+1] ... disk[i]
			fileSize := i - j
			for k := 0; k < len(sc); k++ {
				if sc[k].size >= fileSize && sc[k].idx < j+1 {
					for h := j + 1; h <= i; h++ {
						disk[h] = -1
					}
					for h := sc[k].idx; h < sc[k].idx+fileSize; h++ {
						disk[h] = id
					}
					sc[k].idx += fileSize
					sc[k].size -= fileSize
					break
				}
			}
			i = j + 1
		}
	}
}
