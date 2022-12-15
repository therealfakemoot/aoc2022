package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/scanner"
	// "github.com/davecgh/go-spew/spew"
)

func priority(s string) int {
	for i, p := range priorities {
		if p == s {
			return i + 1
		}
	}

	return 0
}

var priorities = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func splitRecord(s []string) ([]string, []string) {
	left, right := s[:len(s)/2], s[len(s)/2:]
	return left, right
}

func explodeString(s string) []string {
	ret := make([]string, 0)
	for _, r := range s {
		ret = append(ret, string(r))
	}
	return ret
}

func LoadRecords(fn string) ([][]string, error) {
	records := make([][]string, 0)
	f, err := os.Open(fn)
	if err != nil {
		return records, fmt.Errorf("error opening inventory source file: %w", err)
	}
	var s scanner.Scanner
	s.Init(f)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		records = append(records, explodeString(s.TokenText()))
		// fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
	return records, nil
}

func commonCharacters(left, right []string) []string {
	ret := make(map[string]bool)

	for _, cLeft := range left {
		for _, cRight := range right {
			if cLeft == cRight {
				ret[cLeft] = true
			}
		}
	}

	dedupe := make([]string, 0)
	for k := range ret {
		dedupe = append(dedupe, k)
	}
	return dedupe
}

func Part1(records [][]string) {
	sum := 0
	for _, record := range records {
		left, right := splitRecord(record)
		commonChars := commonCharacters(left, right)
		log.Printf("%#+v\n", commonChars)
		p := priority(commonChars[0])
		sum += p
	}
	log.Printf("Priority sum: %d\n", sum)
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func Part2(records [][]string) {
	// chunk the records by 3
	chunked := chunkBy(records, 3)
	// find common item for each chunk
	sum := 0
	for _, triplet := range chunked {

		ab := commonCharacters(triplet[0], triplet[1])
		bc := commonCharacters(triplet[1], triplet[2])
		abc := commonCharacters(ab, bc)

		log.Printf("Common characters for triplet: %+v\n", abc)
		sum += priority(abc[0])
	}
	log.Printf("part 2 sum: %d\n", sum)
	// sum priorities
}

func main() {
	var (
		filename string
	)

	flag.StringVar(&filename, "filename", "input.txt", "inventory source file")

	flag.Parse()
	records, err := LoadRecords(filename)
	if err != nil {
		log.Fatalf("LoadRecords failed: %s\n", err)
	}

	Part1(records)
	Part2(records)

}
