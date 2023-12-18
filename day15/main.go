package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	instructions := readData()

	for _, s := range instructions {

		fmt.Println(getHash(s))

	}

	//profiling part

	// f, err := os.Create("myprogram.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

}

func getHash(s string) int {

	val := 0
	for _, c := range s {

		val += int(c)
		val = val * 17
		val = val % 256
	}

	return val

}

func readData() []string {

	// f, err := os.Open("data.txt")
	content, err := os.ReadFile("testdata.txt")
	if err != nil {
		log.Fatal(err)
	}

	StrContent := strings.ReplaceAll(string(content), "\n", "")

	return strings.Split(StrContent, ",")

}
