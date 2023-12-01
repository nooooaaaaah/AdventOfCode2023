package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// list of hashes
// extract numbers from hashes
// sum numbers

func main() {
	fmt.Println(getHashesFromFile())
}

func getHashesFromFile() int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines int
	re := regexp.MustCompile(`\D`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num := re.ReplaceAllString(scanner.Text(), "")
		fl := (num[0:1] + num[len(num)-1:])
		flInt, err := strconv.Atoi(fl)
		if err != nil {
			panic(err)
		}
		fmt.Println(fl)
		lines += flInt
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
