package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var (
		answer int
	)

	file, err := os.Open("input.txt")
 
	if err != nil {
		fmt.Printf("failed opening input file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input []int
 
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	// Part 1

	// prevInput := 0
	// for idx, i := range input {
	// 	if idx == 0 {
	// 		prevInput = i
	// 	} else {
	// 		if i > prevInput {
	// 			answer++
	// 		}
	// 		prevInput = i
	// 	}
	// }

	// Part 2

	prevInput := 0
	for idx, i := range input {
		fmt.Println(idx, i)
		if prevInput == -1 {
			break
		}
		if idx == 0 {
			prevInput = getWindowInput(idx, input)
		} else {
			if getWindowInput(idx, input) > prevInput {
				answer++
			}
			prevInput = getWindowInput(idx, input)
		}
	}

	fmt.Printf("Answer: %d\n", answer)
}

func getWindowInput(idx int, input []int) int {
	// fmt.Println(len(input)-3, idx)
	if len(input) - 3 < idx {
		return -1
	}
	fmt.Println(input[idx] + input[idx+1] + input[idx+2])
	return input[idx] + input[idx+1] + input[idx+2]
}