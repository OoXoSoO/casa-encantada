package main

import (
	"math/rand"
	"time"
)

const (
	CellTypeDoor  CellType = 0
	CellTypeBlank CellType = 1
	CellTypeGhost CellType = 2
	CellTypeCandy CellType = 3
)

type CellType int

type Cell struct {
	Type    CellType
	Success bool
	Busy    bool
}

func NewCell() Cell {
	ra := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(99)
	callType := CellTypeBlank
	if ra <= 9 {
		callType = CellTypeGhost
	}
	return Cell{
		Type:    callType,
		Success: false,
		Busy:    false,
	}
}

func (c Cell) Print(printBusy bool) {
	/*🚪⬜️👻🍭🎃*/
	if printBusy && c.Busy {
		print("🎃")
		return
	}
	if !c.Success {
		print("❓")
		return
	}
	switch c.Type {
	case CellTypeBlank:
		print("⬜️")
		return
	case CellTypeCandy:
		print("🍭")
		return
	case CellTypeGhost:
		print("👻")
		return
	case CellTypeDoor:
		print("🚪")
		return
	}
}
