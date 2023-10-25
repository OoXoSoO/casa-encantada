package keyboard

import (
	"casa/pkg"

	"github.com/eiannone/keyboard"
)

type keyBoarReader struct {
}

var _ pkg.KeyBoarReader = (*keyBoarReader)(nil)

func NewKeyBoardReader() *keyBoarReader {
	keyboard.Open()
	return &keyBoarReader{}
}

func (kbr *keyBoarReader) ReadKeyboard() pkg.Key {
	r, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	switch key {
	case keyboard.KeyArrowDown:
		return pkg.KeyArrowDown
	case keyboard.KeyArrowUp:
		return pkg.KeyArrowUp
	case keyboard.KeyArrowRight:
		return pkg.KeyArrowRight
	case keyboard.KeyArrowLeft:
		return pkg.KeyArrowLeft
	}
	if r == 121 {
		return pkg.KeyY
	}
	if r == 110 {
		return pkg.KeyN
	}

	return pkg.KeyArrowUnknown
}
