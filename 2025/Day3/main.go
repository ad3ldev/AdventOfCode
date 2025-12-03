package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var totalMaxJolt = 0

func getLargest(bank string, index, length int) (int, int) {
	maxCurrent, _ := strconv.Atoi(string(bank[index]))
	maxIdx := index
	for i := index + 1; i < length; i++ {
		current, _ := strconv.Atoi(string(bank[i]))
		if current > maxCurrent {
			maxCurrent = current
			maxIdx = i
		}
	}
	return maxCurrent, maxIdx
}

func solve(bank string, batteriesRequired int) {
	maxJolt := 0
	maxCurrent := 0
	maxIdx := 0
	for i := range batteriesRequired {
		maxCurrent, maxIdx = getLargest(bank, maxIdx, len(bank)-batteriesRequired+1+i)
		maxJolt = maxJolt*10 + maxCurrent
		maxIdx++
	}
	fmt.Println(maxJolt)
	totalMaxJolt += maxJolt
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	rawInput := string(file)
	for part := range strings.SplitSeq(rawInput, "\n") {
		if part == "" {
			continue
		}
		solve(part, 12)
	}
	fmt.Println(totalMaxJolt)
}
