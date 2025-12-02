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
var passes = 0

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}
func leftRotate(value int) {
	passes += (mod(-dial, 100) + value) / 100
	dial -= value
	dial = mod(dial, 100)
}
func rightRotate(value int) {
	passes += (mod(dial, 100) + value) / 100
	dial += value
	dial = mod(dial, 100)
}

var rotationDirection = map[string]func(value int){
	"L": leftRotate,
	"R": rightRotate,
}

func solve(instruction string) {
	direction := string(instruction[0])
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		fmt.Println("Error")
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
	fmt.Println("Count:\t", count)
	fmt.Println("Passes:\t", passes)
}
