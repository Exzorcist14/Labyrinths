package renderer

import (
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// transition - вспомогательная константа типа клетки, необходимая для пометки о том, что клетка является переходом.
const transition cells.Type = -2

// expanderRenderer - структура "расширяющего" рендера.
type expanderRenderer struct {
	palette Palette
}

// newExpanderRenderer возвращает указатель на новый expanderRenderer.
func newExpanderRenderer(palette Palette) *expanderRenderer {
	return &expanderRenderer{
		palette: expandPalette(palette),
	}
}

// Render отображает лабиринт в готовую для визуализации строку и возвращает её.
func (r *expanderRenderer) Render(mz maze.Maze) string {
	return convertToString(expandMaze(mz), r.palette)
}

// Render отображает лабиринт и путь в нём в готовую для визуализации строку и возвращает её.
func (r *expanderRenderer) RenderPath(mz maze.Maze, path []cells.Coordinates) string {
	return convertToString(expandMaze(overlayPath(mz, path)), r.palette)
}

// expandPalette возвращает расширенную визуализацией вспомогательных типов палитру.
func expandPalette(palette Palette) Palette {
	palette[transition] = "\U0001F532" // 🔲
	palette[Start] = "⭐"
	palette[End] = "🚩"
	palette[Path] = "\U0001F7E9" // 🟩

	return palette
}

// expandMaze возвращает расширенный лабиринт, в котором появляются стены.
func expandMaze(mz maze.Maze) maze.Maze {
	expandedMaze := maze.New(2*mz.Height-1, 2*mz.Width-1) // Между строками и столбаци появляются новые.

	for y, row := range mz.Cells {
		for x, cell := range row {
			expandedX := 2 * x // Отображение координаты X исходного лабиринта в расширенный.
			expandedY := 2 * y // Отображение координаты Y исходного лабиринта в расширенный.

			expandedMaze.Cells[expandedY][expandedX].Type = mz.Cells[y][x].Type // Перенос типа клетки.

			for _, adjacentCoords := range cell.Transitions {
				expandedMaze.Cells[expandedY][expandedX].Transitions = append(
					expandedMaze.Cells[expandedY][expandedX].Transitions,
					cells.Coordinates{
						X: 2 * adjacentCoords.X, // Отображение координаты X перехода исходного лабиринта в расширенный.
						Y: 2 * adjacentCoords.Y, // Отображение координаты Y перехода исходного лабиринта в расширенный.
					},
				)
			}
		}
	}

	return cutEdges(expandedMaze)
}

// cutEdges возвращает лабиринт, в котором между отображёнными в расширенный клетками появляются клетки типа transition.
func cutEdges(mz maze.Maze) maze.Maze {
	for y, row := range mz.Cells {
		for x, cell := range row {
			if x%2 == 0 && y%2 == 0 { // По формуле отображения лишь чётные координаты имеют смысловую нагрузку.
				for _, adjacentCoords := range cell.Transitions {
					edgeCoords := cells.Coordinates{
						X: (x + adjacentCoords.X) / 2, // X получается по формуле середины отрезка.
						Y: (y + adjacentCoords.Y) / 2, // Y получается по формуле середины отрезка.
					}

					_, ok1 := pathParts[cell.Type]
					_, ok2 := pathParts[mz.Cells[adjacentCoords.Y][adjacentCoords.X].Type]

					if ok1 && ok2 { // Если прорезаемое ребро принадлежит пути.
						mz.Cells[edgeCoords.Y][edgeCoords.X].Type = Path
					} else {
						mz.Cells[edgeCoords.Y][edgeCoords.X].Type = transition
					}

					mz.Cells[edgeCoords.Y][edgeCoords.X].Transitions = append(
						mz.Cells[edgeCoords.Y][edgeCoords.X].Transitions,
						cells.Coordinates{X: x, Y: y},
					)

					mz.Cells[y][x].Transitions = append(
						mz.Cells[y][x].Transitions,
						edgeCoords,
					)
				}
			}
		}
	}

	return mz
}

// overlayPath помечает клетки не расширенного лабиринта, принадлежащие path, как Path.
func overlayPath(mz maze.Maze, path []cells.Coordinates) maze.Maze {
	for i, coords := range path {
		switch i {
		case 0:
			mz.Cells[coords.Y][coords.X].Type = Start
		case len(path) - 1:
			mz.Cells[coords.Y][coords.X].Type = End
		default:
			mz.Cells[coords.Y][coords.X].Type = Path
		}
	}

	return mz
}

// convertToString возвращает готовый к отображению лабиринт в форме строки.
func convertToString(mz maze.Maze, palette Palette) string {
	var result strings.Builder

	for _, row := range mz.Cells {
		for _, cell := range row {
			result.WriteString(palette[cell.Type])
		}

		result.WriteString("\n")
	}

	return result.String()
}
