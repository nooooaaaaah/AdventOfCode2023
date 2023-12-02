# Quick Thoughts

I had no problem with part 1 of the challenge. But I don't know how to use regex patterns well enough yet to use them to solve part 2. Nor do i even think it would've been the best way. Not that what I did was elegant. Part 2 was difficult for me cause I feel like at first I was my own worst enemy, Instead of thinking about the algorithm when my initial solution failed on a few edge case i tried to just handle them instead of reapraching how i was actually traversing the string. It took me going to bed and re-looking at my solution to see how I needed to traverse the string properly.

## Part 2 Solution

```go
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
```

### Explanation of my solution

So instead of just pasting the input text into my code I read it from a file cause it looked ugly and duh. Then had to do the classic loop da loop and process each line then character by character to get the first and last number then sum it all up. I used a builder to build up the string of numbers I was extracting from the input; along with a recursive function to process the string built from the builder. I used a for loop to find the number in the string and then added it to the builder and called the function again with the index set to the index of the last character of the number. If the string didn't contain a number I just added the character to the builder and called the function again with the index set to the next character. I used a switch statement to convert the words to numbers. I used a function to get the first and last number from the string of numbers. I used a function to extract the numbers from the input string.
