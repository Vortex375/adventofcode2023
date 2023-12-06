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

	var previous *string
	var current *string
	var next *string
	var sum int
	for i := range lines {
		fmt.Printf("Line %d: ", i)
		if i > 0 {
			previous = &lines[i-1]
		} else {
			previous = nil
		}
		current = &lines[i]
		if i < len(lines)-1 {
			next = &lines[i+1]
		} else {
			next = nil
		}

		var currentNumber func() Number = nil
		var putDigit func(rune, int) = nil
		for index, r := range *current {
			if unicode.IsDigit(r) {
				if currentNumber == nil {
					putDigit, currentNumber = numberCollector()
				}
				putDigit(r, index)
			} else {
				if currentNumber != nil {
					number := currentNumber()
					fmt.Printf("%d adjacent: ", number.value)
					if checkAdjacent(number, previous, current, next) {
						sum += number.value
						fmt.Print("true; ")
					} else {
						fmt.Print("false; ")
					}
					putDigit, currentNumber = nil, nil
				}
			}
		}
		if currentNumber != nil {
			number := currentNumber()
			fmt.Printf("%d adjacent: ", number)
			if checkAdjacent(number, previous, current, next) {
				sum += number.value
				fmt.Print("true; ")
			} else {
				fmt.Print("false; ")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println(sum)
}
