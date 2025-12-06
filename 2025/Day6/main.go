package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var total = 0
var cephalopodTotal = 0

func multipy(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result *= num
	}
	return result
}
func add(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}
func splitAndRemoveEmpty(input string, splitter string) []string {
	processedInput := strings.Split(input, splitter)
	processedInput = slices.DeleteFunc(processedInput, func(s string) bool {
		return s == ""
	})
	return processedInput
}
func solve(numbers [][]int, operations []string) {
	for i := range numbers {
		result := 0
		switch operations[i] {
		case "+":
			result += add(numbers[i])
		case "*":
			result += multipy(numbers[i])
		}
		total += result
	}

}

func mapToMatrix(input []string) {
	matrix := make([][]string, len(input))
	for i := range input {
		matrix[i] = make([]string, len(input[i]))
	}
	for i := range input {
		for j := range input[i] {
			matrix[i][j] = string(input[i][j])
		}
	}
	convertToCephalopod(matrix)
}

func convertToCephalopod(matrix [][]string) {
	numbersArray := []int{}
	for y := len(matrix[0]) - 1; y >= 0; y-- {
		strNumber := ""
		operation := ""
		for x := range matrix {
			// if matrix[x][y] == "*" || matrix[x][y] == "+" {
			// 	continue
			// }
			strNumber += matrix[x][y]
		}
		if strings.Contains(strNumber, "*") {
			strNumber = strNumber[:len(strNumber)-1]
			operation = "*"
		}
		if strings.Contains(strNumber, "+") {
			strNumber = strNumber[:len(strNumber)-1]
			operation = "+"
		}
		number, err := strconv.Atoi(strings.Trim(strNumber, " "))
		if err != nil {
			continue
		}
		numbersArray = append(numbersArray, number)
		if operation == "*" {
			cephalopodTotal += multipy(numbersArray)
			numbersArray = slices.Delete(numbersArray, 0, len(numbersArray))
		}
		if operation == "+" {
			cephalopodTotal += add(numbersArray)
			numbersArray = []int{}
		}
	}
}

func main() {
	file, _ := os.ReadFile("input.txt")
	rawInput := splitAndRemoveEmpty(string(file), "\n")
	mapToMatrix(rawInput)
	w := len(splitAndRemoveEmpty(rawInput[0], " "))
	h := len(rawInput) - 1
	numbers := make([][]int, w)
	for i := range numbers {
		numbers[i] = make([]int, h)
	}
	operations := make([]string, w)
	for i, s := range rawInput {
		line := strings.Split(s, " ")
		line = slices.DeleteFunc(line, func(s string) bool {
			return s == ""
		})
		for j := range line {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				operations[j] = line[j]
				continue
			}
			numbers[j][i] = num
		}
	}
	solve(numbers, operations)
	fmt.Println("Normal:", total)
	fmt.Println("Cephalopod:", cephalopodTotal)
}
