package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Span struct {
	Start, Stop int
}

func (s Span) Overlap(other Span) bool {
	if s.Stop <= other.Stop && other.Start < s.Stop {
		return true
	}
	return false
}

func NewSpan(raw string) Span {
	var s Span
	split := strings.Split(raw, "-")
	start, _ := strconv.ParseInt(split[0], 10, 64)
	s.Start = int(start)
	stop, _ := strconv.ParseInt(split[1], 10, 64)
	s.Stop = int(stop)

	return s
}

func LoadRecords(fn string) ([][2]Span, error) {
	records := make([][2]Span, 0)
	f, err := os.Open(fn)
	if err != nil {
		return records, fmt.Errorf("error opening inventory source file: %w", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var pair [2]Span
		split := strings.Split(s.Text(), ",")
		pair[0] = NewSpan(split[0])
		pair[1] = NewSpan(split[1])
		records = append(records, pair)
	}
	return records, nil
}

func Part1(records [][2]Span) {
	sum := 0
	for _, r := range records {
		if r[0].Overlap(r[1]) || r[1].Overlap(r[0]) {
			log.Printf("overlap found: %#v\n", r)
			sum++
		}
	}
	fmt.Printf("Part 1 sum: %d\n", sum)
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
	Part1(records)

}
