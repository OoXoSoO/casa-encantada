package main

import (
	"math/rand"
	"time"
)

type Quizzer struct {
	list []Quizz
}

func NewQuizzer() Quizzer {
	return Quizzer{
		list: []Quizz{
			{
				Text:     "¿2 > 4?",
				Response: false,
			},
			{
				Text:     "¿2 = 4?",
				Response: false,
			},
			{
				Text:     "¿2 < 4?",
				Response: true,
			},
			{
				Text:     "¿4 >= 4?",
				Response: true,
			},
		},
	}
}

func (q Quizzer) NewQuizz() Quizz {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	return q.list[ra.Intn(len(q.list)-1)]
}

type Quizz struct {
	Text     string
	Response bool
}
