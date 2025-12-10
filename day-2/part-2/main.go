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

	for segmentLen := 1; segmentLen <= length/2; segmentLen++ {
		if length%segmentLen != 0 {
			continue
		}

		firstSegment := idStr[:segmentLen]
		numSegments := length / segmentLen

		allEqual := true
		for i := 1; i < numSegments; i++ {
			segment := idStr[i*segmentLen : (i+1)*segmentLen]
			if segment != firstSegment {
				allEqual = false
				break
			}
		}

		if allEqual && numSegments >= 2 {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("day-2/part-2/input.txt")
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
