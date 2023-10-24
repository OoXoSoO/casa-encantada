package main

import "fmt"

type Challenger struct {
	quizzer Quizzer
}

func NewChallenger() Challenger {
	return Challenger{
		quizzer: NewQuizzer(),
	}
}
func (c Challenger) Challenge() bool {
	success := false
	for !success {
		success = c.triggerChallenge()
	}
	return success
}

func (c Challenger) triggerChallenge() bool {
	q := c.quizzer.NewQuizz()
	r := getResponse(q)
	return r == q.Response
}

func getResponse(q Quizz) bool {
	// y = true
	// n = false
	// any other value, relauch the request
	for {
		fmt.Println(q.Text)
		key := readKeyboard()
		switch key {
		case KeyY:
			return true
		case KeyN:
			return false
		}
		fmt.Println("Invalid response, please try again")
	}

}
