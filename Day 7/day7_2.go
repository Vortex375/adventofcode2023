package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	PICTURE_STRENGTH_HIGH_CARD = iota
	PICTURE_STRENGTH_ONE_PAIR
	PICTURE_STRENGTH_TWO_PAIR
	PICTURE_STRENGTH_THREE_OF_A_KIND
	PICTURE_STRENGTH_FULL_HOUSE
	PICTURE_STRENGTH_FOUR_OF_A_KIND
	PICTURE_STRENGTH_FIVE_OF_A_KIND
)

var cardRunes = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

type HandAndBid struct {
	hand Hand
	bid  int
}

type Hand struct {
	hand       string
	runes      []rune
	cardCounts map[rune]int
}

func NewHand(hand string) Hand {
	cardCounts := make(map[rune]int)
	runes := make([]rune, len(hand))
	for i, r := range hand {
		runes[i] = r
		if r != 'J' { /* do not count jokers as regular cards */
			cardCounts[r] += 1
		}
	}

	return Hand{hand, runes, cardCounts}
}

func (h *Hand) countCount(count int) int {
	countCount := 0
	for _, v := range h.cardCounts {
		if v == count {
			countCount++
		}
	}
	return countCount
}

func (h *Hand) jokerCount() int {
	count := 0
	for _, r := range h.hand {
		if r == 'J' {
			count++
		}
	}
	return count
}

func (h *Hand) pictureStrength() int {
	jokerCount := h.jokerCount()

	switch {
	case jokerCount == 5:
		return PICTURE_STRENGTH_FIVE_OF_A_KIND
	case jokerCount == 4:
		return PICTURE_STRENGTH_FIVE_OF_A_KIND
	case jokerCount == 3:
		if h.countCount(2) > 0 {
			return PICTURE_STRENGTH_FIVE_OF_A_KIND
		}
		return PICTURE_STRENGTH_FOUR_OF_A_KIND
	case jokerCount == 2:
		if h.countCount(3) > 0 {
			return PICTURE_STRENGTH_FIVE_OF_A_KIND
		}
		if h.countCount(2) > 0 {
			return PICTURE_STRENGTH_FOUR_OF_A_KIND
		}
		return PICTURE_STRENGTH_THREE_OF_A_KIND
	case jokerCount == 1:
		if h.countCount(4) > 0 {
			return PICTURE_STRENGTH_FIVE_OF_A_KIND
		}
		if h.countCount(3) > 0 {
			return PICTURE_STRENGTH_FOUR_OF_A_KIND
		}
		if h.countCount(2) == 2 {
			return PICTURE_STRENGTH_FULL_HOUSE
		}
		if h.countCount(2) == 1 {
			return PICTURE_STRENGTH_THREE_OF_A_KIND
		}
		return PICTURE_STRENGTH_ONE_PAIR

	/* no jokers */
	case h.countCount(5) == 1:
		return PICTURE_STRENGTH_FIVE_OF_A_KIND
	case h.countCount(4) == 1:
		return PICTURE_STRENGTH_FOUR_OF_A_KIND
	case h.countCount(3) == 1 && h.countCount(2) == 1:
		return PICTURE_STRENGTH_FULL_HOUSE
	case h.countCount(3) == 1:
		return PICTURE_STRENGTH_THREE_OF_A_KIND
	case h.countCount(2) == 2:
		return PICTURE_STRENGTH_TWO_PAIR
	case h.countCount(2) == 1:
		return PICTURE_STRENGTH_ONE_PAIR
	default:
		return PICTURE_STRENGTH_HIGH_CARD
	}
}

func (h *Hand) less(o *Hand) bool {
	if h.pictureStrength() < o.pictureStrength() {
		// fmt.Printf("Less: %s %s %t\n", h.hand, o.hand, true)
		return true
	}
	if h.pictureStrength() == o.pictureStrength() {
		for i := range h.runes {
			thisRune := h.runes[i]
			otherRune := o.runes[i]
			if runeStrength(thisRune) > runeStrength(otherRune) {
				return false
			}
			if runeStrength(thisRune) < runeStrength(otherRune) {
				// fmt.Printf("RuneStrength: %s %d %s %d", string(thisRune), runeStrength(thisRune), string(otherRune), runeStrength(otherRune))
				// fmt.Printf(" Less: %s %s %t\n", h.hand, o.hand, true)
				return true
			}
		}
	}
	// fmt.Printf("Less: %s %s %t\n", h.hand, o.hand, false)
	return false
}

func runeStrength(r rune) int {
	for i, c := range cardRunes {
		if r == c {
			return i
		}
	}
	panic("Not a card rune: " + string(r))
}

func main() {
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	handsAndBids := make([]HandAndBid, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		var hand string
		var bid int
		fmt.Sscanf(line, "%s %d", &hand, &bid)
		handsAndBids = append(handsAndBids, HandAndBid{NewHand(hand), bid})
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Slice(handsAndBids, func(i, j int) bool {
		return handsAndBids[i].hand.less(&handsAndBids[j].hand)
	})

	for _, h := range handsAndBids {
		fmt.Printf("%s %d\n", h.hand.hand, h.hand.pictureStrength())
	}

	result := 0
	for i, handAndBid := range handsAndBids {
		result += (i + 1) * handAndBid.bid
	}

	fmt.Println(result)
}
