package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	if idLength%2 != 0 {
		return
	}
	halfs := make([]string, 2)
	index := 0
	for i := 0; i <= idLength/2; i += idLength / 2 {
		j := i + idLength/2
		halfs[index] = id[i:j]
		index++
	}
	for i := range 1 {
		for j := 1; j < 2; j++ {
			if halfs[i] == halfs[j] {
				idInt, _ := strconv.ParseInt(id, 10, 64)
				invalidCountPart2 += idInt
			}
		}
	}
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
