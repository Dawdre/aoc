package main

import (
	"aoc/2022/lib"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Fetch AoC input
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5fcbd4cb1717a7e8ab1e0a8afa456c3977ee26da5b259adc6f6822fb4dbf72e9c1cefa99b615488e424d9c64a1a185ffce91e14787b1b24ec7",
	}

	// -------------------
	// Day One
	// -------------------

	// Read and split
	day_one_response, one_nil := lib.FetchAOCInput("1", cookie, client)
	calories := strings.Split(day_one_response, "\n\n")

	var total_calories []int

	// First we loop the string
	for _, value := range calories {
		var number_calories []int
		// Capture whitespaced string groups
		numbers := strings.Fields(value)

		// Convert to numbers
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != one_nil {
				fmt.Println("failed to convert string to int")
			}

			// Add converted numbers to array
			number_calories = append(number_calories, num)
		}

		// Sum and append
		total_calories = append(total_calories, []int{sum(number_calories)}...)
	}

	// Sort
	slices.Sort(total_calories)
	slices.Reverse(total_calories)

	// Answer Part One and Part Two
	fmt.Println("Day One Answer")
	fmt.Printf("Max calories: %d\n", total_calories[0])
	fmt.Printf("Top 3 Elves: %d\n\n", sum(total_calories[:3]))

	// -------------------
	// Day Two
	// -------------------
	day_two_response, _ := lib.FetchAOCInput("2", cookie, client)

	strategy_raw := strings.Split(day_two_response, "\n")
	strategy_raw = strategy_raw[:len(strategy_raw)-1]
	total_score_p1 := 0
	total_score_p2 := 0

	for _, val := range strategy_raw {
		strat := strings.Split(val, " ")

		total_score_p1 += CalcRoundScore(strat)
		total_score_p2 += CalcExpectedRoundScore(strat)
	}

	fmt.Println("Day Two Answer")
	fmt.Printf("Total Score P1: %d\n", total_score_p1)
	fmt.Printf("Total Score P2: %d\n\n", total_score_p2)

	DayThree()
}

func CalcExpectedRoundScore(strat []string) int {
	score := 0
	opponent, result := strat[0], strat[1]

	resultMap := map[string]int{
		"X": 0, // loss
		"Y": 3, // draw
		"Z": 6, // win
	}

	rulesMap := map[string][]int{
		// loss, draw, win
		"A": {3, 1, 2}, // rock
		"B": {1, 2, 3}, // paper
		"C": {2, 3, 1}, // scissors
	}

	// calc round score
	score += resultMap[result]

	// calc shape score
	if score == 6 {
		score += rulesMap[opponent][2]
	} else if score == 3 {
		score += rulesMap[opponent][1]
	} else {
		score += rulesMap[opponent][0]
	}

	return score
}

func CalcRoundScore(strat []string) int {
	score := 0

	rock, paper, scissors := []string{"A", "X"}, []string{"B", "Y"}, []string{"C", "Z"}
	rock_shape, paper_shape, scissors_shape := 1, 2, 3
	draw, win := 3, 6

	rules := [][]string{
		{rock[0], scissors[1]},  // loss A Z
		{paper[0], rock[1]},     // loss B X
		{scissors[0], paper[1]}, // loss C Y

		{rock[0], paper[1]},     // win A Y
		{paper[0], scissors[1]}, // win B Z
		{scissors[0], rock[1]},  // win C X

	}

	player := strat[1]

	// calc win or loss
	for index, rule := range rules {
		if slices.Equal(strat, rule) && index > 2 {
			score += win
		}
	}

	// calc draws
	if slices.Equal(strat, rock) || slices.Equal(strat, paper) || slices.Equal(strat, scissors) {
		score += draw
	}

	// shape scores
	if player == rock[1] {
		score += rock_shape
	} else if player == paper[1] {
		score += paper_shape
	} else if player == scissors[1] {
		score += scissors_shape
	}

	return score
}

// func FetchAOCInput(url string, cookie *http.Cookie, client *http.Client) (string, error) {
// 	url_day := fmt.Sprintf("https://adventofcode.com/2022/day/%s/input", url)

// 	request, err := http.NewRequest("GET", url_day, nil)

// 	if err != nil {
// 		return "", fmt.Errorf("failed to create request: %w", err)
// 	}

// 	request.AddCookie(cookie)

// 	response, err := client.Do(request)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to make GET request: %w", err)
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("error code: %d", response.StatusCode)
// 	}

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("can't read response body: %w", err)
// 	}

// 	return string(body), nil
// }

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
