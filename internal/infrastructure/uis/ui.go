package uis

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

type renderer interface {
	Render(mz maze.Maze) string                               // Отображает лабиринт в готовую для визуализации строку.
	RenderPath(mz maze.Maze, path []cells.Coordinates) string // Отображает лабиринт и путь в готовую для визуализации строку.
}

type userInterface interface {
	AskMazeDimensions() (height, width int)                          // Спрашивает ширину и высоту.
	AskCoordinates(height, width int) (start, end cells.Coordinates) // Спрашивает координаты start и end.
	DisplayMaze(mz maze.Maze)                                        // Отображает лабиринт.
	DisplayMazeWithPath(mz maze.Maze, path []cells.Coordinates)      // Отображает лабиринт и путь на нём.
}

// New как фабрика возвращает конкретную реализацию userInterface по строке, обозначающей желаемую реализацию.
func New(consoleType string, r renderer) userInterface {
	switch consoleType {
	case "console":
		return newConsole(r)
	default:
		return newConsole(r)
	}
}
