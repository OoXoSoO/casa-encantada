package main

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
		},
	}
}

func (q Quizzer) NewQuizz() Quizz {
	return q.list[0]
}

type Quizz struct {
	Text     string
	Response bool
}
