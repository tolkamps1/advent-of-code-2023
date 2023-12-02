package main

import (
	"os"
	"bufio"
	"log"
	"regexp"
	"strconv"
)

/* 
 *Helper to error check
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


func main(){
// Read input file
	input_file, err := os.Open("day2_input.txt")
	check(err)
	defer input_file.Close()

	sumPart1 := 0
	line := ""
	colours := [3]string{"blue", "red", "green"}
	maxCubeInput := make(map[string]int)
	maxCubeInput["blue"] = 14
	maxCubeInput["red"] = 12
	maxCubeInput["green"] = 13

	scanner := bufio.NewScanner(input_file)
	for scanner.Scan(){
		line = scanner.Text()
		gameId := getGame(line)
		isGamePossible := true
		for _, colour := range colours{
			if getMaxOfColour(line, colour) > maxCubeInput[colour]{
				isGamePossible = false
				break
			}
		}
		if isGamePossible{
			sumPart1 += gameId
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
        return	
	}
	log.Println("Part 1. Sum of games:", sumPart1)
	log.Println("Done.")
}