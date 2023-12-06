package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line
	var numbers []int64
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		var first, last string
		for index, char := range line {
			var digit string
			switch {
			case unicode.IsDigit(char):
				digit = string(char)
			case index <= len(line)-3 && string(line[index:index+3]) == "one":
				digit = "1"
			case index <= len(line)-3 && string(line[index:index+3]) == "two":
				digit = "2"
			case index <= len(line)-5 && string(line[index:index+5]) == "three":
				digit = "3"
			case index <= len(line)-4 && string(line[index:index+4]) == "four":
				digit = "4"
			case index <= len(line)-4 && string(line[index:index+4]) == "five":
				digit = "5"
			case index <= len(line)-3 && string(line[index:index+3]) == "six":
				digit = "6"
			case index <= len(line)-5 && string(line[index:index+5]) == "seven":
				digit = "7"
			case index <= len(line)-5 && string(line[index:index+5]) == "eight":
				digit = "8"
			case index <= len(line)-4 && string(line[index:index+4]) == "nine":
				digit = "9"
			}
			if digit == "" {
				continue
			}
			if first == "" {
				first = digit
			}
			last = digit
		}

		number, _ := strconv.ParseInt(first+last, 10, 64)
		numbers = append(numbers, number)
		fmt.Printf("%s %d\n", line, number)
	}

	var sum int64 = 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
