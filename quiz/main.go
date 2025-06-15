package main

import (
	"flag"
	"strings"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
    // Help flag
	csvQuiz := flag.String("csv", "problems.csv", "csv file containg questions and answers in format `[question],[answer]`")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	flag.Parse()

	// *csvQuiz - pointer to a string 
	file, err := os.Open(*csvQuiz)
	if err != nil {
		exit( fmt.Sprintf("Failed to open csv file %s\nERROR : %s", *csvQuiz, err))
	}

	//
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) 
	// blocks until it gets message from the channel
	// <- timer.C

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
		answerChannel := make(chan string)
		go func(){
			var answer string
			// Scanf will trim input/ spaces
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <- timer.C:
			fmt.Printf("\nOh no! You ran out of time! ... Score: %d / %d\n", counter, len(quiz))
			return
		case answer := <- answerChannel:
			if answer == q.answer{
				fmt.Printf("Correct\n")
				counter++
			} else {
				fmt.Printf("Inorrect\n")
			}
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
			answer: strings.TrimSpace(line[1]), // string validator function?
		}
	}
	return problems

}


// To Do
// Add a string validator funtion?



// Structs

//Define Problem Object
type problem struct {
	question string
	answer string 
}

