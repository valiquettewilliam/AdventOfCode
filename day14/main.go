package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Platform [][]rune

type RockType = rune

const (
	Empty   RockType = '.'
	Rounded RockType = 'O'
	Squared RockType = '#'
)

func main() {
	readData()
}

func readData() {
	f, err := os.Open("testdata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0

	var platform Platform

	for scanner.Scan() {

		platform = append(platform, []rune(scanner.Text()))
		// fmt.Printf("value for line %d: %c\n", i, platform[i])
		i++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("final value: %d\n", i)

	shiftedP := shiftNorth(platform)

	platform.Print()
	shiftedP.Print()

	fmt.Println(shiftedP.getNorthLoad())

}

func Copy(p Platform) (copiedP Platform) {

	copiedP = make(Platform, len(p))
	for i := range p {
		copiedP[i] = make([]rune, len(p[i]))
		copy(copiedP[i], p[i])
	}

	return copiedP
}

func (p Platform) Print() {
	fmt.Println("Printing of platform:")
	for i := range p {
		fmt.Printf("final value: %c\n", p[i])
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

func shiftNorth(p Platform) (shiftedPlat Platform) {

	fmt.Println("start of shiftNorth")
	shiftedPlat = Copy(p)
	shiftedPlat.Print()

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
