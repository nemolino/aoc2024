package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(day int, inputFilename string, f func([]string) (int, int, error)) {
	var lines []string = ReadInputFromFile(inputFilename)
	if res1, res2, err := f(lines); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	} else {
		fmt.Printf("Day %02d %s\n", day, strings.Repeat("-", 73))
		fmt.Println()
		fmt.Println("part1 : ", res1)
		fmt.Println("part2 : ", res2)
		fmt.Println()
	}
}

func ReadInputFromFile(filepath string) []string {

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func StringToInt(s *string) int {
	x, err := strconv.Atoi(*s)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func MapStringToInt(s []string) []int {
	x := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		x[i] = StringToInt(&s[i])
	}
	return x
}
