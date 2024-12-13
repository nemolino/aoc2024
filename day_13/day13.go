package day13

import (
	"fmt"
)

func Part1And2(lines []string) (int, int, error) {

	var aIncX, aIncY, bIncX, bIncY, prizeX, prizeY int
	result1 := 0
	result2 := 0
	for i := 0; i < len(lines); {
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &aIncX, &aIncY)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bIncX, &bIncY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)
		i = i + 4
		result1 += tokensToWinPrize(aIncX, aIncY, bIncX, bIncY, prizeX, prizeY)
		result2 += tokensToWinPrize(aIncX, aIncY, bIncX, bIncY, prizeX+10000000000000, prizeY+10000000000000)
	}

	return result1, result2, nil
}

func tokensToWinPrize(aIncX, aIncY, bIncX, bIncY, prizeX, prizeY int) int {

	/*
		Sistema di equazioni con incognite intere aPresses, bPresses

		prizeX = aIncX * aPresses + bIncX * bPresses
		prizeY = aIncY * aPresses + bIncY * bPresses

		A = Bx + Cy
		D = Ex + Fy

		ha soluzione

				x = (AF - CD) / (-CE + BF)
				y = (BD - AE) / (-CE + BF)

		con
				B != 0
				(-CE + BF) != 0
	*/

	den := aIncX*bIncY - bIncX*aIncY
	if den == 0 {
		return 0
	}

	num1 := prizeX*bIncY - bIncX*prizeY
	if num1%den != 0 {
		// aPresses non è intero
		return 0
	}
	aPresses := num1 / den

	num2 := aIncX*prizeY - prizeX*aIncY
	if num2%den != 0 {
		// bPresses non è intero
		return 0
	}
	bPresses := num2 / den

	return 3*aPresses + bPresses
}
