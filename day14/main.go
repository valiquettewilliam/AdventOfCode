package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type RockType rune

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

	var platform []string

	for scanner.Scan() {

		platform = append(platform, scanner.Text())
		fmt.Printf("final value: %s\n", platform[i])
		i++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final value: %d\n", i)

}
