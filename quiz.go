package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"
	"io/ioutil"
)

// define a structure for the quiz
type Quiz struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	Question string `json:"question"`
	Answer   []string `json:"answer"`
	Options []string `json:"options,omitempty"` //optional field
}

/*
	create quiz question structure instance
*/
func NewQuiz() *Quiz {
	return &Quiz{
		Questions: make([]Question,20),
	}
}

func (quiz *Quiz) run() error {
	var input string 
	var isgood bool
	var sinput []string

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
			if input == "exit" || input == "quit" {
				log.Println("Exiting quiz.")
				return nil
			}
			sinput := strings.split(input, ",") // take only the first line of input
			if len(sinput) != len(question.Answer) {
				log.Printf(
					"Incorrect number of answers provided. Expected %d, got %d.\n", 
					len(question.Answer), len(sinput)
				)
				continue
			}
			isgood = true
		}

		isgood = false

		// loop through input and answers checking for equality ignoring case
		for idx, ans := range inputs {
			ans = strings.TrimSpace(ans) // remove leading and trailing whitespace
			rans = strings.TrimSpace(question.Answer[idx]) // get the answer from the question
			
			if strings.EqualFold(ans, answer) {
				fmt.Println("Correct: ",answer)
			} else {
				fmt.Println("Incorrect, the correct answer is:", answer)
			}
		}

		fmt.Println()
	}
	return nil
}

/*
Function to load quiz from json file, file name is passed to function
structure is:
{
	questions: {
		"question1": {
			"answer": "answer1",
			"question": "question1"
		}
		"question2": {
			"answer": "answer2",
			"question": "question2"
		}
		...
	}
}
*/
func (quiz *Quiz) load_from_file(file_name string) error {
	//verify input
	if file_name == "" {
		//create error message
		return errors.New("file name cannot be empty")
	}
	
	//read contents of file
	data,err := ioutil.ReadFile(file_name)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to read file %s: %v", file_name, err))
	}

	//unmarshal json data into quiz map
	err = json.Unmarshal(data, &quiz)
	if err != nil {
		return errors.New(
			fmt.Sprintf(
				"failed to unmarshal json data from file %s: %v", file_name, err
			)
		)
	}

	// check if the quiz map is empty
	if len(quiz.Questions) == 0 {
		return errors.New("quiz is empty, no questions found in file")
	}

	// If we reach here, the quiz has been successfully loaded
	log.Printf("Quiz loaded successfully from file: %s", file_name)

	return nil
}
