package main

import (
	"flag"
	"strings"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
    // Help flag
	csvQuiz := flag.String("csv", "problems.csv", "csv file containg questions and answers in format `[question],[answer]`")
	flag.Parse()
	_ = csvQuiz

	// *csvQuiz - pointer to a string 
	file, err := os.Open(*csvQuiz)

	if err != nil {
		exit( fmt.Sprintf("Failed to open csv file %s\nERROR : %s", *csvQuiz, err))
	}

	// csv reader takes io.reader && parse entire file
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil{
		exit( fmt.Sprintf("Failed to parse csv file %s",*csvQuiz))
	}

	quiz := parseLines(lines)
	var counter int


	for i, q := range quiz{
		fmt.Printf("Question %d: %s?  ", i+1,q.question)
		var answer string
		// Scanf will trim input/ spaces
		fmt.Scanf("%s\n", &answer)
		if answer == q.answer{
			fmt.Printf("Correct\n")
			counter++
		} else {
			fmt.Printf("Inorrect\n")
		}
	}

	fmt.Printf("Score: %d / %d\n", counter, len(quiz))


}



// Functions

// func quizzical(){}	

func exit(msg string){
	log.Printf(msg)
	os.Exit(1)
}


func parseLines(lines [][]string) []problem {
	//takes lines 2d string slice // returns problems
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem {
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return problems

}



// Structs

//Define Problem Object
type problem struct {
	question string
	answer string 
}

