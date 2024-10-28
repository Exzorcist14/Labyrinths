package renderer

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

type Renderer interface {
	Render(mz maze.Maze) string                               // Отображает лабиринт в готовую для визуализации строку.
	RenderPath(mz maze.Maze, path []cells.Coordinates) string // Отображает лабиринт и путь в готовую для визуализации строку.
}

// New как фабрика возвращает конкретную реализацию Renderer по строке, обозначающей желаемую реализацию, и палитре.
func New(rendererType string) (Renderer, error) {
	switch rendererType {
	case "expander":
		renderer, err := newExpanderRenderer()
		if err != nil {
			return nil, fmt.Errorf("can`t initialize expander renderer: %v", err)
		}

		return renderer, nil
	default:
		renderer, err := newExpanderRenderer()
		if err != nil {
			return nil, fmt.Errorf("can`t initialize expander renderer: %v", err)
		}

		return renderer, nil
	}
}
