package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	
)

func main() {
	file, err := os.Open("inputData.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var time []int
	var distance []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}
		
		nums := strings.Split(parts[1], "  ")
		
		for _, val := range nums {
			num, err := strconv.Atoi(strings.TrimSpace(val))
			if val == "" {
				continue
			}
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}

			if strings.Contains(parts[0], "Time") {
				time = append(time, num)
			} else if strings.Contains(parts[0], "Distance") {
				distance = append(distance, num)
			}
		}
		
	}

	// Print the arrays
	fmt.Println("distance =", distance)
	fmt.Println("time =", time)

	answerOne := partOne(time, distance)
	answerTwo := partTwo(time, distance)
	fmt.Println("Part 1 =", answerOne)
	fmt.Println("Part 2 =", answerTwo)

}


func partOne(time []int , distance []int) int {
	// Calculate the average speed
	waysToBeat := 1
	
	
	for i := 0; i < len(time); i++ {
		t := time[i]
		d := distance[i]
		j := 0
		 for (t- j) * j <= d {
			j += 1
		 }
		 waysToBeat *= (t- (2 * (j)))+1
		 }

	return waysToBeat
}

func intsToString(nums []int) string {
	var strNums []string
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}
	return concatStrings(strNums)
}

func concatStrings(strings []string) string {
	var result string
	for _, s := range strings {
		result += s
	}
	return result
}


func partTwo(time []int, distance []int) int {
	timeInt  := 0
	distanceInt := 0
	waysToBeat := 1
	timeStr := intsToString(time)
	distanceStr := intsToString(distance)

	timeInt, _ = strconv.Atoi(timeStr)
	distanceInt, _ = strconv.Atoi(distanceStr)

	
	j := 0
	for (timeInt- j) * j <= distanceInt {
		j += 1
	}
	waysToBeat *= (timeInt- (2 * (j)))+1
		
	return waysToBeat

}
