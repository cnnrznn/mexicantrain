package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Domino struct {
	n1, n2 int
}

type Node struct {
	visited bool
	Domino
}

func main() {
	ds, start, err := readDominos()
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes := []*Node{}

	for _, d := range ds {
		nodes = append(nodes, &Node{
			Domino:  d,
			visited: false,
		})
	}

	longest := findLongest(nodes, start)

	for _, d := range longest {
		fmt.Println(d)
	}
}

func findLongest(nodes []*Node, prev int) []Domino {
	result := []Domino{}

	for _, n := range nodes {
		if n.visited {
			continue
		}
		if prev != -1 && prev != n.n1 && prev != n.n2 {
			// -1 if first domino
			// previous domino value is not equal to either face on current domino
			continue
		}

		subtrain := tryBase(n, nodes, prev)

		if len(subtrain) > len(result) {
			result = subtrain
		}
	}

	return result
}

func tryBase(n *Node, nodes []*Node, prev int) []Domino {
	n.visited = true
	defer func() {
		n.visited = false
	}()

	d := Domino{}
	if n.n1 == prev {
		d.n1 = n.n1
		d.n2 = n.n2
	} else {
		d.n1 = n.n2
		d.n2 = n.n1
	}

	result := []Domino{d}

	result = append(result, findLongest(nodes, d.n2)...)

	return result
}

func readDominos() ([]Domino, int, error) {
	result := []Domino{}
	start := -1

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Round?")
	fmt.Print("> ")

	startstr, err := reader.ReadString('\n')
	if err != nil {
		return nil, -1, err
	}

	start, err = strconv.Atoi(strings.Split(startstr, "\n")[0])
	if err != nil {
		return nil, -1, err
	}

	fmt.Println("How many Dominos?")
	fmt.Print("> ")

	nstr, err := reader.ReadString('\n')
	if err != nil {
		return nil, -1, err
	}

	n, err := strconv.Atoi(strings.Split(nstr, "\n")[0])
	if err != nil {
		return nil, -1, err
	}

	for i := 0; i < n; i++ {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, -1, err
		}

		ls := strings.Split(line, ",")

		n1, err := strconv.Atoi((strings.TrimSpace(ls[0])))
		if err != nil {
			return nil, -1, err
		}
		n2, err := strconv.Atoi(strings.TrimSpace(ls[1]))
		if err != nil {
			return nil, -1, err
		}

		result = append(result, Domino{
			n1: n1,
			n2: n2,
		})
	}

	return result, start, nil
}
