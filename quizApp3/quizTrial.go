//credits to:https://dev.to/adil_w3nomad/gopher-gym-quiz-game-part-1-4lbo

package main

import (
"bufio"
"encoding/csv"
"fmt"
"log"
"os"
"strings"
)

type quizItem struct {
  question string
  answer   string
}

func main() {
  csvFile, err := os.Open("quiz.csv")
  if err != nil {
    log.Fatal(err)
  }
  defer csvFile.Close()

  reader := csv.NewReader(csvFile)
  records, err := reader.ReadAll()

  if err != nil {
    log.Fatal(err)
  }

  for i := 0; i < len(records); i++ {
    // Create quizItem object
    quizItem := quizItem{records[i][0], records[i][1]}
    // Print out the question
    fmt.Println("Question:", quizItem.question)
    // Create reader and allow user to input their answer
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your answer now: ")
    // Expect answer to be given once they hit return
    text, _ := reader.ReadString('\n')
    fmt.Println("Your answer is:", text)
    // Trim the newline suffix from the input
    text = strings.TrimSuffix(text, "\n")
    if text == quizItem.answer {
      fmt.Println("Correct! Well done")
    } else {
      fmt.Println("WRONG! Answer is:", quizItem.answer)
    }
  }

}