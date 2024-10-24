package ui

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/renderer"
)

type UserInterface interface {
	AskMazeDimensions() (height, width int, err error)                          // Спрашивает ширину и высоту.
	AskCoordinates(height, width int) (start, end cells.Coordinates, err error) // Спрашивает координаты start и end.
	DisplayMaze(mz maze.Maze)                                                   // Отображает лабиринт.
	DisplayMazeWithPath(mz maze.Maze, path []cells.Coordinates)                 // Отображает лабиринт и путь на нём.
}

// New как фабрика возвращает конкретную реализацию UserInterface по строке, обозначающей желаемую реализацию.
func New(consoleType, rendererType string, palette renderer.Palette) UserInterface {
	switch consoleType {
	case "console":
		return newConsole(rendererType, palette)
	default:
		return newConsole(rendererType, palette)
	}
}
