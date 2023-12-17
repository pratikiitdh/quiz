package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file that contains question,answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read csv file")
	}
	problems := parseLines(lines)

	correct := 0
	for index, problem := range problems {
		fmt.Printf("Problem #%d, %s = \n", index+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.a {
			correct++
		}
	}
	fmt.Printf("You score %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for index, line := range lines {
		ret[index] = problem{q: line[0], a: strings.TrimSpace(line[1])}
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
