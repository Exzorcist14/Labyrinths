package sutils

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"

const (
	MissingX = -1
	MissingY = -1
)

// Predecessors - словарь {координаты - координаты, приведшие к ним}.
type Predecessors map[cells.Coordinates]cells.Coordinates

// NewPredecessors возвращает инициализированный Predecessors.
func NewPredecessors(height, width int) Predecessors {
	predecessors := make(Predecessors)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Изначально предшественник данных координат отсутствует.
			predecessors[cells.Coordinates{X: x, Y: y}] = cells.Coordinates{X: MissingX, Y: MissingY}
		}
	}

	return predecessors
}
