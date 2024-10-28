package ui

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

type UserInterface interface {
	AskMazeDimensions() (height, width int)                          // Спрашивает ширину и высоту.
	AskCoordinates(height, width int) (start, end cells.Coordinates) // Спрашивает координаты start и end.
	DisplayMaze(mz maze.Maze)                                        // Отображает лабиринт.
	DisplayMazeWithPath(mz maze.Maze, path []cells.Coordinates)      // Отображает лабиринт и путь на нём.
}

// New как фабрика возвращает конкретную реализацию UserInterface по строке, обозначающей желаемую реализацию.
func New(consoleType, rendererType string) (UserInterface, error) {
	switch consoleType {
	case "console":
		c, err := newConsole(rendererType)
		if err != nil {
			return nil, fmt.Errorf("can`t initialize console: %w", err)
		}

		return c, nil
	default:
		c, err := newConsole(rendererType)
		if err != nil {
			return nil, fmt.Errorf("can`t initialize console: %w", err)
		}

		return c, nil
	}
}
