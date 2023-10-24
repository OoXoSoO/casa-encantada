package main

import "fmt"

func main() {
	sc := NewScreen(4, 4)
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
