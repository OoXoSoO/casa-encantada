package main

import "fmt"

func main() {

	sc := NewScreen()
	sc.Print()
	for {
		key := readKeyboard()
		sc.Move(key)
		sc.Print()
		if sc.Game() {
			fmt.Println("Congratulaciones")
			return
		}
	}

}
