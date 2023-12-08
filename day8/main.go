package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Graph struct {
	nodes map[string]Node
}

type Node struct {
	Left  *Node
	Right *Node
}

func (g *Graph) AddNode(Nodeid string) {
	if _, exists := g.nodes[Nodeid]; !exists {
		newNode := Node{}
		g.nodes[Nodeid] = newNode
		fmt.Println("New node added to graph")
	} else {
		fmt.Println("Node already exists!")
	}
}

func (n *Node) AddNeighbors(left, right *Node) {
	n.Left = left
	n.Right = right
}

func CreateGraph(left, right *Node) {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0

	var graph Graph

	re := regexp.MustCompile("[A-Z]{3}")

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

func main() {

}
