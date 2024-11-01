package solvers

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

type solver interface {
	Solve(mz maze.Maze, begin, end cells.Coordinates) []cells.Coordinates
}

// New как фабрика возвращает конкретную реализацию Solver по строке, обозначающей желаемую реализацию.
func New(solverType string) solver {
	switch solverType {
	case "dijkstra":
		return newDijkstraSolver()
	case "mdfs":
		return newDfsSolver()
	default:
		return newDijkstraSolver()
	}
}