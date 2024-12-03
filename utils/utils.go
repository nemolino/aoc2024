package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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
