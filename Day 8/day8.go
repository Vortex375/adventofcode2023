package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair[T any] [2]T

func (p *Pair[T]) Left() T {
	return p[0]
}

func (p *Pair[T]) Right() T {
	return p[1]
}

func (p *Pair[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p[0], p[1])
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

	scanner.Scan()
	path := strings.TrimSpace(scanner.Text())
	maps := make(map[string]*Pair[string])
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		key := line[0:3]
		left := line[7:10]
		right := line[12:15]

		maps[key] = &Pair[string]{left, right}
	}

	fmt.Println(path)
	fmt.Println(maps)

	dist := 0
	currentKey := "AAA"
	found := false
	for !found {
		for _, direction := range path {
			dist++
			if direction == 'L' {
				currentKey = maps[currentKey].Left()
			} else {
				currentKey = maps[currentKey].Right()
			}
			if currentKey == "ZZZ" {
				found = true
				break
			}
		}
	}

	fmt.Println(dist)
}
