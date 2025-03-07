package main

import (
	"aoc/2022/lib"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func DayThree() {
	// Fetch AoC input
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5fcbd4cb1717a7e8ab1e0a8afa456c3977ee26da5b259adc6f6822fb4dbf72e9c1cefa99b615488e424d9c64a1a185ffce91e14787b1b24ec7",
	}

	// -------------------
	// Day Three - Part 1
	// -------------------

	day_three_response, err := lib.FetchAOCInput("3", cookie, client)
	if err != nil {
		fmt.Printf("Failed to get input - %s", err)
	}

	rucksacks := strings.Split(day_three_response, "\n")
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabet_upper := strings.ToUpper(alphabet)

	prio_score := 0

	for _, sack := range rucksacks {
		sack_length := len(sack) / 2
		comp_one := sack[:sack_length]
		comp_two := sack[sack_length:]

		for _, huh := range comp_one {
			if strings.ContainsRune(comp_two, huh) {
				letter := string(huh)

				if strings.Contains(alphabet, letter) {
					prio_score += strings.Index(alphabet, letter) + 1

				} else {
					prio_score += strings.Index(alphabet_upper, letter) + 27
				}

				break
			}
		}
	}

	fmt.Println("Day Three Answer")
	fmt.Printf("Priority Total Rucksacks: %d\n", prio_score)

	// -------------------
	// Day Three - Part 2
	// -------------------

	prio_score_two := 0

	for i := 0; i < len(rucksacks); i += 3 {
		end := i + 3

		if end > len(rucksacks) {
			end = len(rucksacks)
		}

		sack_group := rucksacks[i:end]
		// create first and primary map
		common := make(map[rune]bool)

		for _, sack := range sack_group[0] {
			// mark the first letters in the first string
			common[sack] = true
		}

		for _, sack := range sack_group[1:] {
			// create second and comparison map
			current := make(map[rune]bool)

			for _, item := range sack {
				current[item] = true
			}

			for item := range common {
				// if each common letter does not exist in the current string
				// delete letter from the common map
				if !current[item] {
					delete(common, item)
				}
			}
		}

		for item := range common {
			prio_score_two += AlphabetScore(string(item))
		}
	}

	fmt.Printf("Priority Total Elf Group: %d\n", prio_score_two)
}

func AlphabetScore(letter string) int {
	score := 0
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabet_upper := strings.ToUpper(alphabet)

	if strings.Contains(alphabet, letter) {
		score += strings.Index(alphabet, letter) + 1
	} else {
		score += strings.Index(alphabet_upper, letter) + 27
	}

	return score
}
