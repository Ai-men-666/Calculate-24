package main

import (
	"fmt"
	"log"
)

var (
	card = []string{"", "", "", ""}
	used = []bool{false, false, false, false}
	path = []float64{0, 0, 0}
	num  = []float64{0, 0, 0, 0}
)

func ParseCard() {
	for i := 0; i < 4; i++ {
		if card[i] == "A" {
			num[i] = 1
		} else if card[i] == "J" {
			num[i] = 11
		} else if card[i] == "Q" {
			num[i] = 12
		} else if card[i] == "K" {
			num[i] = 13
		} else if card[i] == "10" {
			num[i] = 10
		} else if len(card[i]) == 1 && card[i][0] >= '2' && card[i][0] <= '9' {
			num[i] = float64(card[i][0] - '0')
		} else {
			log.Fatalln("Invalid card:", card[i])
		}
	}
}

func dfs() {
	if len(path) == 4 {
		calculate()
		return
	}
	for i := 0; i < 4; i++ {
		if !used[i] {
			used[i] = true
			path = append(path, num[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func calculate() {
	//TODO
}

func tryParentheses(a float64, b float64, c float64, d float64) string {
	//TODO
	return ""
}

func main() {
	for i := 0; i < 4; i++ {
		_, err := fmt.Scan(&card[i])
		if err != nil {
			log.Fatalln("Failed to read input:", err)
		}
	}
	log.Println("Input cards:", card)
	ParseCard()
	log.Println("Parsed numbers:", num)
	dfs()
	log.Println("No Answer")
	return
}
