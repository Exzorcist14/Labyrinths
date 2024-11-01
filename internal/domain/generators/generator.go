package generators

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
)

type generator interface {
	Generate(height, width int) (maze.Maze, error)
}

// New как фабрика возвращает конкретную реализацию generators по строке, обозначающей желаемую реализацию.
func New(generatorType string) generator {
	switch generatorType {
	case "prim":
		return newPrimGenerator()
	case "wilson":
		return newWilsonGenerator()
	default:
		return newPrimGenerator()
	}
}
