package pkg

const (
	KeyArrowUnknown Key = -1
	KeyArrowUp      Key = 0
	KeyArrowDown    Key = 1
	KeyArrowLeft    Key = 2
	KeyArrowRight   Key = 3
	KeyY            Key = 4
	KeyN            Key = 5
)

type KeyBoarReader interface {
	ReadKeyboard() Key
}
type Key int
