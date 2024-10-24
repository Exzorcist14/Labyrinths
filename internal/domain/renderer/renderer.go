package renderer

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

type Renderer interface {
	Render(mz maze.Maze) string                               // Отображает лабиринт в готовую для визуализации строку.
	RenderPath(mz maze.Maze, path []cells.Coordinates) string // Отображает лабиринт и путь в готовую для визуализации строку.
}

// New как фабрика возвращает конкретную реализацию Renderer по строке, обозначающей желаемую реализацию, и палитре.
func New(rendererType string, palette Palette) Renderer {
	switch rendererType {
	case "expander":
		return newExpanderRenderer(palette)
	default:
		return newExpanderRenderer(palette)
	}
}
