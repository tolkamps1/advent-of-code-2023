package main

import (
	"os"
	"bufio"
	"log"
	"regexp"
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
 * Get Time info of race
 */
 func getTimeInfo(line string, raceInfo [][]int) [][]int{
	regex := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(line, -1)
	races := []int{0,0}
	for _, match := range matches{
		race, e := strconv.Atoi(match)
		check(e)
		races[0] = race
		copySlice := make([]int, len(races))
		copy(copySlice, races)
		raceInfo = append(raceInfo, copySlice)
	}
	return raceInfo
}

/**
 * Get Distance record info of race
 */
 func getDistanceRecord(line string, raceInfo [][]int) [][]int{
	regex := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(line, -1)
	for ind, match := range matches{
		distance, e := strconv.Atoi(match)
		check(e)
		raceInfo[ind][1] = distance
	}
	return raceInfo
}

/**
 * Get Race Info Part 2 (read top line as single ms, and bottom as single mm)
 */
 func getRaceInfoP2(line string) int{
	regex := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(line, -1)
	numAsStr := ""
	for _, unit := range matches{
		numAsStr = numAsStr + unit
	}
	num, e := strconv.Atoi(numAsStr)
	check(e)
	return num
}

/**
 * Calculate the number of ways that this boat could win. This is the difference between the maximum and minimum 
 * win values (+1) which are calculated at how long the button must be held to increase the speed.
 * Formula calculation:
 *	Need to travel more than recordDistance during raceTime
 *	Minimum distance would then be recordDistance + 1
 *	raceTime = holdButtonTime + travelTime                     -> hbTime = raceTime - travelTime
 *	travelDistance = travelTime * holdButtonTime (aka speed)   -> hbTime = travelDistance/travelTime
 *	raceTime - travelTime = travelDistance/travelTime
 *	distance <= raceTime*t - t^2
 *	Algebraic equation: t^2 - (raceTime)t + (recordDistance+1) = 0
 *		aka (a=1;b=(-raceTime);c=(recordDistance+1))
 *		(-b±√(b²-4ac))/(2a) .
 *		Get min and max by + and -
 */
func calculateWinWays(raceTime int, recordDistance int) int {
	min := (float64(raceTime) - (math.Sqrt(float64((raceTime*raceTime) - 4 * (recordDistance +1 ) )))) / 2.0
	max :=  (float64(raceTime) + (math.Sqrt(float64((raceTime*raceTime) - 4 * (recordDistance +1 ) )))) / 2.0
	return int(math.Floor(max)) - int(math.Ceil(min)) + 1
}


/**
 * Day 6: Part 1 - 
 * Time is race time in ms and distance is the current record distance in mm. 
 * Each mm the boat's button is held down increases (from zero) the mm/ms the boat travels.
 * Determine the number of ways the record can be beat in each race; in this example, if you multiply these values together, you get 288 (4 * 8 * 9).
 * Part 2 - One race (ignore spaces between columns of time and distance)
 */
func main(){
// Read input file
	input_file, err := os.Open("day6_input.txt")
	check(err)
	defer input_file.Close()
	line := ""
	scanner := bufio.NewScanner(input_file)
	raceInfo := make([][]int, 0)
	scanner.Scan()
	// Read first line to get times
	line = scanner.Text()
	raceInfo = getTimeInfo(line, raceInfo)
	raceTimeP2 := getRaceInfoP2(line)
	scanner.Scan()
	// Read second line to get distances
	line = scanner.Text()
	getDistanceRecord(line, raceInfo)
	raceDistanceP2 := getRaceInfoP2(line)
	// Calculate ways to win for P1
	winWays := 0 
	for _, race := range raceInfo{
		if winWays == 0 {
			winWays = calculateWinWays(race[0], race[1])
		} else{
			winWays *= calculateWinWays(race[0], race[1])
		}
	}
	// Calculate win ways for P2
	winWaysP2 := calculateWinWays(raceTimeP2, raceDistanceP2)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
        return
	}

	log.Println("Part 1. Ways to win for each race multiplied together: ", winWays)
	log.Println("Part 2. Ways to win for one master race: ", winWaysP2)

}