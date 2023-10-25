package quizz

import (
	"casa/pkg"
	"fmt"
)

type Challenger struct {
	quizzer  Quizzer
	kbReader pkg.KeyBoarReader
}

func NewChallenger(kbReader pkg.KeyBoarReader) Challenger {
	return Challenger{
		quizzer:  NewQuizzer(),
		kbReader: kbReader,
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
	r := c.getResponse(q)
	return r == q.Response
}

func (c Challenger) getResponse(q Quizz) bool {
	// y = true
	// n = false
	// any other value, relauch the request
	for {
		fmt.Println(q.Text)
		key := c.kbReader.ReadKeyboard()
		switch key {
		case pkg.KeyY:
			return true
		case pkg.KeyN:
			return false
		}
		fmt.Println("Invalid response, please try again")
	}

}
