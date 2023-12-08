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

func allMatch[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(n ...int) int {
	result := n[0] * n[1] / gcd(n[0], n[1])

	for i := 2; i < len(n); i++ {
		result = lcm(result, n[i])
	}

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

	currentKeys := make([]string, 0)
	for key := range maps {
		if strings.Contains(key, "A") {
			currentKeys = append(currentKeys, key)
		}
	}
	fmt.Println(currentKeys)
	found := make([]bool, len(currentKeys))
	dist := make([]int, len(currentKeys))
	for !allMatch(found, func(b bool) bool { return b }) {
		for _, direction := range path {
			for i := range dist {
				if !found[i] {
					dist[i]++
				}
			}
			for i := range currentKeys {
				if direction == 'L' {
					currentKeys[i] = maps[currentKeys[i]].Left()
				} else {
					currentKeys[i] = maps[currentKeys[i]].Right()
				}
				if strings.Contains(currentKeys[i], "Z") {
					found[i] = true
				}
			}
		}
	}

	fmt.Println(dist)
	fmt.Println(lcm(dist...))
}
