//credits to:https://github.com/LordRahl90/quizmanager/blob/master/readquestion/main.go

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

func main() {

	filename := flag.String("q", "quiz.csv", "Please provide the questions file")
	flag.Parse()

	csvFile, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	l := len(records)
	pass := 0

	r := bufio.NewReader(os.Stdin)

	for i := 0; i < l; i++ {
		q := records[i][0]
		ans := records[i][1]
		fmt.Printf("What is the Answer to %s\n", q)
		resp, _ := r.ReadString('\n')
		resp = strings.Trim(resp, "\n")

		if resp == ans {
			pass++
			fmt.Printf("Very Correct!!!\n")
		} else {
			fmt.Printf("Sadly, You are wrong. The right answer is: %s\n", ans)
		}

		// fmt.Printf("Your Response is: %s, While the answer is: %s\n", strings.Trim(resp, "\n"), ans)
	}

	fail := l - pass
	percent := float64(pass) / float64(l) * 100.0

	fmt.Printf("Well Done, You passed %d and failed: %d, Total Pass Percentage is %.1f% \n", pass, fail, percent)
}