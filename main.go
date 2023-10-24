package main

import "fmt"

func main() {
	sc := NewScreen()
	sc.Print()
	for !sc.Game() {
		sc.Move()
		sc.Print()
	}
	fmt.Println("Congratulaciones")
}
