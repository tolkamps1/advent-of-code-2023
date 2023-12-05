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
 * Get Mapping
 */
 func getMapping(line string) (int, int, int) {
	regex := regexp.MustCompile(`\d+`)
	allNumbers := regex.FindAllString(line, -1)
	dest, e := strconv.Atoi(allNumbers[0])
	check(e)
	src, e := strconv.Atoi(allNumbers[1])
	check(e)
	r, e := strconv.Atoi(allNumbers[2])
	check(e)
	return dest, src, r
}

/**
 * Get Seeds from line Part 1 implementation
 */
 func getSeedsPart1(line string) []int{
	regex := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(line, -1)
	var seeds []int
	for _, match := range matches{
		seed, e := strconv.Atoi(match)
		check(e)
		seeds = append(seeds, seed)
	}
	return seeds
}

/**
 * Get Seeds from line Part 2 implementation
 * Read pair (start range), add seeds from start to start + range -1
 */
 func getSeeds(line string) []int{
	regex := regexp.MustCompile(`(\d+) (\d+)`)
	matches := regex.FindAllStringSubmatch(line, -1)
	var seeds []int
	for _, match := range matches{
		limit, e := strconv.Atoi(match[2])
		check(e)
		for num := 0; num < limit; num++{
			seed, e := strconv.Atoi(match[1])
			check(e)
			seeds = append(seeds, seed + num)
		}
	}
	log.Println(seeds)
	return seeds
}

/**
 * Day 5: Part 1 - If You Give A Seed A Fertilizer
 * Map seeds to location and get lowest location number. 
 * Map is Dest Src Range, aka (src to src+range-1 = dest to dest+range-1) All others are 1-1 
 */
func main(){
// Read input file
	input_file, err := os.Open("day5_input.txt")
	check(err)
	defer input_file.Close()
	line := ""
	getNextMap := false
	scanner := bufio.NewScanner(input_file)
	// Read first line to get seeds
	scanner.Scan()
	line = scanner.Text()
	currentSource := getSeeds(line)
	currentDestination := make([]int, len(currentSource))
	copy(currentDestination, currentSource)
	for scanner.Scan(){
		line = scanner.Text()
		if getNextMap{
			// start next mapping
			getNextMap = false
			continue
		}
		if line == ""{
			// finished mapping, convert remaining 1-1
			copy(currentSource, currentDestination)
			getNextMap = true
			continue
		}
		// map 
		dest, src, r := getMapping(line)
		for ind, curr_src := range currentSource{
			// if mapping applies to this current item
			upper := src + r -1
			if src <= curr_src && curr_src <= upper{
				currentDestination[ind] = curr_src + (dest-src)
			}
		}
	}

	// Find smallest location number
	closestLocation := currentDestination[0]
	for _, location := range currentDestination{
		if(location < closestLocation){
			closestLocation = location
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
        return
	}

	log.Println("Part 1. Closest Location ", closestLocation)
}