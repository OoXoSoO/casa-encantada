package main

import (
	"math/rand"
	"time"
)

type Screen [][]Cell

const MaxRows = 4
const MaxCol = 4

func NewScreen() *Screen {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))

	ret := make(Screen, MaxRows)
	for rowIdx := range ret {
		row := make([]Cell, MaxCol)
		for colIdx := range row {
			row[colIdx] = NewCell()
		}
		ret[rowIdx] = row
	}

	doorPosRow := ra.Intn(MaxRows - 1)
	doorPosCol := ra.Intn(MaxCol - 1)
	ret[doorPosRow][doorPosCol] = Cell{
		Type:    CellTypeDoor,
		Visible: true,
	}

	candyPosRow := doorPosRow
	candyPosCol := doorPosCol

	for candyPosCol == doorPosCol && candyPosRow == doorPosRow {
		candyPosRow = ra.Intn(MaxRows - 1)
		candyPosCol = ra.Intn(MaxCol - 1)
	}
	ret[candyPosRow][candyPosCol] = Cell{
		Type:    CellTypeCandy,
		Visible: true,
	}

	return &ret
}

func (sc *Screen) Print() {
	for _, row := range *sc {
		for _, col := range row {
			col.Print()
		}
		println("")
	}
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
