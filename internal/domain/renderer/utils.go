package renderer

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"

const (
	Start cells.Type = -10 // Вспомогательная константа типа клетки, помечающая начальную клетки.
	End   cells.Type = -20 // Вспомогательная константа типа клетки, помечающая конечную клетку.
	Path  cells.Type = -30 // Вспомогательная константа типа клетки, помечающая остальную часть пути.
)

// pathParts - множество типов клеток, обозначающих часть пути.
var pathParts = map[cells.Type]struct{}{
	Start: {},
	End:   {},
	Path:  {},
}

// Palette - словарь {тип клетки: строчной визуализация}.
type Palette map[cells.Type]string
