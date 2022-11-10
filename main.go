package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", " a csv file in the format of question, answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// need to keep track of how many we get correct
	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// make a go routine that sends a message when we get an answer from the user
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		// it's always waiting for one case or the other
		select {
		case <-timer.C:
			fmt.Println()
			// breaks out of prob loop so you don't have to replicate the line
			break problemloop
			// return breaks out of the for loop, break would break out of the select
			//return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
				fmt.Println("Correct 🥰")
			} else {
				fmt.Println("🤡 haha loser")
			}
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
