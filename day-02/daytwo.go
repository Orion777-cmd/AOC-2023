package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)
func main(){
	// Open the file
	file, err := os.Open("inputData.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to store game values
	gameData := make(map[int][][]int)
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by colon
		parts := strings.SplitN(line, ":", 2)

		// Trim spaces from key and values
		gameNumberStr := strings.TrimPrefix(parts[0], "Game ")
		
		gameNumber, err := strconv.Atoi(gameNumberStr)

		if err != nil {
			fmt.Println("Error converting game number:", err)
			return
		}
		// Split values by semicolon and trim spaces
		colorMapper := map[string]int{
			"red": 0,
			"green": 1,
			"blue": 2,
		}
		
		values := strings.Split(parts[1], ";")
		gameValues := make([][]int, len(values))

		// Parse and count the occurrences of red, green, and blue cubes
		
		for i, v := range values {
			cubeCounts := []int{0, 0, 0} // red, green, blue

			// Split cube counts by comma and trim spaces
			counts := strings.Split(v, ",")
		
			for _, countStr := range counts {
				
				countStr = strings.TrimSpace(countStr)
				temp := strings.Split(countStr, " ")
				
				count, err := strconv.Atoi(temp[0])
				if err != nil {
					fmt.Println("Error converting cube count:", err)
					return
				}

				// Use the colorMapper to determine the correct index for cubeCounts
				color := strings.TrimSpace(temp[1])
				cubeCounts[colorMapper[color]] = count
			}

			// Store values in the map
			gameValues[i] = cubeCounts
		}

		// Store values in the map
		gameData[gameNumber] = gameValues
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the result
	for key, values := range gameData {
		fmt.Printf("%d: %v\n", key, values)
	}

	answerOne := partOne(gameData )
	answerTwo := partTwo(gameData)
	fmt.Println("The answer for the first part is : " , answerOne)
	fmt.Println("The answer for the second part is : ", answerTwo)
}


func partOne(parsedData map[int][][]int) int {
	red := 12
	green := 13
	blue := 14
	var id_sum int = 0
	sum_limit := red + green + blue

	for key, values := range parsedData {
		possible := true
		for _, value := range values {
			if value[0] > red || value[1] > green || value[2] > blue || value[0] + value[1] + value[2] > sum_limit {
				possible = false
			}
		}
		if possible {
			id_sum += key
		}
	}
	return id_sum
}



func partTwo(parsedData map[int][][]int) int {
	

	total_sum := 0
	for _, values := range parsedData {
		min_red := 0
		min_green := 0
		min_blue := 0
		for _, value := range values {
			if value[0] > min_red {
				min_red = value[0]
			}
			if value[1] > min_green {
				min_green = value[1]
			}
			if value[2] > min_blue {
				min_blue = value[2]
			}
		}
	
		total_sum += min_red * min_green * min_blue
	}

	return total_sum
}