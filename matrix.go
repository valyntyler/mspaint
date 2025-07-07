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
