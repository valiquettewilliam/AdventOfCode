package main

import (
	"fmt"
	"slices"
	"sync"
)

type Platform [][]rune

type Direction = int

type RockType = rune

const (
	Empty   RockType  = '.'
	Rounded RockType  = 'O'
	Squared RockType  = '#'
	NORTH   Direction = -1
	EAST    Direction = 1
	WEST    Direction = -1
	SOUTH   Direction = 1
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

func shiftXAxis(p Platform, d Direction) {
	// defer timer("shiftEast")()

	var wg sync.WaitGroup
	for i := range p {
		wg.Add(1)
		go func(plat Platform, i int) {
			// defer timer("shiftEast")()
			defer wg.Done()
			shift(plat[i], d)
		}(p, i)
	}
	wg.Wait()
}

func shiftEast(p Platform) {
	shiftXAxis(p, 1)
}

func shiftWest(p Platform) {
	shiftXAxis(p, -1)
}

func shiftSouth(p Platform) {
	shiftYAxis(p, SOUTH)
}
func shiftNorth(p Platform) {
	shiftYAxis(p, NORTH)
}

func shiftYAxis(p Platform, d Direction) {
	var wg sync.WaitGroup
	for j := range p[0] {
		wg.Add(1)
		go func(plat Platform, j int) {
			// defer timer("shiftEast")()
			defer wg.Done()
			shiftColumn(plat, j, d)
		}(p, j)
	}
	wg.Wait()

}

func shiftColumn(p Platform, j int, d Direction) {
	i := 0
	if d < 0 {
		i = len(p) - 1
	}

	for i >= 0 && i < len(p) {
		if p[i][j] == Rounded {
			//find next non empty space
			nextI := i + d
			for nextI >= 0 && nextI < len(p) && p[nextI][j] == Empty {
				nextI += d
			}
			p[i][j] = Empty
			p[nextI-d][j] = Rounded
		}
		i -= d
	}

}

func shift(line []RockType, d Direction) {
	i := 0
	if d > 0 {
		i = len(line) - 1
	}

	for i >= 0 && i < len(line) {
		if line[i] == Rounded {
			//find next non empty space
			nextI := i + d
			for nextI >= 0 && nextI < len(line) && line[nextI] == Empty {
				nextI += d
			}
			line[i] = Empty
			line[nextI-d] = Rounded
		}
		i -= d
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
