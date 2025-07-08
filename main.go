package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

func recurse(m matrix, screen tcell.Screen, style tcell.Style, x, y int) {
	if m.getCell(x, y) == '.' {
		m.setCell(x, y, '*')
		m.draw(screen, style)
		screen.Show()
		time.Sleep(20 * time.Millisecond)
		recurse(m, screen, style, x, y+1)
		recurse(m, screen, style, x, y-1)
		recurse(m, screen, style, x+1, y)
		recurse(m, screen, style, x-1, y)
	}
}

func main() {
	m := newMatrix(24, 24, '.')

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	recurse(m, s, defStyle, m.width()/2, m.height()/2)

	// Event loop
	for {
		// Update screen
		s.Show()
		m.draw(s, defStyle)

		// Poll and process event
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'Q' || ev.Rune() == 'q' {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		case *tcell.EventMouse:
			xscreen, yscreen := ev.Position()
			x, y := m.screenToMatrix(s, xscreen, yscreen)

			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				m.setCell(x, y, 'x')
			}
		}
	}
}
