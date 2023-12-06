package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	numbers := make([][]int, 0)
	winningNumbers := make([]map[int]bool, 0)
	// collect all lines
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		var card int
		currentNumbers := make([]int, 10)
		currentWinningNumbers := make([]int, 25)
		currentWinningNumbersMap := make(map[int]bool)
		fmt.Sscanf(line, "Card %d: %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d | %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d %2d",
			&card,
			&currentNumbers[0],
			&currentNumbers[1],
			&currentNumbers[2],
			&currentNumbers[3],
			&currentNumbers[4],
			&currentNumbers[5],
			&currentNumbers[6],
			&currentNumbers[7],
			&currentNumbers[8],
			&currentNumbers[9],
			&currentWinningNumbers[0],
			&currentWinningNumbers[1],
			&currentWinningNumbers[2],
			&currentWinningNumbers[3],
			&currentWinningNumbers[4],
			&currentWinningNumbers[5],
			&currentWinningNumbers[6],
			&currentWinningNumbers[7],
			&currentWinningNumbers[8],
			&currentWinningNumbers[9],
			&currentWinningNumbers[10],
			&currentWinningNumbers[11],
			&currentWinningNumbers[12],
			&currentWinningNumbers[13],
			&currentWinningNumbers[14],
			&currentWinningNumbers[15],
			&currentWinningNumbers[16],
			&currentWinningNumbers[17],
			&currentWinningNumbers[18],
			&currentWinningNumbers[19],
			&currentWinningNumbers[20],
			&currentWinningNumbers[21],
			&currentWinningNumbers[22],
			&currentWinningNumbers[23],
			&currentWinningNumbers[24])

		for _, winningNumber := range currentWinningNumbers {
			currentWinningNumbersMap[winningNumber] = true
		}

		numbers = append(numbers, currentNumbers)
		winningNumbers = append(winningNumbers, currentWinningNumbersMap)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sum := 0
	for i := range numbers {
		currentNumbers := numbers[i]
		currentWinningNumbers := winningNumbers[i]

		value := 0
		for _, number := range currentNumbers {
			if _, wins := currentWinningNumbers[number]; wins {
				if value == 0 {
					value = 1
				} else {
					value = value * 2
				}
			}
		}

		sum += value
	}

	fmt.Println(sum)
}
