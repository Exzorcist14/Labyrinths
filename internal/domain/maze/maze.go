package maze

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// Maze хранит двумерную таблицу клеток, высоту и ширину.
type Maze struct {
	Cells  [][]cells.Cell
	Height int
	Width  int
}

// New возвращает инициализированный Maze.
func New(height, width int) Maze {
	maze := Maze{
		Cells:  make([][]cells.Cell, height),
		Height: height,
		Width:  width,
	}

	for y := 0; y < height; y++ {
		maze.Cells[y] = make([]cells.Cell, width)
	}

	return maze
}
