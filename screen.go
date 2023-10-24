package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Screen struct {
	cells      [][]Cell
	cellCandy  *Cell
	challenger Challenger
	maxRows    int
	maxCol     int
}

func NewScreen(maxRows int, maxCol int) *Screen {

	ra := rand.New(rand.NewSource(time.Now().UnixNano()))

	ret := Screen{
		cells:      make([][]Cell, maxRows+1),
		maxRows:    maxRows,
		maxCol:     maxCol,
		challenger: NewChallenger(),
	}
	for rowIdx := range ret.cells {
		row := make([]Cell, ret.maxCol+1)
		for colIdx := range row {
			row[colIdx] = NewCell()
		}
		ret.cells[rowIdx] = row
	}

	doorPosRow := ra.Intn(ret.maxRows)
	doorPosCol := ra.Intn(ret.maxCol)
	ret.cells[doorPosRow][doorPosCol] = Cell{
		Type:    CellTypeDoor,
		Success: true,
		Busy:    true,
	}

	candyPosRow := doorPosRow
	candyPosCol := doorPosCol

	for candyPosCol == doorPosCol && candyPosRow == doorPosRow {
		candyPosRow = ra.Intn(ret.maxRows)
		candyPosCol = ra.Intn(ret.maxCol)
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
	sc.print(true)
}
func (sc *Screen) PrintFinal() {
	sc.print(false)
}
func (sc *Screen) print(printBusy bool) {
	runCmd("cmd", "/c", "cls")
	for _, row := range sc.cells {
		for _, col := range row {
			col.Print(printBusy)
		}
		println("")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func (sc *Screen) Move() {

	key := readKeyboard()
	if key == KeyArrowUnknown {
		return
	}

	busyRow, busyCol := sc.findBusy()
	// Verifica si se presionó la tecla ESC para salir del bucle
	newBusyRow := busyRow
	newBusyCol := busyCol
	switch key {
	case KeyArrowDown:
		if busyRow+1 <= sc.maxRows {
			newBusyRow = busyRow + 1
		}
	case KeyArrowUp:
		if busyRow-1 >= 0 {
			newBusyRow = busyRow - 1
		}
	case KeyArrowRight:
		if busyCol+1 <= sc.maxCol {
			newBusyCol = busyCol + 1
		}
	case KeyArrowLeft:
		if busyCol-1 >= 0 {
			newBusyCol = busyCol - 1
		}
	}

	if sc.cells[busyRow][busyCol].Success || (!sc.cells[busyRow][busyCol].Success && sc.cells[newBusyRow][newBusyCol].Success) {
		sc.cells[busyRow][busyCol].Busy = false
		sc.cells[newBusyRow][newBusyCol].Busy = true
		return
	}
	fmt.Printf("invalid movement")

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

func (sc *Screen) EvaluateAction() {
	busyRow, busyCol := sc.findBusy()
	if sc.cells[busyRow][busyCol].Success {
		return
	}
	success := sc.challenger.Challenge()
	if success {
		sc.cells[busyRow][busyCol].Success = true
	}
}
