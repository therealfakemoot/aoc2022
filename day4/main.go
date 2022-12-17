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

func (s Span) Overlaps(other Span) bool {
	if s.Stop < other.Stop && other.Start < s.Stop {
		return true
	}
	return false
}

func (s Span) Contains(other Span) bool {
	//if s.Stop <= other.Stop && other.Start < s.Stop {
	if other.Start >= s.Start && other.Stop <= s.Stop {
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
		if r[0].Contains(r[1]) || r[1].Contains(r[0]) {
			sum++
		}
	}
	fmt.Printf("Part 1 sum: %d\n", sum)
}

func Part2(records [][2]Span) {
	sum := 0
	for _, pair := range chunkBy(records, 2) {
		//if r[0].Overlaps(r[1]) || r[1].Overlaps(r[0]) {
		a, b := pair[0], pair[1]
		if (a[0].Overlaps(a[1]) || (a[1].Overlaps(a[0]))) && (b[0].Overlaps(b[1]) || (b[1].Overlaps(b[0]))) {
			sum++
		}
	}
	fmt.Printf("Part 2 sum: %d\n", sum)
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
	Part2(records)

}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
