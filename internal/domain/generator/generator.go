package generator

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
)

type Generator interface {
	Generate(height, width int) (maze.Maze, error)
}

// New как фабрика возвращает конкретную реализацию Generator по строке, обозначающей желаемую реализацию.
func New(generatorType string) Generator {
	switch generatorType {
	case "prim":
		return newPrimGenerator()
	case "wilson":
		return newWilsonGenerator()
	default:
		return newPrimGenerator()
	}
}
