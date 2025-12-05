package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

var fresh = 0
var totalFresh = 0

func combine(ranges []Range) []Range {
	result := []Range{}
	currentRange := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].Min <= currentRange.Max+1 {
			currentRange.Max = max(ranges[i].Max, currentRange.Max)
			continue
		}
		result = append(result, currentRange)
		currentRange = ranges[i]
	}
	result = append(result, currentRange)
	return result
}

func calculateTotalFreshIds(ranges []Range) {
	for _, oneRange := range ranges {
		totalFresh += oneRange.Max - oneRange.Min + 1
	}
}

func solve(ranges []Range, id int) {
	for _, oneRange := range ranges {
		if id >= oneRange.Min && id <= oneRange.Max {
			fresh++
			break
		}
	}
}

func main() {
	file, _ := os.ReadFile("input.txt")
	rawInput := strings.Split(string(file), "\n\n")
	rawRanges := strings.Split(rawInput[0], "\n")
	rawIdsToCheck := strings.Split(rawInput[1], "\n")
	rawIdsToCheck = slices.DeleteFunc(rawIdsToCheck, func(s string) bool {
		return s == ""
	})
	ranges := make([]Range, len(rawRanges))
	for i, s := range rawRanges {
		line := strings.Split(s, "-")
		rangeMin, _ := strconv.Atoi(line[0])
		rangeMax, _ := strconv.Atoi(line[1])
		ranges[i] = Range{
			Min: rangeMin,
			Max: rangeMax,
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})
	ranges = combine(ranges)
	calculateTotalFreshIds(ranges)
	for _, s := range rawIdsToCheck {
		id, _ := strconv.Atoi(s)
		solve(ranges, id)
	}
	fmt.Println(fresh)
	fmt.Println(totalFresh)
}
