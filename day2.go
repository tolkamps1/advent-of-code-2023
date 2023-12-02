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
 * Get game ID
 */
func getGame(line string) int{
	regex := regexp.MustCompile(`(\d+):`)
	match := regex.FindStringSubmatch(line)
	id := match[1]
	gameId, e := strconv.Atoi(id)
	check(e)
	return gameId
}

/**
 * Get max cubes of a colour
 */
 func getMaxOfColour(line string, colour string) int{
	regex := regexp.MustCompile(` (\d+) `+colour)
	match := regex.FindAllStringSubmatch(line, -1)
	maxRow := len(match)
	maxColour, e := strconv.Atoi(match[0][1])
	check(e)
	for i, j := 0, 1; i < maxRow; i++{
		curr, e := strconv.Atoi(match[i][j])
		check(e)
		if(curr > maxColour){
			maxColour = curr
		}
	}
	return maxColour
}

/**
 * Day 2: Part 1 and 2
 * Sample Input: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
 * 				 Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
 * 				 Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
 * 				 Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
 * 				 Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
 */
func main(){
// Read input file
	input_file, err := os.Open("day2_input.txt")
	check(err)
	defer input_file.Close()
	sumPart1, sumPart2 := 0, 0
	line := ""
	colours := [3]string{"blue", "red", "green"}
	// Set up input for Part 1
	maxCubeInput := make(map[string]int)
	maxCubeInput["blue"] = 14
	maxCubeInput["red"] = 12
	maxCubeInput["green"] = 13

	scanner := bufio.NewScanner(input_file)
	for scanner.Scan(){
		line = scanner.Text()
		gameId := getGame(line)
		power := 1
		isGamePossible := true
		// Get max occurrence of each colour in a game
		for _, colour := range colours{
			maxOfColour := getMaxOfColour(line, colour)
			power *= maxOfColour
			if maxOfColour > maxCubeInput[colour]{
				isGamePossible = false
			}
		}
		// If game is possible, add ID to sum for Part 1
		if isGamePossible{
			sumPart1 += gameId
		}
		sumPart2 += power
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
        return	
	}
	log.Println("Part 1. Sum of games:", sumPart1)
	log.Println("Part 2. Sum of powers:", sumPart2)
}