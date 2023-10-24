package main

import (
	"github.com/eiannone/keyboard"
)

const (
	KeyArrowUnknown Key = -1
	KeyArrowUp      Key = 0
	KeyArrowDown    Key = 1
	KeyArrowLeft    Key = 2
	KeyArrowRight   Key = 3
)

type Key int

func init() {
	keyboard.Open()
}
func readKeyboard() Key {
	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	switch key {
	case keyboard.KeyArrowDown:
		return KeyArrowDown
	case keyboard.KeyArrowUp:
		return KeyArrowUp
	case keyboard.KeyArrowRight:
		return KeyArrowRight
	case keyboard.KeyArrowLeft:
		return KeyArrowLeft
	}
	return KeyArrowUnknown
}
