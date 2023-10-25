package main

import (
	"casa/keyboard"
	"casa/pkg"
	"casa/quizz"
	"fmt"
)

func main() {
	kbreader := keyboard.NewKeyBoardReader()
	challenger := quizz.NewChallenger(kbreader)
	sc := pkg.NewScreen(4, 4, challenger, kbreader)

	sc.Print()
	for !sc.Game() {
		sc.Move()
		sc.Print()
		sc.EvaluateAction()
		sc.Print()
	}
	sc.PrintFinal()
	fmt.Println("Congratulaciones")
}
