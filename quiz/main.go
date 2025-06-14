package main

import (
	"flag"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func quizzical(){
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

	fmt.Println(lines)
}	

func exit(msg string){
	log.Printf(msg)
	os.Exit(1)
}

func main() {
    quizzical()  
}