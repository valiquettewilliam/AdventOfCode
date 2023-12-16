package main

import (
	"fmt"
	"slices"
	"sync"
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
	// defer timer("transpose")()
	xl := len(p[0])
	yl := len(p)
	result := make(Platform, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}

	var wg sync.WaitGroup
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			wg.Add(1)
			go transposeOneCase(p, result, i, j, &wg)
		}
	}
	wg.Wait()
	return result
}

func transposeOneCase(src, dest Platform, i, j int, wg *sync.WaitGroup) {
	dest[i][j] = src[j][i]
	wg.Done()

}

func shiftSide(p Platform, direction int) {
	// defer timer("shiftEast")()

	var wg sync.WaitGroup
	for i := range p {
		wg.Add(1)
		go func(plat Platform, i int) {
			// defer timer("shiftEast")()
			defer wg.Done()
			shift(plat[i], direction)
		}(p, i)
	}
	wg.Wait()
}

func shiftEast(p Platform) {
	shiftSide(p, 1)
}

func shiftWest(p Platform) {
	shiftSide(p, -1)
}

func shiftNorth(p Platform) (shiftedPlat Platform) {
	shiftedPlat = Copy(p)

	for j := range shiftedPlat[0] {
		for i := range shiftedPlat {
			if i == 0 {
				continue
			}

			if shiftedPlat[i][j] == Rounded {
				//find next up non empty space
				upI := i - 1
				for upI >= 0 && shiftedPlat[upI][j] == Empty {
					upI--
				}
				shiftedPlat[i][j] = Empty
				shiftedPlat[upI+1][j] = Rounded
			}
		}
	}
	return
}

func shiftSouth(p Platform) {

	for j := range p[0] {
		for i := range p {
			if i == 0 {
				continue
			}

			if p[i][j] == Rounded {
				//find next up non empty space
				downI := i + 1
				for downI < len(p) && p[downI][j] == Empty {
					downI++
				}
				p[i][j] = Empty
				p[downI-1][j] = Rounded
			}
		}
	}
	return
}

func shiftRight(line []RockType) {
	shift(line, 1)
}

func shiftLeft(line []RockType) {
	shift(line, -1)
}

func shift(line []RockType, step int) {
	// defer timer("shift")()

	// shiftedLine = make([]RockType, len(line))
	// copy(shiftedLine, line)

	i := 0
	if step > 0 {
		i = len(line) - 1
	}

	for i >= 0 && i < len(line) {
		if line[i] == Rounded {
			//find next non empty space
			nextI := i + step
			for nextI >= 0 && nextI < len(line) && line[nextI] == Empty {
				nextI += step
			}
			line[i] = Empty
			line[nextI-step] = Rounded
		}
		i -= step
	}

}

func (p *Platform) rotate() {
	defer timer("rotate")()
	TPlatform := transpose(*p)
	var wg sync.WaitGroup
	for i := range TPlatform {
		wg.Add(1)
		go func(row []rune) {
			defer wg.Done()
			defer timer("Reverse")()
			slices.Reverse(row)
		}(TPlatform[i])
	}
	wg.Wait()

	*p = TPlatform
}
