package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() string {

	var safeCount int
	var unsafeCount int

	// Read the file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		isSafe := false

		for i := 1; i < len(numbers); i++ {
			if checkSafety(numbers) {
				isSafe = true
			}
		}

		if isSafe {
			safeCount++
		} else {
			unsafeCount++
		}
	}
	return "Safe Reports: " + strconv.Itoa(safeCount) + " || Unsafe Reports: " + strconv.Itoa(unsafeCount)
}

func partTwo() string {
	var safeCount int
	var unsafeCount int

	// Read the file
	file, err := os.Open("input.txt")
	//file, err := os.Open("input_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		isSafe := false
		if checkSafety(numbers) {
			isSafe = true
		} else {
			for i := 0; i < len(numbers); i++ {
				tempList := append([]string{}, numbers[:i]...)
				tempList = append(tempList, numbers[i+1:]...)

				if checkSafety(tempList) {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			safeCount++
		} else {
			unsafeCount++
		}
	}
	return "Safe Reports: " + strconv.Itoa(safeCount) + " || Unsafe Reports: " + strconv.Itoa(unsafeCount)
}

func checkSafety(numbers []string) bool {
	numAsc := false
	numDesc := false

	for i := 1; i < len(numbers); i++ {
		prev, _ := strconv.Atoi(numbers[i-1])
		curr, _ := strconv.Atoi(numbers[i])

		// Check if all numbers are ascending or descending, but not both.
		// Also, if the numbers are the same, it's unsafe
		if curr > prev {
			numAsc = true
		} else if curr < prev {
			numDesc = true
		}
		if numAsc && numDesc || curr == prev {
			return false
		}

		// Check the difference between the numbers, seeing if its between -3 and 3
		diff := curr - prev
		if diff < -3 || diff > 3 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Part One:")
	fmt.Println(partOne())
	fmt.Println("Part Two:")
	fmt.Println(partTwo())
}
