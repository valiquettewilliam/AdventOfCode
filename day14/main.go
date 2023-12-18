package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var previousPlatform Platform

func main() {
	platform := readData()
	//profiling part

	f, err := os.Create("myprogram.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	platform.Print()
	// platform.rotate()

	// shiftColumn(platform, 0, NORTH)

	// shiftNorth(platform)
	// platform.Print()
	// shiftWest(platform)
	// platform.Print()
	// shiftSouth(platform)
	// platform.Print()
	// shiftEast(platform)
	// platform.Print()
	nbOfCycle := 1000000000
	// nbOfCycle := 3

	start := time.Now()
	previousPlatform = make(Platform, len(platform))
	for i := range previousPlatform {
		previousPlatform[i] = make([]rune, len(platform[i]))
	}

	isStabalized := false
	for i := 0; i < nbOfCycle && !isStabalized; i++ {

		logStep := 1000000

		if (i+logStep)%logStep == 0 {
			fmt.Printf("Progress %f %%\n", (float64(i)/float64(nbOfCycle))*100)
		}

		isStabalized = Cycle(platform)
		// platform.Print()
		if isStabalized {
			fmt.Printf("number of cycle done: %d\n", i)
			break
		}
	}
	fmt.Printf("was stabalized: %t\n", isStabalized)
	fmt.Printf("previous platoform: \n")
	previousPlatform.Print()
	platform.Print()
	fmt.Printf("took %v\n", time.Since(start))
	// // shiftedP.Print()

	fmt.Println(platform.getNorthLoad())
}

func readData() Platform {

	f, err := os.Open("data.txt")
	// f, err := os.Open("testdata.txt")
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

	return platform

}

func Cycle(p Platform) bool {

	shiftNorth(p)
	shiftWest(p)
	shiftSouth(p)
	shiftEast(p)
	if IsEqual(p, previousPlatform) {
		return true
	} else {
		Copy(p, previousPlatform)
	}
	return false

}

// todo move
func Copy(src, dest Platform) {
	// defer timer("Copy")()
	for i := range src {
		dest[i] = make([]rune, len(src[i]))
		copy(dest[i], src[i])
	}
}

func IsEqual(p1, p2 Platform) bool {
	for i := range p1 {
		for j := range p1[i] {
			if p1[i][j] != p2[i][j] {
				return false
			}
		}
	}
	return true
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

/*
0.01s 0.0072% 45.34%     11.13s  8.05%  runtime.lock (partial-inline)
     0.11s  0.08% 45.42%     10.22s  7.40%  main.shiftEast
     9.65s  6.98% 52.40%      9.71s  7.03%  main.shift (inline)
     0.06s 0.043% 52.45%      9.35s  6.77%  main.shiftWest
     0.19s  0.14% 52.58%      9.18s  6.64%  runtime.newproc.func1
*/
