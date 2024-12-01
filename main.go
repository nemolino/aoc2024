package main

import (
	"fmt"
	"log"
	"os"

	day01 "github.com/nemolino/aoc2024/day_01"
	"github.com/nemolino/aoc2024/utils"
)

func main() {

	var lines []string = utils.ReadInputFromFile("inputs/day_01/input.txt")

	res1, res2, err := day01.Part12(lines)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	fmt.Println("part1 : ", res1)
	fmt.Println("part2 : ", res2)

}
