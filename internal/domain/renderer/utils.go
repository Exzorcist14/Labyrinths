package renderer

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"

// Path - вспомогательная константа типа клетки, необходимая для пометки о принадлежности пути.
const Path cells.Type = -1

// Palette - словарь {тип клетки: строчной визуализация}.
type Palette map[cells.Type]string
