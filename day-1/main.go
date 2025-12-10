package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		fmt.Printf("Error opening input.txt: %v\n", err)
		fmt.Println("Please create an input.txt file with the rotation instructions.")
		return
	}
	defer file.Close()

	position := 50
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing distance in line '%s': %v\n", line, err)
			continue
		}

		if direction == 'L' {
			position = (position - distance + 100) % 100
		}
		if direction == 'R' {
			position = (position + distance) % 100
		}

		if direction != 'L' && direction != 'R' {
			fmt.Printf("Invalid direction '%c' in line '%s'\n", direction, line)
			continue
		}

		if position == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("The password is: %d\n", count)
}
