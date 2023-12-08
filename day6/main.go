package main

import (
	"fmt"
	"math"
)

func main() {

	ways := WaysToBeatRecord(7, 9)

	fmt.Printf("%f", ways)

	ways = WaysToBeatRecord(38, 234)

	fmt.Printf("%f", ways)

	ways *= WaysToBeatRecord(67, 1027)

	fmt.Printf("%f", ways)

	ways *= WaysToBeatRecord(76, 1157)

	fmt.Printf("%f", ways)

	ways *= WaysToBeatRecord(73, 1236)

	fmt.Printf("%f", ways)

}

func WaysToBeatRecord(time int, distance int) float64 {

	var a float64 = -1
	b := float64(time)
	c := float64(-distance)

	sol1 := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	sol2 := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	fmt.Println("%f", sol1)
	fmt.Println("%f", sol2)

	return math.Ceil(sol2) - math.Ceil(sol1)

}
