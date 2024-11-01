package generators

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators/prim"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators/wilson"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
)

type generator interface {
	Generate(height, width int) (maze.Maze, error)
}

// New как фабрика возвращает конкретную реализацию generators по строке, обозначающей желаемую реализацию.
func New(generatorType string) generator {
	switch generatorType {
	case "prim":
		return prim.NewPrimGenerator()
	case "wilson":
		return wilson.NewWilsonGenerator()
	default:
		return prim.NewPrimGenerator()
	}
}
