//credits to:https://medium.com/@moriagape/how-to-read-csv-files-in-golang-a-quiz-app-b5c8891207a0

package main

import (
	"encoding/csv"
	"flag"
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
)


type csvLine struct {
	Questions string
	Answers   string
}

var score int

func processQuiz(lines [][]string) int {
	var total int
	for _, line := range lines {
		total++
		data := csvLine{
			Questions: line[0],
			Answers:   line[1],
		}
		fmt.Println(data.Questions)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Your answer: ")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error occured with answer input", err)
		}
		fmt.Println("Your answer is:", response, "Correct answer is:", data.Answers)
		processResult(response, data.Answers)
	}
	return total
}

func increment() {
	score = score + 1
}

func processResult(response string, answer string) {
	res := strings.TrimSpace(response)
	ans := strings.TrimSpace(answer)
	if res == ans {
		score++
	}
}



func processFile(filename *string) [][]string {
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
func main() {
	filename := flag.String("csv", "quiz.csv", "a csv file in the format of 'question,answer'")
	var totalScore int
	lines := processFile(filename)
	totalScore = processQuiz(lines)
	fmt.Println("Number of correct questions is ", score, "Total question is", totalScore)
}