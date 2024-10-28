package maze

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// Maze хранит двумерную таблицу клеток, высоту и ширину.
type Maze struct {
	Cells  map[cells.Coordinates]*cells.Cell
	Height int
	Width  int
}

// New возвращает инициализированный Maze.
func New(height, width int) Maze {
	maze := Maze{
		Cells:  make(map[cells.Coordinates]*cells.Cell, height*width),
		Height: height,
		Width:  width,
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			maze.Cells[cells.Coordinates{X: x, Y: y}] = &cells.Cell{}
		}
	}

	return maze
}
