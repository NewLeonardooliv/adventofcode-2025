package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)

	if length%2 != 0 {
		return false
	}

	half := length / 2
	firstHalf := idStr[:half]
	secondHalf := idStr[half:]

	return firstHalf == secondHalf
}

func main() {
	file, err := os.Open("day-2/part-1/input.txt")
	if err != nil {
		fmt.Printf("Error opening input.txt: %v\n", err)
		fmt.Println("Please create an input.txt file with the product ID ranges.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line != "" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if line == "" {
		fmt.Println("No input found in file")
		return
	}

	ranges := strings.Split(line, ",")
	sum := 0

	for _, rangeStr := range ranges {
		rangeStr = strings.TrimSpace(rangeStr)
		if rangeStr == "" {
			continue
		}

		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			fmt.Printf("Invalid range format: %s\n", rangeStr)
			continue
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error parsing start of range '%s': %v\n", rangeStr, err)
			continue
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error parsing end of range '%s': %v\n", rangeStr, err)
			continue
		}

		for id := start; id <= end; id++ {
			if isInvalidID(id) {
				sum += id
			}
		}
	}

	fmt.Printf("Sum of all invalid IDs: %d\n", sum)
}
