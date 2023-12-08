package main

import (
	"os"
	"bufio"
	"log"
	"regexp"
	"strconv"
	"slices"
	"cmp"
)

const(
	FIVE_OF_A_KIND = 7
	FOUR_OF_A_KIND = 6
	FULL_HOUSE = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR = 3
	ONE_PAIR = 2
	HIGH_CARD = 1
)

type Card struct{
	name rune
	value int
}

type Hand struct{
	cards [5]Card
	handType int
	bid int
}

/* 
 * Helper to error check
 */
func check(e error){
	if e != nil {
		panic(e)
	}
}

/**
 * Custom compare for hands
 * Compare hand type, if the same, compare card by card
 */
 func compareHands(i Hand, j Hand) int{
	if i.handType == j.handType{
		for ind, card := range i.cards{
			if card.value > j.cards[ind].value{
				return 1
			}
			if card.value < j.cards[ind].value{
				return -1
			}
		}
	}
	return cmp.Compare(i.handType, j.handType)
}

/**
 * Helper function to return card value
 */
func getCardValue(cardName rune) int {
	switch cardName{
		case 'A':
			return 14
		case 'K':
			return 13
		case 'Q':
			return 12
		case 'J':
			return 11
		case 'T':
			return 10
		case '9':
			return 9
		case '8':
			return 8
		case '7':
			return 7
		case '6':
			return 6
		case '5':
			return 5
		case '4':
			return 4
		case '3':
			return 3
		case '2':
			return 2
	}
	panic("Undefined card found.")
}

/**
 * Get Hand type
 */
func getHandType(hand Hand) int {
	dict := countOccurencesRunes(hand)
	handType := HIGH_CARD
	for _, item := range dict{
		if(item == 5){
			handType = FIVE_OF_A_KIND
			break
		}
		if(item == 4){
			handType = FOUR_OF_A_KIND
			break
		}
		if(item == 3){
			// if pair was already found
			if (handType == ONE_PAIR){
				handType = FULL_HOUSE
				break
			}
			handType = THREE_OF_A_KIND
		}
		if(item == 2){
			// if three of a kind was already found
			if (handType == THREE_OF_A_KIND){
				handType = FULL_HOUSE
				break
			}
			// if pair was already found
			if (handType == ONE_PAIR){
				handType = TWO_PAIR
				break
			}
			handType = ONE_PAIR
		}
	}
	return handType
}

/**
 * Helper function to count occurences
 */
func countOccurencesRunes(hand Hand) map[rune]int {
	dict := make(map[rune]int, 0)
	for _, card := range hand.cards{
		dict[card.name] +=1
	}
	return dict
}

/**
 * Get Hand and Bid from line
 */
 func getHandAndBid(line string) Hand{
	regex := regexp.MustCompile(`(\w+) (\d+)`)
	matches := regex.FindAllStringSubmatch(line, -1)
	var hand Hand
	for ind, rune := range matches[0][1]{
		card := Card{name: rune, value: getCardValue(rune) }
		hand.cards[ind] = card
	}
	hand.bid, _ = strconv.Atoi(matches[0][2])
	hand.handType = getHandType(hand)
	return hand
}


/**
 * Day 7: Part 1 - Camel Cards
 * hand  bid
 * 32T3K 765
 * T55J5 684
 * KK677 28
 * KTJJT 220
 * QQQJA 483
 */
func main(){
// Read input file
	input_file, err := os.Open("day7_input.txt")
	check(err)
	defer input_file.Close()
	line := ""
	partOne := 0
	var hands []Hand
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan(){
		line = scanner.Text()
		hand := getHandAndBid(line)
		hands = append(hands, hand)
	}
	// Sort hands by strength/rank
	slices.SortFunc(hands, compareHands)
	// Sum rank * bid for all hands
	for ind, hand := range hands{
		partOne += hand.bid * (ind+1)
	}
	log.Println("Part 1. ", partOne)
}