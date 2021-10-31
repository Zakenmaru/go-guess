package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter the CSV of math questions(make sure it's in the same folder as the main file!): ")
	var file string

	fmt.Scanln(&file)

	f, err := os.Open(file)

	if err != nil {
		error_msg(err)
	}

	eqMap := makeQuestionMap(f)
	total, correct := quizGame(eqMap)
	// printMap(eqMap)
	fmt.Println("Total Questions: ", total)
	fmt.Println("Correct: ", correct)
}

func quizGame(eqMap map[string]int) (int, int) {
	total := 0
	correct := 0
	for eq, res := range eqMap {
		fmt.Println("Enter the answer: " + eq)

		var ans string
		fmt.Scanln(&ans)

		intAns, err := strconv.Atoi(ans)
		if err != nil {
			error_msg(err)
		}
		if intAns == res {
			correct++
		}
		total++

	}
	return total, correct
}

func printMap(eqMap map[string]int) {
	for eq, res := range eqMap {
		fmt.Println("Equation: ", eq, "\tResult: ", res)
	}
}

func error_msg(err error) {
	fmt.Errorf("error: ", err)
	os.Exit(1)
}

func makeQuestionMap(f *os.File) map[string]int {
	eqMap := make(map[string]int)

	r := csv.NewReader(f)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			error_msg(err)
		}
		// parse answer
		a, err := strconv.Atoi(line[1])
		if err != nil {
			error_msg(err)
		}
		eqMap[line[0]] = a
	}

	return eqMap
}
