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
		fmt.Printf("value for line %d: %c\n", i, platform[i])
		i++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	copiedP := Copy(platform)

	platform[0][0] = 'A'
	fmt.Printf("final value: %c\n", platform[0])
	fmt.Printf("final value: %c\n", copiedP[0])

	fmt.Printf("final value: %d\n", i)

}

func Copy(p Platform) (copiedP Platform) {

	copiedP = make(Platform, len(p))
	for i := range p {
		copiedP[i] = make([]rune, len(p[i]))
		copy(copiedP[i], p[i])
	}

	return copiedP
}

func shiftNorth(p Platform) (shiftedPlat Platform) {

	shiftedPlat = Copy(p)

	for j := range p[0] {
		for i := range p {
			switch p[i][j] {
			case Empty:
				//do nothing
			case Rounded:

			case Squared:
				//do nothing
			default:
				fmt.Printf("Opps something is not a rock at [%d][%d]: %s\n", i, j, p[i][j])

			}

		}

	}
	return
}
