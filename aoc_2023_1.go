package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "words.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Print(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var listOfNums []int
	for scanner.Scan() {
		numsOnly := findNumber(scanner.Text())
		readyNums := convertToInt(numsOnly)
		listOfNums = append(listOfNums, readyNums)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file", err)
	}
	fmt.Print(totalSum(listOfNums))
	defer file.Close()
}

func findNumber(words string) []int {
	numbers := make([]int, 0)
	for _, word := range words {
		if digit, err := strconv.Atoi(string(word)); err == nil {
			numbers = append(numbers, digit)
		}
	}
	firstElement := numbers[0]
	lastElement := numbers[len(numbers)-1]
	numbers = nil
	numbers = append(numbers, firstElement, lastElement)
	return numbers
}

func totalSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func convertToInt(numbers []int) int {
	if len(numbers) < 2 {
		return 0
	}
	return numbers[0]*10 + numbers[1]
}
