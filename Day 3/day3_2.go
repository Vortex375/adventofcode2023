package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	value int
	first int
	last  int
}

func numberCollector() (func(rune, int), func() Number) {
	var value string
	var first int
	var last int

	putDigit := func(digit rune, pos int) {
		if value == "" {
			first = pos
		}
		last = pos
		value += string(digit)
	}

	getValue := func() Number {
		v, _ := strconv.Atoi(value)
		return Number{v, first, last}
	}

	return putDigit, getValue
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAdjacentToSymbol(number Number, line string) bool {
	s := line[max(0, number.first-1):min(len(line), number.last+2)]

	for _, r := range s {
		if isSymbol(r) {
			return true
		}
	}
	return false
}

func isSymbol(r rune) bool {
	if unicode.IsDigit(r) {
		return false
	} else if r == '.' {
		return false
	}
	return true
}

func checkAdjacent(number Number, lines ...*string) bool {
	for _, s := range lines {
		if s == nil {
			continue
		}
		if isAdjacentToSymbol(number, *s) {
			return true
		}
	}
	return false
}

func adjacentNumbers(index int, numbers ...[]Number) []Number {
	ret := make([]Number, 0)
	for _, numberLine := range numbers {
		if numberLine == nil {
			continue
		}
		// fmt.Println(numberLine)
		for _, number := range numberLine {
			if index >= number.first-1 && index <= number.last+1 {
				// fmt.Printf("adjacent: %d, %d\n", index, number)
				ret = append(ret, number)
			}
		}
	}
	return ret
}

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

	lines := make([]string, 0)
	// collect all lines
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	numbers := make([][]Number, 0)
	var currentNumber func() Number = nil
	var putDigit func(rune, int) = nil
	for _, line := range lines {
		numbersOnLine := make([]Number, 0)
		for index, r := range line {
			if unicode.IsDigit(r) {
				if currentNumber == nil {
					putDigit, currentNumber = numberCollector()
				}
				putDigit(r, index)
			} else {
				if currentNumber != nil {
					number := currentNumber()
					numbersOnLine = append(numbersOnLine, number)
					putDigit, currentNumber = nil, nil
				}
			}
		}
		if currentNumber != nil {
			number := currentNumber()
			numbersOnLine = append(numbersOnLine, number)
			putDigit, currentNumber = nil, nil
		}
		numbers = append(numbers, numbersOnLine)
	}

	var previous []Number
	var current []Number
	var next []Number
	var sum int
	for i, line := range lines {
		if i > 0 {
			previous = numbers[i-1]
		} else {
			previous = nil
		}
		current = numbers[i]
		if i < len(numbers)-1 {
			next = numbers[i+1]
		} else {
			next = nil
		}

		for index, r := range line {
			if r == '*' {
				adjacent := adjacentNumbers(index, previous, current, next)
				if len(adjacent) == 2 {
					// fmt.Printf("Found Gear on line %d at %d, adjacent: %d, %d\n", i, index, adjacent[0].value, adjacent[1].value)
					sum += adjacent[0].value * adjacent[1].value
					// sum += 1
				}
			}
		}
	}

	fmt.Println(sum)
}
