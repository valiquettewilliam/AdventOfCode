package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	readData()
}

func readData() {
	// f, err := os.Open("data.txt")
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

	platform.Print()
	platform.rotate()

	// shiftedP := shiftNorth(platform)
	// shiftedP := shiftEast(platform)

	platform.Print()
	// shiftedP.Print()

	// fmt.Println(shiftedP.getNorthLoad())

}

func Copy(p Platform) (copiedP Platform) {

	copiedP = make(Platform, len(p))
	for i := range p {
		copiedP[i] = make([]rune, len(p[i]))
		copy(copiedP[i], p[i])
	}

	return copiedP
}
