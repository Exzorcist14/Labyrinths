package sutils

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// RestorePath восстанавливает путь из start, end и ps и возвращает его в виде []cells.Coordinates.
func RestorePath(start, end cells.Coordinates, ps Predecessors) []cells.Coordinates {
	var (
		invertedPath []cells.Coordinates
		path         []cells.Coordinates
	)

	current := end

	for current != ps[start] {
		invertedPath = append(invertedPath, current) // Проходясь по Predecessors, получается путь в обратном порядке.
		current = ps[current]
	}

	for i := len(invertedPath) - 1; i >= 0; i-- {
		path = append(path, invertedPath[i]) // Формируем прямой порядок.
	}

	return path
}
