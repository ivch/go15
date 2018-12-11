package grid

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	expectedResultGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	actualResultGrid, _ := Init(3)

	if !reflect.DeepEqual(expectedResultGrid, actualResultGrid) {
		t.Errorf("got different result grids:\nER:%s\nAR:%s", expectedResultGrid, actualResultGrid)
	}
}

func TestGetZeroTilePosition(t *testing.T) {
	grid := Grid{
		{1, 2},
		{3, 0},
	}

	expX, expY := byte(1), byte(1)
	actX, actY, err := GetZeroTilePosition(grid)
	if err != nil {
		t.Fatal(err)
	}

	if actX != expX || actY != expY {
		t.Errorf("got wrong zero point coords: expected %d/%d got %d/%d", expX, expY, actX, actY)
	}

	grid[1][1] = 4
	_, _, err = GetZeroTilePosition(grid)
	if err == nil {
		t.Error("got nil error, expected 'unable to find zero tile'")
	}
}

func TestGrid_String(t *testing.T) {
	grid := Grid{
		{1, 2},
		{13, 0},
	}

	var i interface{} = grid

	if _, ok := i.(fmt.Stringer); !ok {
		t.Fatal("should implement fmt.Stringer")
	}

	expected := "  [  1 ] [  2 ]\n  [ 13 ] [    ]\n"
	actual := grid.String()

	if expected != actual {
		t.Errorf("got wront grid string\nER\n%s\nAR\n%s", expected, actual)
	}
}

func TestReload(t *testing.T) {
	grid, _ := Init(4)
	newGrid := Reload(grid)

	if reflect.DeepEqual(grid, newGrid) {
		t.Error("reload should shuffle greed so two grids has to be unequal")
	}
}
