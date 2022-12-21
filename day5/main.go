package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	// "regexp"
	//"unicode"
)

func _transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func transpose(slice [][]string) [][]string {
	// reverse the input so you append items items onto each stack
	x := make([][]string, len(slice))
	//for i, l := range sort.Reverse(slice {
	for i := 0; i < len(slice); i++ {
		l := slice[i]
		for j := 0; j < len(slice); j++ {
			x[i][j] = make([]string, 0)
		}
	}

	return x
}

type Move struct {
	Count, From, To int
}

type Bay struct {
	Stacks [][]string
}

func LoadInputs(fn string) (Bay, []Move, error) {
	var b Bay
	m := make([]Move, 0)

	f, err := os.Open(fn)
	if err != nil {
		return b, m, fmt.Errorf("error opening input file: %w", err)
	}

	s := bufio.NewScanner(f)
	lines := make([]string, 0)
	for s.Scan() {
		t := s.Text()
		lines = append(lines, t)

	}

	var split int
	for i, line := range lines {
		if line == "" {
			split = i
		}
	}

	stacks := lines[:split-1]
	// moves := lines[split+1:]
	x := make([][]string, 0)
	for i, line := range stacks {
		x = append(x, make([]string, 0))
		for j := 1; j < len(line); j += 4 {
			x[i] = append(x[i], string(line[j]))
		}
	}

	log.Printf("%#v\n", transpose(x))

	// b.Stacks = parseStacks(lines[:split-1])
	// m = parseMoves(lines[split+1:])

	return b, m, nil
}

func main() {
	var (
		input string
	)
	flag.StringVar(&input, "input", "input.txt", "filename containing inputs")

	flag.Parse()
	b, m, err := LoadInputs(input)
	if err != nil {
		log.Fatalf("Error loading inputs: %s\n", err)
	}
	log.Printf("loaded bay: %#v\n", b)
	log.Printf("loaded moves: %#v\n", m)

}
