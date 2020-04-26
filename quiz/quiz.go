package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	defaultfilename = "problems.csv"
)

var filename = flag.String("file", defaultfilename, "the file you want to use")

func main()  {
	flag.Parse()

	// opening the file
	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("couldn't open file", err)
	}
	records, _ := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		log.Fatalln("error reading file contents",err)
	}
	score := 0
	total := len(records)
	for _, val := range records {
		fmt.Print(val[0]+"=")
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Invalid Input", err)
		}
		answer = trim(answer)
		expectedAns := trim(val[1])

		if answer ==  expectedAns {
			score++
		}
	}
	fmt.Printf("Score: %d correct out of %d\n", score, total)
}

func trim(s string) string {
	t := strings.Trim(s, "\n")
	t = strings.Trim(t, " ")
	return t
}