package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)
	splitInput := strings.Split(input, "\n")
	fmt.Println(sumOfParts(splitInput))
}

func sumOfParts(input []string) int {
	rows := len(input)
	cols := len(input[0])
	total := 0
	var partNumberStr strings.Builder

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if checkForSymbol(input[i], j) {
				partNumber := getSurroundingDigits(input, i, j, rows, cols)
				fmt.Println(partNumber)
				total += partNumber
				partNumberStr.Reset()
			}
		}
	}
	return total
}

func checkForSymbol(input string, index int) bool {
	// fmt.Println(input)
	symbols := map[rune]bool{'*': true, '#': true, '+': true, '$': true}
	if _, ok := symbols[rune(input[index])]; ok {
		// fmt.Printf("Found symbol: %c", input[index])
		return true
	}
	return false
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func getSurroundingDigits(input []string, i, j, rows, cols int) int {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	var results []string
	for k := 0; k < 8; k++ {
		nx, ny := i+dx[k], j+dy[k]
		if 0 <= nx && nx < rows && 0 <= ny && ny < cols && isDigit(rune(input[nx][ny])) {
			digits := string(input[nx][ny])
			for 0 <= nx+dx[k] && nx+dx[k] < rows && 0 <= ny+dy[k] && ny+dy[k] < cols && isDigit(rune(input[nx+dx[k]][ny+dy[k]])) {
				nx, ny = nx+dx[k], ny+dy[k]
				digits += string(input[nx][ny])
			}
			results = append(results, digits)
		}
	}
	// convert strings to integers and sum them up
	resultsSums := sumList(convertStringsToInts(results))
	return resultsSums
}

// convert strings to integers
func convertStringsToInts(input []string) []int {
	output := make([]int, len(input))
	for i, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to convert string to integer: %v", err)
		}
		output[i] = num
	}
	return output
}

// sum list of ints
func sumList(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total
}
