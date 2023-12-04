package main

import (
	"os"
	"bufio"
	"log"
	"regexp"
	"strings"
	"strconv"
	"math"
)

/* 
 * Helper to error check
 */
func check(e error){
	if e != nil {
		panic(e)
	}
}

/**
 * Get winning numbers
 */
 func getWinningNumbers(line string) []string {
	regex := regexp.MustCompile(`(\d+) `)
	allNumbers := regex.FindAllStringSubmatch(line, -1)
	var winNum []string
	for _, num := range allNumbers{
		winNum = append(winNum, num[1])
	}
	return winNum
}

/**
 * Check and count number of winners in line
 */
 func countWinners(winningNums []string, nums string) int {
	count := 0
	for _, reg := range winningNums{
		regex := regexp.MustCompile(" "+reg+" ")
		allMatch := regex.FindAllString(nums, -1)
		count += len(allMatch)
	}
	return count
}

/**
 * Get Card Number from line
 */
 func getCardNumber(line string) int{
	regex := regexp.MustCompile(`(\d+):`)
	match := regex.FindStringSubmatch(line)
	id := match[1]
	cardId, e := strconv.Atoi(id)
	check(e)
	return cardId
}

/**
 * Day 4: Part 1 - Scratch cards. Counts number of winning numbers (before '|') that occur after '|'
 * Part 2 - Scratch card sums. Based off number of wins, copies of next scratch cards are added. Calculates total number of cards. 
 */
func main(){
// Read input file
	input_file, err := os.Open("day4_input.txt")
	check(err)
	defer input_file.Close()
	sumPart1, sumPart2 := 0, 0
	line := ""
	//Could pre-read file to get line count, but input is static in this case
	numberOfCards := [197]int{}
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan(){
		line = scanner.Text()
		cardNumber := getCardNumber(line)
		// Start with one card of each number
		numberOfCards[cardNumber] += 1
		lineSplit := strings.Split(line, "|")
		winningNumbers := getWinningNumbers(lineSplit[0])
		winningCount := countWinners(winningNumbers, lineSplit[1]+" ")
		if winningCount > 0{
			sumPart1 += int(math.Pow(2, float64(winningCount-1)))
			// increase card size for each
			for count := 1; count <= winningCount; count++{
				numberOfCards[cardNumber+count] += numberOfCards[cardNumber]
			}
		}
	}
	// Sum total card count for Part 2
	for _, count := range numberOfCards{
		sumPart2 += count
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
        return
	}

	log.Println("Part 1. ", sumPart1)
	log.Println("Part 2. ", sumPart2)
}