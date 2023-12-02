package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(getHashesFromFile())
}

func getHashesFromFile() int {
	f, err := openFile("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return processLines(f)
}

func openFile(filename string) (*os.File, error) {
	return os.Open(filename)
}

func processLines(f *os.File) int {
	var total int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		total += processLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func processLine(line string) int {
	num := extractNumbers(line)
	fmt.Println("extracted: ", num)
	return getFirstAndLast(num)
}

func extractNumbers(input string) string {
	var builder strings.Builder
	processString(input, &builder, 0)
	return builder.String()
}

func processString(input string, builder *strings.Builder, index int) {
	if index >= len(input) {
		return
	}

	longestNum := ""
	for length := 1; length <= len(input)-index; length++ {
		substr := input[index : index+length]
		if num, ok := wordToNumber(substr); ok {
			if len(num) > len(longestNum) {
				longestNum = num
			}
		}
	}

	if longestNum != "" {
		builder.WriteString(longestNum)
		processString(input, builder, index+len(longestNum))
	} else if unicode.IsDigit(rune(input[index])) {
		builder.WriteRune(rune(input[index]))
		processString(input, builder, index+1)
	} else {
		processString(input, builder, index+1)
	}
}

func getFirstAndLast(num string) int {
	fl := num[0:1] + num[len(num)-1:]
	flInt, err := strconv.Atoi(fl)
	if err != nil {
		panic(err)
	}
	fmt.Println(flInt)
	return flInt
}

func wordToNumber(word string) (string, bool) {
	switch strings.ToLower(word) {
	case "one":
		return "1", true
	case "two":
		return "2", true
	case "three":
		return "3", true
	case "four":
		return "4", true
	case "five":
		return "5", true
	case "six":
		return "6", true
	case "seven":
		return "7", true
	case "eight":
		return "8", true
	case "nine":
		return "9", true
	default:
		return "", false
	}
}
