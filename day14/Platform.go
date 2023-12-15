package main

import (
	"fmt"
	"slices"
)

type Platform [][]rune

type RockType = rune

const (
	Empty   RockType = '.'
	Rounded RockType = 'O'
	Squared RockType = '#'
)

func (p Platform) Print() {
	fmt.Println("Printing of platform:")
	for i := range p {
		fmt.Printf("%c\n", p[i])
	}
}

func (p Platform) getNorthLoad() int {
	lenght := len(p)
	load := 0
	for i, line := range p {
		for j := range line {
			if p[i][j] == Rounded {
				load += lenght
			}
		}
		lenght--
	}
	return load
}

func transpose(p Platform) Platform {
	xl := len(p[0])
	yl := len(p)
	result := make(Platform, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = p[j][i]
		}
	}
	return result
}

func shiftEast(p Platform) (shiftedPlat Platform) {
	shiftedPlat = Copy(p)

	for i := range shiftedPlat {
		shiftedPlat[i] = shiftRight(shiftedPlat[i])
	}

	return
}

func shiftNorth(p Platform) (shiftedPlat Platform) {
	shiftedPlat = Copy(p)

	for j := range shiftedPlat[0] {
		for i := range shiftedPlat {
			if i == 0 {
				continue
			}
			switch shiftedPlat[i][j] {
			case Empty:
				//do nothing
			case Rounded:
				//find next up non empty space
				upI := i - 1
				for upI >= 0 && shiftedPlat[upI][j] == Empty {
					upI--
				}
				shiftedPlat[i][j] = Empty
				shiftedPlat[upI+1][j] = Rounded

			case Squared:
				//do nothing
			default:
				fmt.Printf("Opps something is not a rock at [%d][%d]: %c\n", i, j, shiftedPlat[i][j])

			}

		}

	}
	return
}

func shiftRight(line []RockType) []RockType {
	return shift(line, 1)
}

func shiftLeft(line []RockType) []RockType {
	return shift(line, -1)
}

func shift(line []RockType, step int) (shiftedLine []RockType) {
	shiftedLine = make([]RockType, len(line))
	copy(shiftedLine, line)

	i := 0
	if step > 0 {
		i = len(line) - 1
	}

	for i >= 0 && i < len(shiftedLine) {
		if shiftedLine[i] == Rounded {
			//find next non empty space
			nextI := i + step
			for nextI >= 0 && nextI < len(shiftedLine) && shiftedLine[nextI] == Empty {
				nextI += step
			}
			shiftedLine[i] = Empty
			shiftedLine[nextI-step] = Rounded
		}
		i -= step
	}

	return
}

func (p *Platform) rotate() {

	TPlatform := transpose(*p)

	for i := range TPlatform {
		slices.Reverse(TPlatform[i])
	}

	*p = TPlatform
}
