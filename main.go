package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	quiz := NewQuiz()
	var input string 
	var answer string 
	var isgood bool

	err := quiz.load_from_file("irc.json")
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}

	for _, question := range quiz.Questions {

		fmt.Printf("Question: %s\n", question.Question)
		if len(question.Options) > 0 {
			fmt.Printf("Options: %v\n", question.Options)
		}

		//read line from console
		for isgood == false {
			fmt.Scanln(&input)
			if input == "" {	
				log.Println("No input provided, skipping question.")
				continue
			}
			isgood = true
		}

		isgood = false
		// clean up input
		input = strings.TrimSpace(input) // remove leading and trailing whitespac
		//clean up answer
		answer = strings.TrimSpace(question.Answer) // remove leading and trailing whitespac

		if strings.EqualFold(input, answer) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect, the correct answer is:", answer)
		}

		fmt.Println()
	}
}
