package renderer

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"

// Значения универсальных вспомогательных типов стоит кодировать двухзначными отрицательными числами,
// а значения вспомогательных типов конкретного рендерера - трёхзначными отрицательными.
const (
	Start cells.Type = -10 // Вспомогательный тип клетки, помечающий начальную клетки.
	End   cells.Type = -20 // Вспомогательный тип клетки, помечающий конечную клетку.
	Path  cells.Type = -30 // Вспомогательный тип клетки, помечающий остальную часть пути.
)

// pathParts - множество типов клеток, обозначающих часть пути.
var pathParts = map[cells.Type]struct{}{
	Start: {},
	End:   {},
	Path:  {},
}

// Palette - словарь {тип клетки: строчной визуализация}.
type Palette map[cells.Type]string
