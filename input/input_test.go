package input

import (
	"testing"

	"github.com/nsf/termbox-go"
)

var table = []struct {
	key string
	in  termbox.Event
	out byte
}{
	{"s", termbox.Event{Ch: 's'}, KeyShuffle},
	{"q", termbox.Event{Ch: 'q'}, KeyQuit},
	{"up", termbox.Event{Ch: 0, Key: termbox.KeyArrowUp}, KeyUp},
	{"right", termbox.Event{Ch: 0, Key: termbox.KeyArrowRight}, KeyRight},
	{"down", termbox.Event{Ch: 0, Key: termbox.KeyArrowDown}, KeyDown},
	{"left", termbox.Event{Ch: 0, Key: termbox.KeyArrowLeft}, KeyLeft},
	{"ctrlc", termbox.Event{Ch: 0, Key: termbox.KeyCtrlC}, KeyQuit},
	{"esc", termbox.Event{Ch: 0, Key: termbox.KeyEsc}, KeyQuit},
	{"any", termbox.Event{Ch: 'b'}, 0},
}

func Test_getAction(t *testing.T) {
	for _, test := range table {
		res := getAction(test.in)
		if res != test.out {
			t.Errorf("error on key '%s': expected %d got %d", test.key, test.out, res)
		}
	}
}
