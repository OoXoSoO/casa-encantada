package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Screen struct {
	cells     [][]Cell
	cellCandy *Cell
}

const MaxRows = 3
const MaxCol = 3

func NewScreen() *Screen {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))

	ret := Screen{
		cells: make([][]Cell, MaxRows+1),
	}
	for rowIdx := range ret.cells {
		row := make([]Cell, MaxCol+1)
		for colIdx := range row {
			row[colIdx] = NewCell()
		}
		ret.cells[rowIdx] = row
	}

	doorPosRow := ra.Intn(MaxRows)
	doorPosCol := ra.Intn(MaxCol)
	ret.cells[doorPosRow][doorPosCol] = Cell{
		Type:    CellTypeDoor,
		Success: true,
		Busy:    true,
	}

	candyPosRow := doorPosRow
	candyPosCol := doorPosCol

	for candyPosCol == doorPosCol && candyPosRow == doorPosRow {
		candyPosRow = ra.Intn(MaxRows)
		candyPosCol = ra.Intn(MaxCol)
	}
	ret.cells[candyPosRow][candyPosCol] = Cell{
		Type:    CellTypeCandy,
		Success: false,
	}
	ret.cellCandy = &ret.cells[candyPosRow][candyPosCol]

	return &ret
}
func (sc *Screen) Game() bool {
	return sc.cellCandy.Busy
}
func (sc *Screen) Print() {
	runCmd("cmd", "/c", "cls")
	for _, row := range sc.cells {
		for _, col := range row {
			col.Print()
		}
		println("")
	}
}
func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func (sc *Screen) Move(key Key) {

	if key == KeyArrowUnknown {
		return
	}

	busyRow, busyCol := sc.findBusy()
	// Verifica si se presionó la tecla ESC para salir del bucle
	newBusyRow := busyRow
	newBusyCol := busyCol
	switch key {
	case KeyArrowDown:
		if busyRow+1 <= MaxRows {
			newBusyRow = busyRow + 1
		}
	case KeyArrowUp:
		if busyRow-1 >= 0 {
			newBusyRow = busyRow - 1
		}
	case KeyArrowRight:
		if busyCol+1 <= MaxCol {
			newBusyCol = busyCol + 1
		}
	case KeyArrowLeft:
		if busyCol-1 >= 0 {
			newBusyCol = busyCol - 1
		}
	}

	sc.cells[busyRow][busyCol].Busy = false
	sc.cells[newBusyRow][newBusyCol].Busy = true

	fmt.Printf("old busy = (%d,%d) %t, move to (%d,%d) %t \n", busyRow, busyCol, sc.cells[busyRow][busyCol].Busy, newBusyRow, newBusyCol, sc.cells[newBusyRow][newBusyCol].Busy)

}
func (sc *Screen) findBusy() (int, int) {
	for rowIdx, row := range sc.cells {
		for colIdx, col := range row {
			if col.Busy {
				return rowIdx, colIdx
			}
		}
	}
	panic("invalid busy state")
}

/*
 * Este es un reto especial por Halloween.
 * Te encuentras explorando una mansión abandonada llena de habitaciones.
 * En cada habitación tendrás que resolver un acertijo para poder avanzar a la siguiente.
 * Tu misión es encontrar la habitación de los dulces.
 *
 * Se trata de implementar un juego interactivo de preguntas y respuestas por terminal.
 * (Tienes total libertad para ser creativo con los textos)
 *
 * - 🏰 Casa: La mansión se corresponde con una estructura cuadrada 4 x 4
 *   que deberás modelar. Las habitaciones de puerta y dulces no tienen enigma.
 *   (16 habitaciones, siendo una de entrada y otra donde están los dulces)
 *   Esta podría ser una representación:
 *   🚪⬜️⬜️⬜️
 *   ⬜️👻⬜️⬜️
 *   ⬜️⬜️⬜️👻
 *   ⬜️⬜️🍭⬜️
 * - ❓ Enigmas: Cada habitación propone un enigma aleatorio que deberás responder con texto.
 *   Si no lo aciertas no podrás desplazarte.
 * - 🧭 Movimiento: Si resuelves el enigma se te preguntará a donde quieres desplazarte.
 *   (Ejemplo: norte/sur/este/oeste. Sólo deben proporcionarse las opciones posibles)
 * - 🍭 Salida: Sales de la casa si encuentras la habitación de los dulces.
 * - 👻 (Bonus) Fantasmas: Existe un 10% de que en una habitación aparezca un fantasma y
 *   tengas que responder dos preguntas para salir de ella.
 */
