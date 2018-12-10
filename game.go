package go15

import (
	"reflect"
	"time"

	"github.com/ivch/go15/grid"
	"github.com/ivch/go15/input"
	"github.com/ivch/go15/printer"
)

//Game describes game interface
type Game interface {
	Run()
	Stop()
}

type game struct {
	size       byte
	gameGrid   [][]byte
	resultGrid [][]byte
	printer    chan string
	input      chan byte
	exit       chan struct{}
}

//Run is a main game runner. Processes game logic and ticks
func (g *game) Run() {
	go printer.Listen(g.printer)
	go input.Listen(g.input, g.exit)

	go func() {
		gameGrid := grid.Copy(g.gameGrid)
		for {
			g.printer <- gameGrid.String()
			action := <-g.input

			if action == input.KeyQuit {
				g.Stop()
			}

			if action == input.KeyShuffle {
				gameGrid = grid.Restart(gameGrid)
				continue
			}

			var err error
			gameGrid, err = g.move(gameGrid, action)
			if err != nil {
				g.printer <- err.Error()
				time.Sleep(2 * time.Second)
			}

			if reflect.DeepEqual(gameGrid, g.resultGrid) {
				g.printer <- "Winner Winner Chicken Dinner"
				g.Stop()
				break
			}
		}
	}()

	<-g.exit
}

func (g *game) move(currGrid grid.Grid, direction byte) (grid.Grid, error) {
	startX, startY, err := grid.GetZeroTilePosition(currGrid)
	if err != nil {
		return currGrid, err
	}

	if !moveIsPossible(startX, startY, direction, g.size) {
		return currGrid, nil
	}

	endX, endY := getNextCoords(startX, startY, direction)

	newGrid := grid.Copy(currGrid)

	newGrid[endY][endX], newGrid[startY][startX] = currGrid[startY][startX], currGrid[endY][endX]

	return newGrid, nil
}

//Init initalizes the game
func Init(size byte) Game {
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

//Stop is called to stop the game
func (g *game) Stop() {
	g.exit <- struct{}{}
}

func moveIsPossible(x, y, direction byte, size byte) bool {
	switch direction {
	case input.KeyUp:
		return y < size-1
	case input.KeyDown:
		return y > 0
	case input.KeyLeft:
		return x < size-1
	case input.KeyRight:
		return x > 0
	}

	return false
}

func getNextCoords(x, y, direction byte) (byte, byte) {
	switch direction {
	case input.KeyUp:
		y++
	case input.KeyDown:
		y--
	case input.KeyLeft:
		x++
	case input.KeyRight:
		x--
	}

	return x, y
}
