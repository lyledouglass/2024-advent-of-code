package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne() int {
	// Read the file and separate the numbers into two lists
	list1 := []int{}
	list2 := []int{}

	file, err := os.Open("list.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Find the differences between the lowest in list one and lowest in
	// list two, and continue with the next lowest in list one and list
	// two
	sort.Ints(list1)
	sort.Ints(list2)

	differences := []int{}
	for i := 0; i < len(list1); i++ {
		difference := int(math.Abs(float64(list1[i] - list2[i])))
		differences = append(differences, difference)
	}
	//fmt.Println(differences)

	// Add all the differences together
	sum := 0
	for _, difference := range differences {
		sum += difference
	}
	return sum

}

func partTwo() int {
	// Read the file and separate the numbers into two lists
	list1 := []int{}
	list2 := []int{}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Take the first number in list one and multiply it by the number of times it shows in list two
	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	result := 0
	for _, num := range list1 {
		result += num * counts[num]
	}
	return result
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
