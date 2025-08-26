package main

import "log"

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

func main() {

}
