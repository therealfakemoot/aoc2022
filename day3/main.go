package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/scanner"
)

func splitString(s string) (ret [2]string) {
	ret[0], ret[1] = s[:len(s)/2], s[len(s)/2:]
	return
}

func LoadRecords(fn string) ([]string, error) {
	records := make([]string, 0)
	f, err := os.Open(fn)
	if err != nil {
		return records, fmt.Errorf("error opening inventory source file: %w", err)
	}
	var s scanner.Scanner
	s.Init(f)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		records = append(records, s.TokenText())
		// fmt.Printf("%s: %s\n", s.Position, s.TokenText())
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
		log.Fatalf("LoadRecords failed: %s\n", err)
	}

	for _, record := range records {
		log.Println(splitString(record))
	}
}
