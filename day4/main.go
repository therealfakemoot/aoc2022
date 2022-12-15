package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func LoadRecords(fn string) ([]string, error) {
	records := make([]string, 0)
	f, err := os.Open(fn)
	if err != nil {
		return records, fmt.Errorf("error opening inventory source file: %w", err)
	}
	/*
		 var s scanner.Scanner
		 s.Init(f)
		s := scanner.NewScanner(f)
		 s.Whitespace = 1 << '\n'
		for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
			records = append(records, s.TokenText())
			// fmt.Printf("%s: %s\n", s.Position, s.TokenText())
		}
	*/
	s := bufio.NewScanner(f)
	for s.Scan() {
		records = append(records, s.Text())
	}
	return records, nil
}

func main() {
	var (
		filename string
	)

	flag.StringVar(&filename, "filename", "input.txt", "inventory source file")

	flag.Parse()
	records, err := LoadRecords(filename)
	if err != nil {
		log.Fatalf("error loading records: %s\n", err)
	}
	log.Printf("%#+v\n", records)

}
