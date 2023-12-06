package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var possibleGames []int64
	var sumOfPower int64

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		gamesplit := strings.Split(line, ":")
		gameId, _ := strconv.ParseInt(strings.TrimSpace(gamesplit[0])[5:], 10, 64)

		cubes := make(map[string]int64)
		samplesplit := strings.Split(gamesplit[1], ";")
		for _, samplestr := range samplesplit {
			cubesplit := strings.Split(samplestr, ",")
			for _, cubestr := range cubesplit {
				var amount int64
				var key string
				fmt.Sscanf(cubestr, "%d %s", &amount, &key)
				if amount > cubes[key] {
					cubes[key] = amount
				}
			}
		}
		fmt.Print(line)
		fmt.Printf(" | Game %d ", gameId)
		for key, value := range cubes {
			fmt.Printf("%s %d ", key, value)
		}
		fmt.Println("")
		totalRed := cubes["red"]
		totalGreen := cubes["green"]
		totalBlue := cubes["blue"]
		if totalRed <= 12 && totalGreen <= 13 && totalBlue <= 14 {
			possibleGames = append(possibleGames, gameId)
			fmt.Printf(" | possible : true\n")
		} else {
			fmt.Printf(" | possible : false\n")
		}
		sumOfPower += totalRed * totalGreen * totalBlue
	}

	var sum int64
	for _, g := range possibleGames {
		sum += g
	}

	fmt.Println(sum)
	fmt.Println(sumOfPower)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
