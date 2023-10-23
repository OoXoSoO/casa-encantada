package main

import (
	"math/rand"
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
	Visible bool
}

func NewCell() Cell {
	t := rand.Intn(99)
	callType := CellTypeBlank
	if t <= 0 {
		callType = CellTypeGhost
	}
	return Cell{
		Type:    callType,
		Visible: false,
	}
}

func (c Cell) Print() {
	if !c.Visible {
		print("❓")
		return
	}
	/*🚪⬜️👻🍭*/
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
