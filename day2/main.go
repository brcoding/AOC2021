package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	var (
		answer int
	)

	file, err := os.Open("testinput.txt")
 
	if err != nil {
		fmt.Printf("failed opening input file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input []string
 
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Part 1

	// position := 0
	// depth := 0

	// for _, command := range input {
	// 	commandParts := strings.Split(command, " ")
	// 	if len(commandParts) < 2 {
	// 		continue
	// 	}
	// 	number, err := strconv.Atoi(commandParts[1])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	switch (commandParts[0]) {
	// 		case "forward":
	// 			position += number
	// 		case "down":
	// 			depth += number
	// 		case "up":
	// 			depth -= number
	// 	}
	// }
	// answer = position * depth

	// Part 2

	position := 0
	depth := 0
	aim := 0

	for _, command := range input {
		commandParts := strings.Split(command, " ")
		if len(commandParts) < 2 {
			continue
		}
		number, err := strconv.Atoi(commandParts[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch (commandParts[0]) {
			case "forward":
				position += number
				depth += aim * number
			case "down":
				aim += number
			case "up":
				aim -= number
		}
	}
	answer = position * depth


	fmt.Printf("Answer: %d\n", answer)
}
