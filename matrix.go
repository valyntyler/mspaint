package main

import (
	"github.com/gdamore/tcell/v2"
)

type matrix [][]rune

func newMatrix(width, height int, fill rune) matrix {
	m := make(matrix, height)
	for i := range m {
		m[i] = make([]rune, width)
	}
	for i, row := range m {
		for j := range row {
			m[i][j] = fill
		}
	}
	return m
}

func (m matrix) screenToMatrix(screen tcell.Screen, x, y int) (int, int) {
	_, ymax := screen.Size()

	xnew := x/2 - ymax/2 - m.width()/2
	ynew := y - ymax/2 + m.height()/2

	return xnew, ynew
}

func (m matrix) isInside(x, y int) bool {
	return !(y < 0 || x < 0 || y >= len(m) || x >= len(m[0]))
}

func (m matrix) getCell(x, y int) rune {
	if m.isInside(x, y) {
		return m[y][x]
	} else {
		return ' '
	}
}

func (m matrix) setCell(x, y int, value rune) {
	if !m.isInside(x, y) {
		return
	}

	m[y][x] = value
}

func (m matrix) width() int {
	return len(m[0])
}

func (m matrix) height() int {
	return len(m)
}

func (m matrix) draw(screen tcell.Screen, style tcell.Style) {
	xmax, ymax := screen.Size()

	for x := range m.width() {
		for y := range m.height() {
			xnew := xmax/2 - m.width() + x*2
			ynew := ymax/2 - m.height()/2 + y

			screen.SetContent(xnew, ynew, m[y][x], nil, style)
		}
	}
}
