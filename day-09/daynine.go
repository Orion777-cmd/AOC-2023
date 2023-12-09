package main;

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"

)

func main() {
	filePath := "inputData.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	defer file.Close()

	report := make([][]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		eachReport := strings.Split(line, " ")
		
		reportLine := make([]int, 0)

		for _, each := range eachReport {
			num, _ := strconv.Atoi(each)
			reportLine = append(reportLine, num)
		}

		report = append(report, reportLine)

	
	}

	// fmt.Println(report)

	//part 1
	// answerOne := partOne(report)
	// fmt.Println(answerOne)

	//part 2
	answerTwo := partTwo(report)
	fmt.Println(answerTwo)
}


func partOne(report [][]int) int {

	totalSum := 0

	for i, each := range report {
		// fmt.Println(each)
		curArray := report[i][:len(each)]
		// fmt.Println(curArray)
		lastNum := curArray[len(curArray)-1]

		for len(curArray) > 1{

			temp := make([]int, 0)
			for i := 0; i < len(curArray)-1; i++ {
				temp = append(temp, curArray[i+1] - curArray[i])
			}
			lastNum += temp[len(temp)-1]

			curArray = temp

		}
		// fmt.Println(curArray, lastNum)
		
		totalSum += lastNum 
 

	}
	
	return totalSum
}


func partTwo (report [][]int ) int {
	totalSum :=0

	
	

	for i, each:= range report {
		curArray := report[i][:len(each)]
	
		
		firstValues := make([]int , 0)
		
		for len(curArray) > 1{
			firstValues = append(firstValues, curArray[0])
			
			temp := make([]int, 0)
			for i := 0; i < len(curArray)-1; i++ {
				temp = append(temp, curArray[i+1] - curArray[i])
			}
	
			curArray = temp
		}
		
		lastNum := 0
		for i:= len(firstValues)-1; i >= 0; i-- {
			
			firstValues[i] = firstValues[i] - lastNum
			lastNum = firstValues[i]
		
		}
		
		totalSum += firstValues[0]
	}

	return totalSum
}