package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	// Open the file
	file, err := os.Open("mockData.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize a map to store the card data
	cardData := make(map[int][][]int)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line based on the ":" symbol
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue // Skip invalid lines
		}
		gameNumberStr := strings.TrimPrefix(parts[0], "Card ")
		// Extract the card number
		cardNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			fmt.Println("Error converting card number:", err)
			continue
		}

		// Split the second part based on the "|" symbol
		cardValues := strings.Split(parts[1], "|")

		// Process the values before the "|"
		beforePipe := parseValues(cardValues[0])

		// Process the values after the "|"
		afterPipe := parseValues(cardValues[1])

		// Store the data in the map
		cardData[cardNumber] = [][]int{beforePipe, afterPipe}
	}

	// Print the resulting dictionary
	// fmt.Println(cardData)

	answerOne := partOne(cardData)
	answerTwo := partTwo(cardData)
	fmt.Println("the answer for part one is: ", answerOne)
	fmt.Println("the answer for part two is: ", answerTwo)
}

// Helper function to parse a string of space-separated numbers into a 1D slice of integers
func parseValues(input string) []int {
	values := strings.Fields(input)
	result := make([]int, len(values))

	for i, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting value:", err)
			continue
		}
		result[i] = num
	}

	return result
}


func partOne( data map[int][][]int) int{
	totalSum := 0
	for _, values := range data{
		winningNumbers := values[0]
		recievedNumbers := values[1]

		temp := 0
		for _, num := range winningNumbers{
			for _, numb := range recievedNumbers{
				if num == numb{
					if temp == 0{
						temp += 1
					}else{
						temp *= 2
					}
					break
				}
			}
		}
		totalSum += temp
	}

	return totalSum
}

func partTwo( data map[int][][]int) int {
	scratchCards := 0
	copiesDict := make(map[int]int)
	var sortedCards []int

	for card := range data{
		sortedCards = append(sortedCards, card)
	}
	sort.Ints(sortedCards)
	for _, card := range sortedCards{
		values := data[card]
		winningNumbers := values[0]
		recievedNumbers := values[1]

		temp := 0
		for _, num := range winningNumbers{
			for _, numb := range recievedNumbers{
				if num == numb{
					temp += 1
				}
			}
		}
		// fmt.Println(copiesDict, temp)
		copiesDict[card] += 1
		j :=0
		for j < copiesDict[card]{
			i := 1
			for i <= temp{
				copiesDict[card+i] += 1
				i += 1
			}
			j += 1
		}
		// fmt.Println(copiesDict, temp)
		
		scratchCards += copiesDict[card]
	}
	// fmt.Println(copiesDict)
	return scratchCards
}
