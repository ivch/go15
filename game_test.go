package go15

import (
	"reflect"
	"testing"
	"time"

	"github.com/ivch/go15/grid"
	"github.com/ivch/go15/input"
)

func TestGameMove(t *testing.T) {
	var fieldSize byte = 3
	game := initTestGame(fieldSize)

	currGrid := grid.Copy(game.gameGrid)
	startX, startY, err := grid.GetZeroTilePosition(currGrid)
	if err != nil {
		t.Fatal(err)
	}
	direction := input.KeyUp
	if !moveIsPossible(startX, startY, input.KeyUp, fieldSize) {
		direction = input.KeyDown
	}

	endX, endY := getNextCoords(startX, startY, direction)

	newGrid := grid.Copy(currGrid)
	newGrid[endY][endX], newGrid[startY][startX] = currGrid[startY][startX], currGrid[endY][endX]

	resGrid, err := game.move(currGrid, direction)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(newGrid, resGrid) {
		t.Error("got different result for moves in test and in game")
	}
}

func TestGameMoveErrors(t *testing.T) {
	game := initTestGame(2)

	wrongGrid := grid.Grid{
		{1, 2},
		{3, 4},
	}

	_, err := game.move(wrongGrid, input.KeyUp)
	if err == nil {
		t.Error("should have got error 'unable to find zero tile")
	}

	wrongGrid[0][0] = 0
	newGrid, _ := game.move(wrongGrid, input.KeyDown)
	if !reflect.DeepEqual(wrongGrid, newGrid) {
		t.Error("got changed grid, should have got same one")
	}
}

func Test_gameStop(t *testing.T) {
	game := initTestGame(2)

	go game.Stop()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-game.exit:
		return
	default:
		t.Error("haven't received write to exit channel")
	}
}

func Test_moveIsPossible(t *testing.T) {
	var table = []struct {
		key string
		in  [4]byte
		out bool
	}{
		{"up", [4]byte{1, 0, input.KeyUp, 2}, true},
		{"up", [4]byte{1, 0, input.KeyDown, 2}, false},
		{"down", [4]byte{1, 1, input.KeyDown, 2}, true},
		{"down", [4]byte{1, 1, input.KeyUp, 2}, false},
		{"left", [4]byte{0, 0, input.KeyLeft, 2}, true},
		{"right", [4]byte{0, 0, input.KeyRight, 2}, false},
		{"left", [4]byte{1, 0, input.KeyLeft, 2}, false},
		{"right", [4]byte{1, 0, input.KeyRight, 2}, true},
		{"any", [4]byte{1, 0, 5, 2}, false},
	}

	var res bool
	for _, test := range table {
		res = moveIsPossible(test.in[0], test.in[1], test.in[2], test.in[3])

		if res != test.out {
			t.Errorf("error checking move '%s' from pos %d/%d for size %d: expected %t got %t", test.key, test.in[0], test.in[1], test.in[3], test.out, res)
		}
	}
}

func Test_getNextCoords(t *testing.T) {
	var table = []struct {
		key string
		in  [3]byte
		out [2]byte
	}{
		{"up", [3]byte{2, 2, input.KeyUp}, [2]byte{2, 3}},
		{"down", [3]byte{2, 2, input.KeyDown}, [2]byte{2, 1}},
		{"left", [3]byte{2, 2, input.KeyLeft}, [2]byte{3, 2}},
		{"right", [3]byte{2, 2, input.KeyRight}, [2]byte{1, 2}},
	}

	var res [2]byte
	for _, test := range table {
		res[0], res[1] = getNextCoords(test.in[0], test.in[1], test.in[2])

		if res != test.out {
			t.Errorf("error on move '%s': expected %d/%d got %d/%d", test.key, test.out[0], test.out[1], res[0], res[1])
		}
	}
}

func initTestGame(size byte) *game {
	resultGrid, gameGrid := grid.Init(size)

	return &game{
		size:       size,
		gameGrid:   gameGrid,
		resultGrid: resultGrid,
		printer:    make(chan string),
		input:      make(chan byte),
		exit:       make(chan struct{}),
	}
}
