package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", " a csv file in the format of question, answer")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse provided csv file 😱")
	}

	problems := parseLines(lines)

	// need to keep track of how many we get correct
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// read in an answer
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("Correct 🥰")
		} else {
			fmt.Println("🤡 haha loser")
		}
	}

	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	// every row in csv is a problem
	// if you know exactly how long the slice is, define it
	// append has to resize the slice everytime
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
