package main

import(
	"bufio"
	"strings"
	"fmt"
	"os"
	
)

func main(){

	filePath := "inputData.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	var instructions string
	dictionary := make(map[string][]string)
	startingPoint := make([]string, 0)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the file into instructions and dictionary parts
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
	
			if key[len(key)-1] == 'A'{
				startingPoint = append(startingPoint, key)
			}
			value := strings.TrimSpace(parts[1])
			value = strings.Trim(value, "()")
			
			array := make([]string, 0)
			for _, v := range strings.Split(value, ", ") {
				newV := strings.TrimSpace(v)
				
				array = append(array, newV)
			}
			

			dictionary[key] = array
		} else {
			if len(line) > 0 {

				instructions = line
			}
		}
	}

	// Process instructions
	// fmt.Println("Instructions:", instructions)
	// fmt.Println("Starting Point:", startingPoint)

	// // Process dictionary
	// fmt.Println("Dictionary:")
	// for key, value := range dictionary {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	// Part 1
	// answerOne := partOne(instructions, dictionary)
	// fmt.Println("part 1: ", answerOne)

	// Part 2
	answerTwo := partTwo(instructions, dictionary, startingPoint)
	fmt.Println("part 2: ", answerTwo)
	
}

func partOne(instructions string, dictionary map[string][]string) int {
		
	steps := 0
	i := 0
	curPos := "AAA"

	for curPos != "ZZZ"{
		
		if instructions[i] == 'L' {
			curPos = dictionary[curPos][0]
		} else {

			curPos = dictionary[curPos][1]
		}

		steps++
		i++

		if i == len(instructions) {
			i = 0
		}
		fmt.Println(i)
	}
	return steps
}

func partTwo(instructions string, dictionary map[string][]string, startingPoint []string) int {
	steps := 0
	curPos := startingPoint
	finish := false
	i := 0

	temp := make([]string, len(curPos)) // Initialize temp outside the loop

	for !finish {
		
		tempFinish := true
		directionToMove := instructions[i]

		for indice, pos := range curPos {
			

			var nexPos string
			if directionToMove == 'L' {
				nexPos = dictionary[pos][0]
			} else {
				nexPos = dictionary[pos][1]
			}

			if nexPos[len(nexPos)-1] != 'Z' {
				tempFinish = false
			}
			temp[indice] = nexPos
		}

		if tempFinish {
			finish = true
		}

		copy(curPos, temp) // Copy temp to curPos

		steps++
		i++
	
		if i == len(instructions) {
			
			i = 0
		}
	}

	return steps
}
