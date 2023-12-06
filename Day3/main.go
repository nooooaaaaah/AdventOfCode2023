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
				partNumberStr.WriteString(getSurroundingDigits(input, i, j, rows, cols))
				fmt.Println(partNumberStr.String())
				partNumber, err := strconv.Atoi(partNumberStr.String())
				if err != nil {
					log.Fatal(err)
				}
				total += partNumber
				partNumberStr.Reset()
			}
		}
	}
	return total
}

func checkForSymbol(input string, index int) bool {
	symbols := map[rune]bool{'*': true, '#': true, '+': true, '$': true}
	if _, ok := symbols[rune(input[index])]; ok {
		return true
	}
	return false
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func getSurroundingDigits(input []string, i, j, rows, cols int) string {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	var builder strings.Builder

	for k := 0; k < 8; k++ {
		nx, ny := i+dx[k], j+dy[k]
		if 0 <= nx && nx < rows && 0 <= ny && ny < cols && isDigit(rune(input[nx][ny])) {
			builder.WriteRune(rune(input[nx][ny]))
		}
	}

	return builder.String()
}
