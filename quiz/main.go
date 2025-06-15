package main

import (
	"flag"
	"strings"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
	"regexp"
)

func main() {
    // Help flag
	csvQuiz := flag.String("csv", "problems.csv", "csv file containg questions and answers in format `[question],[answer]`")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	flag.Parse()
	lines := parseCSV(csvQuiz)
	quizQuestions := parseLines(lines)
	quizzical(quizQuestions, timeLimit)
}

// Functions
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

func parseCSV(csvQuiz *string) [][]string {
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
	return lines 
}


func quizzical(quiz []problem, timeLimit *int ){
	// Quiz functionality, cycle through questions, get user input, feedback correctness 
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) 

	var counter int

	for i, q := range quiz{
		fmt.Printf("Question %d: %s?  ", i+1,q.question)
		answerChannel := make(chan string)
		go func(){
			var answer string
			// Scanf will trim input/ spaces
			fmt.Scanf("%s\n", &answer)
			isNumeric(answer)
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
				fmt.Printf("Incorrect\n")
			}
		}
	}
	fmt.Printf("Score: %d / %d\n", counter, len(quiz))
}	

func isNumeric(answer string){
	// Check input is a number if not break program with warn
	// scanf already trimmed whitespace
    if(!(regexp.MustCompile(`\d`).MatchString(answer))){
		exit("Invalid Input : Must be a numeric value")
	}
}


// Structs

//Define Problem Object
type problem struct {
	question string
	answer string 
}

