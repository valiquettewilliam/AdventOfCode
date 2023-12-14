package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type RockType int64

const (
	Empty RockType = iota
	Rounded
	Squared
)

func main() {

}

func readData() {
	f, err := os.Open("testdata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		allLetters := re.FindAllString(line, -1)
		graph.AddNode()

		counter += val
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final value: %d\n", counter)

}
