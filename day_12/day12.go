package day12

type pos struct {
	i int
	j int
}

func Part1And2(lines []string) (int, int, error) {

	SIZE := len(lines)
	M := make([][]byte, SIZE)
	for i := 0; i < SIZE; i++ {
		M[i] = make([]byte, SIZE)
	}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			M[i][j] = lines[i][j]
		}
	}

	VISITED := make([][]bool, SIZE)
	for i := 0; i < SIZE; i++ {
		VISITED[i] = make([]bool, SIZE)
	}

	result1 := 0
	result2 := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if !VISITED[i][j] {
				area, perimeter, sidesCount := exploreFrom(i, j, M, SIZE, VISITED)
				result1 += area * perimeter
				result2 += area * sidesCount
			}
		}
	}

	return result1, result2, nil
}

func exploreFrom(i int, j int, M [][]byte, SIZE int, VISITED [][]bool) (int, int, int) {

	perimeter := 0
	area := 0

	toVisit := make([]pos, 1)
	toVisit[0] = pos{i, j}

	for len(toVisit) > 0 {

		p := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		if VISITED[p.i][p.j] {
			continue
		}
		VISITED[p.i][p.j] = true

		area++

		if p.i-1 < 0 || (p.i-1 >= 0 && M[p.i-1][p.j] != M[p.i][p.j]) {
			perimeter++
		} else {
			toVisit = append(toVisit, pos{p.i - 1, p.j})
		}
		if p.i+1 >= SIZE || (p.i+1 < SIZE && M[p.i+1][p.j] != M[p.i][p.j]) {
			perimeter++
		} else {
			toVisit = append(toVisit, pos{p.i + 1, p.j})
		}
		if p.j-1 < 0 || (p.j-1 >= 0 && M[p.i][p.j-1] != M[p.i][p.j]) {
			perimeter++
		} else {
			toVisit = append(toVisit, pos{p.i, p.j - 1})
		}
		if p.j+1 >= SIZE || (p.j+1 < SIZE && M[p.i][p.j+1] != M[p.i][p.j]) {
			perimeter++
		} else {
			toVisit = append(toVisit, pos{p.i, p.j + 1})
		}
	}
	return area, perimeter, 0
}
