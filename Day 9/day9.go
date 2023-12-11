package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diff(seq []int) ([]int, bool) {
	diffed := make([]int, len(seq)-1)
	for i := range diffed {
		diffed[i] = seq[i+1] - seq[i]
	}
	// fmt.Println("seq: ", seq)
	// fmt.Println("Diffed: ", diffed)
	return diffed, !allMatch(seq, func(i int) bool { return i == 0 })
}

func extrapolate(diffs [][]int) int {
	extrapolated := 0
	for _, seq := range diffs {
		extrapolated = extrapolated + seq[len(seq)-1]
	}
	return extrapolated
}

func allMatch[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
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

	sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		seq := make([]int, 0)
		for _, s := range strings.Fields(line) {
			i, _ := strconv.Atoi(s)
			seq = append(seq, i)
		}

		diffed := seq
		canDiff := true
		diffs := make([][]int, 0)
		for canDiff {
			diffed, canDiff = diff(diffed)
			// fmt.Println("diffed: ", diffed, "canDiff: ", canDiff)
			if canDiff {
				diffs = append([][]int{diffed}, diffs...)
			}
		}

		extrapolated := extrapolate(append(diffs, seq))

		fmt.Println(seq, extrapolated)
		for _, diff := range diffs {
			fmt.Println(diff)
		}
		fmt.Println()
		sum += extrapolated
	}

	fmt.Println(sum)
}
