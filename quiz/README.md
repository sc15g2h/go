# Numeric Quiz

## Compile and Run Code 
Compile quiz by running `go build main.go`

Play quiz by running `./main` 

### FLags 
`-csv`
By default the quiz will use `problems.csv` 
To use your own csv file use the flag `-csv` and specify your csv file `./main -csv=otherCsvFile.csv` 

`-limit`
By default the quiz has a 30 second time limit 
To specifiy you own time limit use the flag `-limit` and specify your time limit in seconds `./main -limit=60`


## CSV Format 
The format of the csv file is `question,answer`. The answer must be a number. 
