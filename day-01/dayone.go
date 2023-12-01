package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Node struct {
	children map[rune]*Node
	word     string
}

func newNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		word:     "",
	}
}

type Trie struct {
	root *Node
}

func newTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) insert(word string) {
	node := t.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = newNode()
		}
		node = node.children[c]
	}
	node.word = word
}

func (t *Trie) find(idx int, str string) string {
	curNode := t.root
	for i := idx; i < len(str); i++ {
		c := rune(str[i])

		if _, ok := curNode.children[c]; !ok {
			return ""
		}

		curNode = curNode.children[c]

		if curNode.word != "" {
			return curNode.word
		}
	}

	return ""
}

func main() {
	realData := "inputData.txt"
	// mockData := "mockData.txt"
	// mockData2 := "mockData2.txt"
	answerOne := partOne(realData)
	answerTwo := partTwo(realData)
	fmt.Println(answerOne, answerTwo)
}

func partOne(filePath string) int {
	totalSum := 0

	content, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()

		var nums []int

		for _, char := range line {
			if unicode.IsDigit(char) {
				num := int(char - '0')
				nums = append(nums, num)
			}
		}
		if len(nums) > 0 {
			totalSum += nums[0]*10 + nums[len(nums)-1]
		}
	}

	fmt.Println(totalSum)

	return totalSum
}

func partTwo(filePath string) int {
	totalSum := 0

	numberMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	content, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	trie := newTrie()
	for _, num := range numbers {
		trie.insert(num)
	}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()

		firstNum, firstIdx, lastNum, lastIdx := parseInput(line, trie)
		firstValue, firstFound := numberMap[firstNum]
		lastValue, lastFound := numberMap[lastNum]

		if !firstFound {
			firstValue = -1
		}
		if !lastFound {
			lastValue = -1
		}
		var nums [][]int
		
		for i, char := range line {
			if unicode.IsDigit(char) {
				num := int(char - '0')
				nums = append(nums, []int{num, i})
				
			}
		}
		if firstValue == -1 && lastValue == -1 {
			firstValue = nums[0][0]
			
			lastValue = nums[len(nums)-1][0]
			

		}else if firstValue == -1 && lastValue != -1{
			
			firstValue = nums[0][0]
			
			if lastIdx < nums[len(nums)-1][1] {
				lastValue = nums[len(nums)-1][0]
			}
			
		}else if firstValue != -1 && lastValue == -1{
			
			if firstIdx > nums[0][1] {
				firstValue = nums[0][0]
			}
			
			
			lastValue = nums[len(nums)-1][0]
			
		}else{
			
			if firstIdx > nums[0][1] {
				firstValue = nums[0][0]
			}
		
		
			
			if lastIdx < nums[len(nums)-1][1] {
				lastValue = nums[len(nums)-1][0]
			}
			
		}

		
		

		totalSum += firstValue*10 + lastValue
	}

	return totalSum
}

func parseInput(data string, trie *Trie) (string, int, string, int) {
	firstNum := ""
	firstIdx := -1
	lastNum := ""
	lastIdx := -1

	for idx, _ := range data {
		temp := trie.find(idx, data)
		if temp != "" {
			if firstNum == "" {
				firstNum = temp
				firstIdx = idx
			} else {
				lastNum = temp
				lastIdx = idx
			}
		}
		if temp != "" {
			if firstNum == "" {
				firstNum = temp
				firstIdx = idx
			} else {
				lastNum = temp
				lastIdx = idx
			}
		}
	}

	return firstNum, firstIdx, lastNum, lastIdx
}