package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Direction int

const (
	UP_LEFT Direction = iota
	UP
	UP_RIGHT
	LEFT
	RIGHT
	DOWN_LEFT
	DOWN
	DOWN_RIGHT
)

type DirectionCoordinates struct {
	x int
	y int
}

var directions = map[Direction]DirectionCoordinates{
	UP_LEFT: {
		x: -1,
		y: -1,
	},
	UP: {
		x: 0,
		y: -1,
	},
	UP_RIGHT: {
		x: 1,
		y: -1,
	},
	LEFT: {
		x: -1,
		y: 0,
	},
	RIGHT: {
		x: 1,
		y: 0,
	},
	DOWN_LEFT: {
		x: -1,
		y: 1,
	},
	DOWN: {
		x: 0,
		y: 1,
	},
	DOWN_RIGHT: {
		x: 1,
		y: 1,
	},
}

var grid []string
var count = 0
var totalCount = 0

func check8Around(x, y int, tempGrid []string) []string {
	check := 0
	for _, direction := range directions {
		nextX := x + direction.x
		nextY := y + direction.y
		if nextX < 0 || nextY < 0 {
			continue
		}
		if nextX >= len(tempGrid) || nextY >= len(tempGrid[x]) {
			continue
		}
		if tempGrid[nextX][nextY] == '@' || tempGrid[nextX][nextY] == 'x' {
			check++
		}
	}
	if check < 4 {
		row := []rune(tempGrid[x])
		row[y] = 'x'
		tempGrid[x] = string(row)
		count++
	}
	return tempGrid
}

func solve(recursive bool) {
	var newGrid = make([]string, len(grid))
	copy(newGrid, grid)
	for i, row := range grid {
		for j := range row {
			if row[j] == '@' {
				newGrid = check8Around(i, j, newGrid)
			}
		}
	}
	totalCount += count
	if recursive {
		for i := range newGrid {
			for j := range newGrid[i] {

				row := []rune(newGrid[i])
				if row[j] == 'x' {
					row[j] = '.'
					newGrid[i] = string(row)
				}
			}
		}
	}
	if recursive && count > 0 {
		copy(grid, newGrid)
		count = 0
		solve(recursive)
	}
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	rawInput := string(file)
	grid = strings.Split(rawInput, "\n")
	grid = slices.DeleteFunc(grid, func(s string) bool {
		return s == ""
	})
	solve(true)
	fmt.Println(totalCount)

}
