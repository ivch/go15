package grid

import (
	"fmt"
	"math/rand"
	"time"
)

const shuffleTimeLimit = 100

//Grid represents grid type. made as [][]byte to represent matrix
type Grid [][]byte

//Init generates gameGrid and resultGrid of given size
func Init(size byte) (Grid, Grid) {
	grid := make(Grid, size)

	num := byte(0)

	for y := byte(0); y < size; y++ {
		grid[y] = make([]byte, size)
		for x := byte(0); x < size; x++ {
			num++
			grid[y][x] = num

			if num == size*size {
				grid[y][x] = 0
			}
		}
	}

	shuffled := shuffle(grid)

	return grid, shuffled
}

//Reload shuffles givend grid to restart the game
func Reload(g Grid) Grid {
	return shuffle(g)
}

//String returns grid in string representation in order to satisfy fmt.Stringer interface
func (g Grid) String() string {
	var s string
	for _, row := range g {
		var r string
		for _, val := range row {
			var tile string
			if val > 0 && val < 10 {
				tile = fmt.Sprintf("[  %d ]", val)
			}

			if val >= 10 {
				tile = fmt.Sprintf("[ %d ]", val)
			}

			if val == 0 {
				tile = "[    ]"
			}

			r = fmt.Sprintf("%s %s", r, tile)
		}
		s = fmt.Sprintf("%s %s\n", s, r)

	}

	return s
}

func shuffle(g Grid) Grid {
	t := time.NewTimer(time.Millisecond * shuffleTimeLimit)

	done := make(chan bool)

	shuffledGrid := Copy(g)

	go func() {
		for {
			select {
			case <-t.C:
				done <- true
			default:
				shuffledGrid = switchTiles(shuffledGrid)
			}
		}
	}()

	<-done

	return shuffledGrid
}

func switchTiles(g Grid) Grid {
	rand.Seed(time.Now().UnixNano())

	x1, y1 := byte(rand.Intn(len(g))), byte(rand.Intn(len(g)))
	x2, y2 := byte(rand.Intn(len(g))), byte(rand.Intn(len(g)))

	grid := Copy(g)
	grid[y1][x1], grid[y2][x2] = g[y2][x2], g[y1][x1]

	return grid
}

//Copy performs deep copy of given grid
func Copy(g Grid) Grid {
	ng := make(Grid, len(g))
	for y := range g {
		ng[y] = make([]byte, len(g[y]))
		copy(ng[y], g[y])
	}
	return ng
}

//GetZeroTilePosition returns x,y coordinates of empty tile on game field
func GetZeroTilePosition(g Grid) (byte, byte, error) {
	for y, row := range g {
		for x, val := range row {
			if val == 0 {
				return byte(x), byte(y), nil
			}
		}
	}

	return 0, 0, fmt.Errorf("unable to find zero tile")
}
