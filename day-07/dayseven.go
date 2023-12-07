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
	cardMap := map[string]int{
		"A" : 14,
		"K" : 13,
		"Q" : 12,
		"J" : 11,
		"T" : 10,
		"9" : 9,
		"8" : 8,
		"7" : 7,
		"6" : 6,
		"5" : 5,
		"4" : 4,
		"3" : 3,
		"2" : 2,
	
	}
	cardMapTwo := map[string]int{
		"A" : 14,
		"K" : 13,
		"Q" : 12,
		"T" : 10,
		"9" : 9,
		"8" : 8,
		"7" : 7,
		"6" : 6,
		"5" : 5,
		"4" : 4,
		"3" : 3,
		"2" : 2,
		"J" : 1,
	}
	filePath := "inputData.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	camelCardMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		score, err := strconv.Atoi(parts[1])

		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}

		camelCardMap[parts[0]] = score
	}

	answerOne := partOne(camelCardMap, cardMap)
	fmt.Println("part 1: ", answerOne)
	answerTwo := partTwo(camelCardMap, cardMapTwo)
	fmt.Println("part 2: ", answerTwo)
}

func partOne(cards map[string]int, cardMap map[string]int) int{
	totalScore := 0

	correctOrder := compareHands(cards, cardMap)

	for i, card := range correctOrder{
		totalScore += (i + 1) * cards[card]
	}

	return totalScore
}

func partTwo (cards map[string]int, cardMap map[string]int) int {
	totalScore := 0

	correctOrder := compareHandsJoker(cards, cardMap)

	for i, card := range correctOrder{
		totalScore += (i + 1) * cards[card]
	}

	return totalScore
}

func compareHands(cards map[string]int, cardMap map[string]int) []string{
	fiveKinds := make([]string, 0)
	fourKinds := make([]string, 0)
	fullHouses := make([]string, 0)
	threeKinds := make([]string, 0)
	twoPairs := make([]string, 0)
	onePairs := make([]string, 0)
	highCards := make([]string, 0)

	for card, _ := range cards{
		count_ := make(map[rune]int)
		for _, c := range card{
			count_[c]++
		}

		if len(count_) == 1{
			fiveKinds = append(fiveKinds, card)
		}else if len(count_) == 2{
			for _, v := range count_{
				if v == 4 || v == 1{
					fourKinds = append(fourKinds, card)
					break
				}else if v == 3 || v == 2{
					fullHouses = append(fullHouses, card)
					break
				}
			}
		}else if len(count_) == 3{
			for _, v := range count_{
				if v == 3 {
					threeKinds = append(threeKinds, card)
					break
				}else if v == 2{
					twoPairs = append(twoPairs, card)
					break
				}
			}
		}else if len(count_) == 4{
			onePairs = append(onePairs, card)
		}else{
			highCards = append(highCards, card)
		}
	}
	fmt.Println(fiveKinds, fourKinds, fullHouses, threeKinds, twoPairs, onePairs, highCards)

	fiveKinds = sortHands(fiveKinds, cardMap)
	fourKinds = sortHands(fourKinds, cardMap)
	fullHouses = sortHands(fullHouses, cardMap)
	threeKinds = sortHands(threeKinds, cardMap)
	twoPairs = sortHands(twoPairs, cardMap)
	onePairs = sortHands(onePairs, cardMap)
	highCards = sortHands(highCards, cardMap)

	return append(append(append(append(append(append(highCards, onePairs...), twoPairs...), threeKinds...), fullHouses...), fourKinds...), fiveKinds...)

}

func sortHands(hands []string, cardMap map[string]int) []string{
	
	sort.Slice(hands, func(i, j int) bool {
		return compareCards(hands[i], hands[j], cardMap)
	})

	return hands
}
func compareCards(card1, card2 string, cardMap map[string]int) bool {
	for i := 0; i < len(card1) && i < len(card2); i++ {
		rank1, found1 := cardMap[string(card1[i])]
		rank2, found2 := cardMap[string(card2[i])]

		if found1 && found2 {
			if rank1 != rank2 {
				return rank1 < rank2
			}
		}
	}

	return len(card1) < len(card2)
}

// PART TWO

func compareHandsJoker(cards map[string]int, cardMap map[string]int) []string{
	fiveKinds := make([]string, 0)
	fourKinds := make([]string, 0)
	fullHouses := make([]string, 0)
	threeKinds := make([]string, 0)
	twoPairs := make([]string, 0)
	onePairs := make([]string, 0)
	highCards := make([]string, 0)

	for card, _ := range cards{
		count_ := make(map[rune]int)

		for _, c := range card{
			count_[c]++
		}

		joker, found := count_['J']
		
		if found && len(count_) > 1{
			val := upgradeCardWithJoker(card, joker);
			count_ = make(map[rune]int)
			for _, c := range val{
				count_[c]++
			}
		}
		
		if len(count_) == 1{
			fiveKinds = append(fiveKinds, card)
		}else if len(count_) == 2{
			for _, v := range count_{
				if v == 4 || v == 1{
					fourKinds = append(fourKinds, card)
					break
				}else if v == 3 || v == 2{
					fullHouses = append(fullHouses, card)
					break
				}
			}
		}else if len(count_) == 3{
			for _, v := range count_{
				if v == 3 {
					threeKinds = append(threeKinds, card)
					break
				}else if v == 2{
					twoPairs = append(twoPairs, card)
					break
				}
			}
		}else if len(count_) == 4{
			onePairs = append(onePairs, card)
		}else{
			highCards = append(highCards, card)
		}
		
		
	}
	fmt.Println(fiveKinds, fourKinds, fullHouses, threeKinds, twoPairs, onePairs, highCards)

	fiveKinds = sortHands(fiveKinds, cardMap)
	fourKinds = sortHands(fourKinds, cardMap)
	fullHouses = sortHands(fullHouses, cardMap)
	threeKinds = sortHands(threeKinds, cardMap)
	twoPairs = sortHands(twoPairs, cardMap)
	onePairs = sortHands(onePairs, cardMap)
	highCards = sortHands(highCards, cardMap)

	return append(append(append(append(append(append(highCards, onePairs...), twoPairs...), threeKinds...), fullHouses...), fourKinds...), fiveKinds...)

}

func upgradeCardWithJoker(card string, jokerCount int) string {
	count_ := make(map[rune]int)

	// Count the occurrences of each character in the card
	for _, c := range card {
		count_[c]++
	}

	// Replace 'J' characters with characters with the highest frequency
	for i := 0; i < jokerCount; i++ {
		// Find the character with the highest frequency
		maxChar := 'J' // Default to 'J' if no other characters are found
		maxFreq := 0

		for key, value := range count_ {
			if key != 'J' && value > maxFreq {
				maxChar = key
				maxFreq = value
			}
		}

		// Check if the sum of frequencies is not more than 5
		if count_['J']+maxFreq > 5 {
			// Distribute the remaining 'J' count to other characters
			for key := range count_ {
				if key != 'J' {
					remaining := 5 - count_[maxChar] // Remaining space for the chosen character
					count_[key] += count_['J']
					count_['J'] = 0

					if count_[key] > remaining {
						count_['J'] = count_[key] - remaining
						count_[key] = remaining
					}
					break
				}
			}
		}

		// Replace 'J' with the character and increase its strength
		count_[maxChar]++
		delete(count_, 'J')
	}

	// Reconstruct the upgraded card string
	var upgradedCard strings.Builder
	for key, value := range count_ {
		for i := 0; i < value; i++ {
			upgradedCard.WriteRune(key)
		}
	}

	return upgradedCard.String()
}
