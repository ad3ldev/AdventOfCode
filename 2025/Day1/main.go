package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var dial = 50
var count = 0

func leftRotate(value int) {
	dial -= value
	for dial < 0 {
		dial += 100
	}
}
func rightRotate(value int) {
	dial += value
	for dial > 99 {
		dial -= 100
	}
}

var rotationDirection = map[string]func(value int){
	"L": leftRotate,
	"R": rightRotate,
}

func solve(instruction string) {
	direction := string(instruction[0])
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return
	}
	rotationDirection[direction](steps)
	if dial == 0 {
		count++
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		solve(line)
	}
	fmt.Println(count)
}
