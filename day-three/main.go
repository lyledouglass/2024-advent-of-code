package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var matchList []string

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		inputString := strings.Fields(line)
		joinedString := strings.Join(inputString, " ")
		matches := pattern.FindAllStringSubmatch(joinedString, -1)

		for _, match := range matches {
			if len(match) == 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				totalSum += num1 * num2
			}
		}

	}
	return totalSum
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)
	totalSum := 0
	mulEnabled := true

	for scanner.Scan() {
		line := scanner.Text()
		pos := 0

		for pos < len(line) {
			doMatch := doPattern.FindStringIndex(line[pos:])
			dontMatch := dontPattern.FindStringIndex(line[pos:])
			mulMatch := mulPattern.FindStringSubmatchIndex(line[pos:])

			if doMatch != nil && (dontMatch == nil || doMatch[0] < dontMatch[0]) && (mulMatch == nil || doMatch[0] < mulMatch[0]) {
				mulEnabled = true
				pos += doMatch[1]
				// fmt.Println("do() found, enabling mul at position", pos)
			} else if dontMatch != nil && (doMatch == nil || dontMatch[0] < doMatch[0]) && (mulMatch == nil || dontMatch[0] < mulMatch[0]) {
				mulEnabled = false
				pos += dontMatch[1]
				// fmt.Println("don't() found, disabling mul at position", pos)
			} else if mulMatch != nil && (doMatch == nil || mulMatch[0] < doMatch[0]) && (dontMatch == nil || mulMatch[0] < dontMatch[0]) {
				if mulEnabled {
					num1, _ := strconv.Atoi(line[pos+mulMatch[2] : pos+mulMatch[3]])
					num2, _ := strconv.Atoi(line[pos+mulMatch[4] : pos+mulMatch[5]])
					totalSum += num1 * num2
					// fmt.Printf("mul(%d,%d) found and processed at position %d\n", num1, num2, pos)
				}
				pos += mulMatch[1]
			} else {
				break
			}
		}
	}
	return totalSum
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
