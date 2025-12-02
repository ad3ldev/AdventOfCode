package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IntSet map[int]struct{}

// NewIntSet creates and returns a new IntSet.
func NewIntSet() IntSet {
	return make(IntSet)
}

// Add adds an element to the set.
func (s IntSet) Add(element int) {
	s[element] = struct{}{}
}

// Contains checks if an element exists in the set.
func (s IntSet) Contains(element int) bool {
	_, exists := s[element]
	return exists
}

// Remove removes an element from the set.
type Range struct {
	first int64
	last  int64
}

var invalidCountPart1 int64 = 0
var invalidCountPart2 int64 = 0

func checkInvalidPart1(id string) {
	idLength := len(id)
	if idLength%2 != 0 {
		return
	}
	firstHalf := id[0 : idLength/2]
	secondHalf := id[idLength/2:]
	if firstHalf == secondHalf {
		idInt, _ := strconv.ParseInt(id, 10, 64)
		invalidCountPart1 += idInt
	}
}
func checkInvalidPart2(id string) {
	idLength := len(id)
	for i := 1; i <= idLength/2; i++ {
		if idLength%i != 0 {
			continue
		}
		parts := make([]string, idLength/i)
		index := 0
		for j := 0; j < idLength; j += i {
			parts[index] = id[j : j+i]
			index++
		}
		invalid := checkParts(parts)
		if invalid {
			idInt, _ := strconv.ParseInt(id, 10, 64)
			invalidCountPart2 += idInt
			return
		}
	}
}

func checkParts(parts []string) bool {
	set := make(map[string]any)
	for _, part := range parts {
		_, ok := set[part]
		if !ok {
			set[part] = true
		}
	}
	return len(set) == 1
}

func allIds(first int64, last int64) {
	for i := first; i <= last; i++ {
		id := strconv.FormatInt(i, 10)
		checkInvalidPart1(id)
		checkInvalidPart2(id)
	}
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	rawInput := string(file)
	for part := range strings.SplitSeq(rawInput, ",") {
		rawRange := strings.Split(part, "-")
		first, _ := strconv.ParseInt(strings.Trim(rawRange[0], "\n"), 10, 64)
		last, _ := strconv.ParseInt(strings.Trim(rawRange[1], "\n"), 10, 64)
		allIds(first, last)
	}
	fmt.Println(invalidCountPart1)
	fmt.Println(invalidCountPart2)
}
