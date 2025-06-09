package main

import (
	"fmt"
	"log"
)

func main() {
	quiz := NewQuiz()
	err := quiz.load_from_file("irc.json")
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	for _, question := range quiz.Questions {
		fmt.Printf("Question: %s\n", question.Question)
		fmt.Printf("Answer: %s\n", question.Answer)
		if len(question.Options) > 0 {
			fmt.Printf("Options: %v\n", question.Options)
		}
		fmt.Println()
	}
}
