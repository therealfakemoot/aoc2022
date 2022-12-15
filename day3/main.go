package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/scanner"
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

func splitRecord(s []string) (ret [2][]string) {
	ret[0], ret[1] = s[:len(s)/2], s[len(s)/2:]
	return
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

func commonCharacters(pairs [2][]string) []string {
	ret := make(map[string]bool)

	for _, cLeft := range pairs[0] {
		for _, cRight := range pairs[1] {
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
		split := splitRecord(record)
		commonChars := commonCharacters(split)
		log.Printf("%#+v\n", commonChars)
		p := priority(commonChars[0])
		sum += p
	}
	log.Printf("Priority sum: %d\n", sum)
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

}
