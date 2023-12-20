package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type Direction byte

type SpaceType = byte

const (
	Empty         SpaceType = '.'
	Mirror1       SpaceType = '/'
	Mirror2       SpaceType = '\\'
	SplitterXAxis SpaceType = '-'
	SplitterYAxis SpaceType = '|'
)

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Cavern [][]byte

type EnergyBoard [][]bool

type Beam struct {
	PosX int
	PosY int
	Direction
}

func (b *Beam) Mirror1() {
	switch b.Direction {
	case NORTH:
		b.Direction = EAST
	case EAST:
		b.Direction = NORTH
	case SOUTH:
		b.Direction = WEST
	case WEST:
		b.Direction = SOUTH
	}

}

func (b *Beam) reverse() {
	switch b.Direction {
	case NORTH:
		b.Direction = SOUTH
	case EAST:
		b.Direction = WEST
	case SOUTH:
		b.Direction = NORTH
	case WEST:
		b.Direction = EAST
	}

}

func (b *Beam) Mirror2() {
	b.Mirror1()
	b.reverse()
}

func (b *Beam) isDirectionYAxis() bool {
	return b.Direction == NORTH || b.Direction == SOUTH
}

func (b *Beam) isDirectionXAxis() bool {
	return b.Direction == EAST || b.Direction == WEST
}

func (b *Beam) Copy() Beam {
	return Beam{PosX: b.PosX, PosY: b.PosY, Direction: b.Direction}
}

func (b *Beam) do90Turn() {
	switch b.Direction {
	case NORTH:
		b.Direction = WEST
	case EAST:
		b.Direction = NORTH
	case SOUTH:
		b.Direction = EAST
	case WEST:
		b.Direction = SOUTH
	}
}

func createEnergyBoard(c Cavern) EnergyBoard {
	EnergyB := make(EnergyBoard, len(c))
	for i := range c {
		EnergyB[i] = make([]bool, len(c[i]))
	}
	return EnergyB
}

func (e EnergyBoard) getEnergy() int {
	energyCount := 0
	for i := range e {
		for j := range e[i] {
			if e[i][j] {
				energyCount++
			}
		}
	}
	return energyCount
}

func TraverseCavern(b Beam, c Cavern, board EnergyBoard, wg *sync.WaitGroup) {

	wg.Add(1)
	defer wg.Done()
	maxX := len(c[0])
	maxY := len(c)

	for {
		switch b.Direction {
		case NORTH:
			b.PosY--
		case EAST:
			b.PosX++
		case SOUTH:
			b.PosY++
		case WEST:
			b.PosX--
		}

		if b.PosY >= maxY || b.PosY < 0 ||
			b.PosX >= maxX || b.PosX < 0 {
			break
		}

		board[b.PosY][b.PosX] = true

		switch c[b.PosY][b.PosX] {
		case Empty:
		case Mirror1:
			b.Mirror1()
		case Mirror2:
			b.Mirror2()
		case SplitterXAxis:
			if b.isDirectionYAxis() {
				//split
				b.do90Turn()
				b2 := b.Copy()
				b2.reverse()
				go TraverseCavern(b2, c, board, wg)
			}

		case SplitterYAxis:
			if b.isDirectionXAxis() {
				//split
				b.do90Turn()
				b2 := b.Copy()
				b2.reverse()
				go TraverseCavern(b2, c, board, wg)
			}
		}

	}
}

func (c Cavern) Print() {
	fmt.Println("Printing of cavern:")
	for i := range c {
		fmt.Printf("%c\n", c[i])
	}
}

func readData() Cavern {

	// f, err := os.Open("data.txt")
	f, err := os.Open("testdata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0

	var cavern Cavern

	for scanner.Scan() {

		cavern = append(cavern, []byte(scanner.Text()))
		// fmt.Printf("value for line %d: %c\n", i, platform[i])
		i++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cavern

}

func main() {

	cavern := readData()

	cavern.Print()

	b1 := Beam{PosX: 0, PosY: 0, Direction: EAST}

	energyBoard := createEnergyBoard(cavern)

	var wg sync.WaitGroup
	TraverseCavern(b1, cavern, energyBoard, &wg)

	wg.Wait()

	fmt.Printf("Energy Count: %d", energyBoard.getEnergy())

	//profiling part

	// f, err := os.Create("myprogram.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

}
