package day15

import (
	"log"
	"slices"
)

type pos struct {
	i int
	j int
}

type direction struct {
	iIncrement int
	jIncrement int
}

var DIRECTIONS = [4]direction{
	{iIncrement: -1, jIncrement: 0}, // up
	{iIncrement: 0, jIncrement: 1},  // right
	{iIncrement: 1, jIncrement: 0},  // down
	{iIncrement: 0, jIncrement: -1}, // left
}

func Part1And2(lines []string) (int, int, error) {

	// --- part 1 -------------------------------------------------------------

	SIZE := len(lines[0])
	M := make([][]byte, SIZE)
	var robotCurPos pos
	for i := 0; i < SIZE; i++ {
		M[i] = make([]byte, SIZE)
		for j := 0; j < SIZE; j++ {
			M[i][j] = lines[i][j]
			if M[i][j] == '@' {
				robotCurPos = pos{i, j}
			}
		}
	}
	for i := SIZE + 1; i < len(lines); i++ {
		for _, c := range lines[i] {
			simulateMovePart1(c, &robotCurPos, M)
		}
	}
	result1 := calculateResult(M, SIZE, SIZE, 'O')

	// --- part 2 -------------------------------------------------------------

	SIZE2 := 2 * SIZE
	M2 := make([][]byte, SIZE)
	for i := 0; i < SIZE; i++ {
		M2[i] = make([]byte, SIZE2)
		for j := 0; j < SIZE; j++ {
			switch lines[i][j] {
			case '#':
				M2[i][2*j] = '#'
				M2[i][2*j+1] = '#'
			case 'O':
				M2[i][2*j] = '['
				M2[i][2*j+1] = ']'
			case '.':
				M2[i][2*j] = '.'
				M2[i][2*j+1] = '.'
			case '@':
				robotCurPos = pos{i, 2 * j}
				M2[i][2*j] = '@'
				M2[i][2*j+1] = '.'
			default:
				log.Fatal("unreachable")
			}
		}
	}
	for i := SIZE + 1; i < len(lines); i++ {
		for _, c := range lines[i] {
			simulateMovePart2(c, &robotCurPos, M2)
		}
	}
	result2 := calculateResult(M2, SIZE, SIZE2, '[')

	return result1, result2, nil
}

func calculateResult(M [][]byte, nRows int, nCols int, c byte) int {
	res := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if M[i][j] == c {
				res += 100*i + j
			}
		}
	}
	return res
}

func simulateMovePart1(c rune, robotCurPos *pos, M [][]byte) {

	pos := *robotCurPos
	var idx int
	switch c {
	case '^':
		idx = 0
	case 'v':
		idx = 2
	case '>':
		idx = 1
	case '<':
		idx = 3
	default:
		log.Fatal("unreachable")
	}

	for M[pos.i][pos.j] != '#' && M[pos.i][pos.j] != '.' {
		pos.i += DIRECTIONS[idx].iIncrement
		pos.j += DIRECTIONS[idx].jIncrement
	}
	if M[pos.i][pos.j] == '.' {
		if pos != *robotCurPos {
			for p := pos; p != *robotCurPos; {
				M[p.i][p.j] = M[p.i-DIRECTIONS[idx].iIncrement][p.j-DIRECTIONS[idx].jIncrement]
				p.i -= DIRECTIONS[idx].iIncrement
				p.j -= DIRECTIONS[idx].jIncrement
			}
			M[robotCurPos.i][robotCurPos.j] = '.'
			robotCurPos.i += DIRECTIONS[idx].iIncrement
			robotCurPos.j += DIRECTIONS[idx].jIncrement
		}
	}
}

func simulateMovePart2(c rune, robotCurPos *pos, M [][]byte) {

	if c == '>' || c == '<' {
		simulateMovePart1(c, robotCurPos, M)
	} else {
		var inc int
		if c == '^' {
			inc = -1
		} else {
			inc = 1
		}
		posi := *robotCurPos
		toMove := make([][]pos, 0)
		toMove = append(toMove, []pos{posi})
		for len(toMove[len(toMove)-1]) > 0 {
			toMove = append(toMove, make([]pos, 0))
			for _, cell := range toMove[len(toMove)-2] {
				l := len(toMove) - 1
				if M[cell.i+inc][cell.j] == '[' {
					if !slices.Contains(toMove[l], pos{cell.i + inc, cell.j}) {
						toMove[l] = append(toMove[l], pos{cell.i + inc, cell.j})
					}
					if !slices.Contains(toMove[l], pos{cell.i + inc, cell.j + 1}) {
						toMove[l] = append(toMove[l], pos{cell.i + inc, cell.j + 1})
					}
				} else if M[cell.i+inc][cell.j] == ']' {
					if !slices.Contains(toMove[l], pos{cell.i + inc, cell.j}) {
						toMove[l] = append(toMove[l], pos{cell.i + inc, cell.j})
					}
					if !slices.Contains(toMove[l], pos{cell.i + inc, cell.j - 1}) {
						toMove[l] = append(toMove[l], pos{cell.i + inc, cell.j - 1})
					}
				} else if M[cell.i+inc][cell.j] == '#' {
					return
				}
			}
		}
		if len(toMove) == 2 && len(toMove[0]) == 1 && len(toMove[1]) == 0 {
			simulateMovePart1(c, robotCurPos, M)
		} else {
			for idx := len(toMove) - 1; idx >= 0; idx-- {
				for _, p := range toMove[idx] {
					M[p.i+inc][p.j] = M[p.i][p.j]
					M[p.i][p.j] = '.'
				}
			}
			robotCurPos.i += inc
		}

	}
}
