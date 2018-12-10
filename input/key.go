package input

import (
	"log"
	"unicode"

	"github.com/nsf/termbox-go"
)

const (
	KeyUp byte = iota + 1
	KeyRight
	KeyDown
	KeyLeft
	KeyShuffle
	KeyQuit
)

func Listen(input chan byte, exit chan struct{}) {
	for {
		e := termbox.PollEvent()

		switch e.Type {
		case termbox.EventKey:
			input <- getAction(e)
		case termbox.EventInterrupt:
			exit <- struct{}{}
			break
		case termbox.EventError:
			log.Fatal(e.Err)
		}
	}
}

func getAction(e termbox.Event) byte {
	switch unicode.ToLower(e.Ch) {
	case 's':
		return KeyShuffle
	case 'q':
		return KeyQuit
	case 0:
		switch e.Key {
		case termbox.KeyArrowUp:
			return KeyUp
		case termbox.KeyArrowRight:
			return KeyRight
		case termbox.KeyArrowDown:
			return KeyDown
		case termbox.KeyArrowLeft:
			return KeyLeft
		case termbox.KeyCtrlC, termbox.KeyEsc:
			return KeyQuit
		}
	}

	return 0
}
