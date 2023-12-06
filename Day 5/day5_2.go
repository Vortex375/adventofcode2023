package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	start  int
	length int
}

type SeedMap struct {
	name string
	maps []MapDef
}

type MapDef struct {
	name        string
	destination int
	source      int
	length      int
}

func (s *SeedMap) mapSeed(v int) int {
	result := v
	for _, m := range s.maps {
		if v >= m.source && v < m.source+m.length {
			result = v - m.source + m.destination
		}
	}

	// fmt.Printf("%s %d -> %d \n", s.name, v, result)

	return result
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

	// parse seeds
	scanner.Scan()
	seedLine := scanner.Text()
	seedsStr := strings.Split(seedLine, " ")[1:]
	seeds := make([]Seed, 0)
	for i := 0; i < len(seedsStr)-1; i += 2 {
		start, _ := strconv.Atoi(seedsStr[i])
		length, _ := strconv.Atoi(seedsStr[i+1])
		seed := Seed{start, length}
		seeds = append(seeds, seed)
	}

	seedMaps := make([]SeedMap, 0, 7)
	var currentMap *SeedMap = nil
	// collect all maps
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map") {
			if currentMap != nil {
				seedMaps = append(seedMaps, *currentMap)
			}
			currentMap = &SeedMap{line, make([]MapDef, 0)}
			continue
		}
		if currentMap != nil {
			mapDef := MapDef{}
			fmt.Sscanf(line, "%d %d %d", &(mapDef.destination), &(mapDef.source), &(mapDef.length))
			currentMap.maps = append(currentMap.maps, mapDef)
		}
	}
	if currentMap != nil {
		seedMaps = append(seedMaps, *currentMap)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		fmt.Printf("Seed: %d\n", seed)
		for i := seed.start; i < seed.start+seed.length; i++ {
			// fmt.Println(i)
			v := i
			for _, seedMap := range seedMaps {
				v = seedMap.mapSeed(v)
			}
			if v < lowest {
				lowest = v
			}
		}
		fmt.Printf("lowest: %d\n", lowest)
		fmt.Println()
	}

	fmt.Println(lowest)

}
