package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	card   = []string{"", "", "", ""}
	used   = []bool{false, false, false, false}
	path   = []float64{}
	num    = []float64{0, 0, 0, 0}
	arg    = []rune{'+', '-', '*', '/'}
	result = ""
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
	if result != "" {
		return
	}
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
	op := []rune{0, 0, 0}
	for i := 0; i < 4; i++ {
		op[0] = arg[i]
		for j := 0; j < 4; j++ {
			op[1] = arg[j]
			for k := 0; k < 4; k++ {
				op[2] = arg[k]
				result = tryParentheses(path, op)
				if result != "" {
					fmt.Println(result)
					return
				}
			}
		}
	}
}

func operate(a, b float64, oper rune, res *float64) bool {
	if oper == '+' {
		*res = a + b
	} else if oper == '-' {
		*res = a - b
	} else if oper == '*' {
		*res = a * b
	} else {
		if math.Abs(b) < 1e-5 {
			return false
		}
		*res = a / b
	}
	return true
}
func tryParentheses(path []float64, op []rune) string {
	var t1, t2, t3 float64
	res := ""
	// ((path[0] op[0] path[1]) op[1] path[2]) op[2] path[3]
	if operate(path[0], path[1], op[0], &t1) && operate(t1, path[2], op[1], &t2) && operate(t2, path[3], op[2], &t3) && math.Abs(t3-24) < 1e-5 {
		return fmt.Sprint("((", strconv.Itoa(int(path[0])), string(op[0]), strconv.Itoa(int(path[1])), ")", string(op[1]), strconv.Itoa(int(path[2])), ")", string(op[2]), strconv.Itoa(int(path[3])))
	}
	// (path[0] op[0] (path[1] op[1] path[2])) op[2] path[3]
	if operate(path[1], path[2], op[1], &t1) && operate(path[0], t1, op[0], &t2) && operate(t2, path[3], op[2], &t3) && math.Abs(t3-24) < 1e-5 {
		return fmt.Sprint("(", strconv.Itoa(int(path[0])), string(op[0]), "(", strconv.Itoa(int(path[1])), string(op[1]), strconv.Itoa(int(path[2])), "))", string(op[2]), strconv.Itoa(int(path[3])))
	}
	// path[0] op[0] ((path[1] op[1] path[2]) op[2] path[3])
	if operate(path[1], path[2], op[1], &t1) && operate(t1, path[3], op[2], &t2) && operate(path[0], t2, op[0], &t3) && math.Abs(t3-24) < 1e-5 {
		return fmt.Sprint(strconv.Itoa(int(path[0])), string(op[0]), "((", strconv.Itoa(int(path[1])), string(op[1]), strconv.Itoa(int(path[2])), ")", string(op[2]), strconv.Itoa(int(path[3])), ")")
	}
	// path[0] op[0] (path[1] op[1] (path[2] op[2] path[3]))
	if operate(path[2], path[3], op[2], &t1) && operate(path[1], t1, op[1], &t2) && operate(path[0], t2, op[0], &t3) && math.Abs(t3-24) < 1e-5 {
		return fmt.Sprint(strconv.Itoa(int(path[0])), string(op[0]), "(", strconv.Itoa(int(path[1])), string(op[1]), "(", strconv.Itoa(int(path[2])), string(op[2]), strconv.Itoa(int(path[3])), "))")
	}
	// (path[0] op[0] path[1]) op[1] (path[2] op[2] path[3])
	if operate(path[0], path[1], op[0], &t1) && operate(path[2], path[3], op[2], &t2) && operate(t1, t2, op[1], &t3) && math.Abs(t3-24) < 1e-5 {
		return fmt.Sprint("(", strconv.Itoa(int(path[0])), string(op[0]), strconv.Itoa(int(path[1])), ")", string(op[1]), "(", strconv.Itoa(int(path[2])), string(op[2]), strconv.Itoa(int(path[3])), ")")
	}
	return res
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			http.ServeFile(w, r, "./static/main.css")
			return
		}
		http.ServeFile(w, r, "./static/main.html")
	} else if r.Method == "POST" {
		result = ""
		path = []float64{}
		used = []bool{false, false, false, false}

		r.ParseForm()
		card[0] = r.Form.Get("card1")
		card[1] = r.Form.Get("card2")
		card[2] = r.Form.Get("card3")
		card[3] = r.Form.Get("card4")
		log.Println("Input cards:", card)
		ParseCard()
		log.Println("Parsed numbers:", num)
		dfs()
		tmpl, err := template.ParseFiles("./static/res.html")
		if err != nil {
			log.Fatalln(err)
		}
		dataToShow := "No Answer"
		if result != "" {
			dataToShow = result
		}
		err = tmpl.Execute(w, dataToShow)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
	return
}
