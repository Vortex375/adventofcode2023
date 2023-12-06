package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxDistance(buttonTime int, totalTime int) int {
	return buttonTime * (totalTime - buttonTime)
}

func main() {
	// Open the file
	file, err := os.Open("input2")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	lineSplit := strings.Fields(line)[1:]
	fmt.Println(lineSplit)
	times := make([]int, len(lineSplit))
	for i, str := range lineSplit {
		times[i], _ = strconv.Atoi(str)
	}

	scanner.Scan()
	line = scanner.Text()
	lineSplit = strings.Fields(line)[1:]
	fmt.Println(lineSplit)
	distances := make([]int, len(lineSplit))
	for i, str := range lineSplit {
		distances[i], _ = strconv.Atoi(str)
	}

	result := 1
	for i := range times {
		time := times[i]
		distance := distances[i]
		success := 0
		for j := 1; j < time; j++ {
			d := maxDistance(j, time)
			// fmt.Printf("Time %d Distance %d Press %d Total %d Win %t\n", time, distance, j, d, d > distance)
			if d > distance {
				success++
			}
		}
		// fmt.Printf("Success: %d\n", success)
		if success > 0 {
			result *= success
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(result)
}
