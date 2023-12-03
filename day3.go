package main

import (
	"os"
	"bufio"
	"log"
	"regexp"
	"strconv"
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
 * Get all numbers and their indexes in line
 */
func getNumbers(line string, matches map[string][]int){
	regex := regexp.MustCompile(`\d+`)
	allStrings := regex.FindAllString(line, -1)
	allStringIndexes := regex.FindAllStringIndex(line, -1)
	for ind, str := range allStrings{
		// Add to dict with ind to handle duplicates on lines
		matches[str + "_" + strconv.Itoa(ind)] = allStringIndexes[ind]
	}
}


/**
 * Check overlap
 */
func checkSpecCharIndex(allStringIndex[][]int, numIndex[]int) bool{
	// for each special character index, check if matches number index
	numberFirstIndex := numIndex[0] - 1
	numberLastIndex := numIndex[1]
	for _, specCharIndex := range allStringIndex{
		charIndex := specCharIndex[0]
		if numberFirstIndex <= charIndex && charIndex <= numberLastIndex{
			return true
		}
	}
	return false
}


/**
 * Check for collision of number index (+-1) with special character
 */
func checkForCollision(matches map[string][]int, line string, lineBefore string, lineAfter string) int {
	sum := 0
	// regex to match all non-digit and non-period characters
	regex := regexp.MustCompile(`([^.\d]+)`)
	// regex to get num from dict
	regexNum := regexp.MustCompile(`(\d+)_`)
	beforeAllStringIndexes := regex.FindAllStringIndex(lineBefore, -1)
	lineAllStringIndexes := regex.FindAllStringIndex(line, -1)
	afterAllStringIndexes := regex.FindAllStringIndex(lineAfter, -1)
	// for each number
	for number, indexes := range matches{
		// for each special character index, check if matches number index
		numFromMatch := regexNum.FindAllStringSubmatch(number,-1)[0][1]
		num, e := strconv.Atoi(numFromMatch)
		check(e)
		if checkSpecCharIndex(beforeAllStringIndexes, indexes){
			sum += num
			continue
		}
		if checkSpecCharIndex(lineAllStringIndexes, indexes){
			sum += num
			continue
		}
		if checkSpecCharIndex(afterAllStringIndexes, indexes){
			sum += num
			continue
		}
	}
	return sum
}


/**
 * Day 3: Part 1
 * Sum each part (aka number) that is adjacent or immediately diagonal of a special character
 * 
 */
func main(){
// Read input file
	input_file, err := os.Open("day3_input.txt")
	check(err)
	defer input_file.Close()
	sumPart1, sumPart2 := 0, 0
	lineBefore, line, lineAfter := "", "", ""

	scanner := bufio.NewScanner(input_file)
	scanner.Scan() // first line
	line = scanner.Text()
	for scanner.Scan(){
		lineAfter = scanner.Text()
		matches := make(map[string][]int)
		getNumbers(line, matches)
		sumPart1 += checkForCollision(matches, line, lineBefore, lineAfter)
		// increment lines
		lineBefore = line
		line = lineAfter
	}
	// Handle last line
	lineAfter = ""
	matches := make(map[string][]int)
	getNumbers(line, matches)
	sumPart1 += checkForCollision(matches, line, lineBefore, lineAfter)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
        return	
	}
	log.Println("Part 1. ", sumPart1)
	log.Println("Part 2. ", sumPart2)
}