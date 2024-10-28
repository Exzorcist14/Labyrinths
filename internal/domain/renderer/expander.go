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

// RenderPath отображает лабиринт и путь в нём в готовую для визуализации строку и возвращает её.
func (r *expanderRenderer) RenderPath(mz maze.Maze, path []cells.Coordinates) string {
	return convertToString(expandMaze(overlayPath(mz, path)), r.palette)
}

// expandPalette возвращает расширенную визуализацией вспомогательных типов палитру.
func expandPalette(palette Palette) Palette {
	palette[transition] = "\U0001F532" // 🔲
	palette[Start] = "⭐"               // ⭐
	palette[End] = "🚩"                 // 🚩
	palette[Path] = "\U0001F7E9"       // 🟩

	return palette
}

// expandMaze возвращает расширенный лабиринт, в котором появляются стены.
func expandMaze(mz maze.Maze) maze.Maze {
	expandedMaze := maze.New(2*mz.Height-1, 2*mz.Width-1) // Между строками и столбцами появляются новые.

	for coords, cell := range mz.Cells {
		expandedCoords := cells.Coordinates{X: 2 * coords.X, Y: 2 * coords.Y} // Отображение координат исходного лабиринта в расширенный.

		expandedMaze.Cells[expandedCoords].Type = mz.Cells[coords].Type // Перенос типа клетки.

		for _, adjacentCoords := range cell.Transitions {
			expandedMaze.Cells[expandedCoords].Transitions = append(
				expandedMaze.Cells[expandedCoords].Transitions,
				cells.Coordinates{
					X: 2 * adjacentCoords.X, // Отображение координаты X клетки, куда есть переход исходного лабиринта в расширенный.
					Y: 2 * adjacentCoords.Y, // Отображение координаты Y клетки, куда есть переход исходного лабиринта в расширенный.
				},
			)
		}
	}

	return cutEdges(expandedMaze)
}

// cutEdges возвращает лабиринт, в котором между отображёнными в расширенный клетками появляются клетки типа transition.
func cutEdges(mz maze.Maze) maze.Maze {
	for coords, cell := range mz.Cells {
		if coords.X%2 == 0 && coords.Y%2 == 0 { // По формуле отображения лишь чётные координаты имеют смысловую нагрузку.
			for _, adjacentCoords := range cell.Transitions {
				edgeCoords := cells.Coordinates{
					X: (coords.X + adjacentCoords.X) / 2, // X получается по формуле середины отрезка.
					Y: (coords.Y + adjacentCoords.Y) / 2, // Y получается по формуле середины отрезка.
				}

				_, ok1 := pathParts[cell.Type]
				_, ok2 := pathParts[mz.Cells[adjacentCoords].Type]

				if ok1 && ok2 { // Если прорезаемое ребро принадлежит пути.
					mz.Cells[edgeCoords].Type = Path
				} else {
					mz.Cells[edgeCoords].Type = transition
				}

				mz.Cells[edgeCoords].Transitions = append(
					mz.Cells[edgeCoords].Transitions,
					cells.Coordinates{X: coords.X, Y: coords.Y},
				)

				mz.Cells[coords].Transitions = append(
					mz.Cells[coords].Transitions,
					edgeCoords,
				)
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
			mz.Cells[coords].Type = Start
		case len(path) - 1:
			mz.Cells[coords].Type = End
		default:
			mz.Cells[coords].Type = Path
		}
	}

	return mz
}

// convertToString возвращает готовый к отображению лабиринт в форме строки.
func convertToString(mz maze.Maze, palette Palette) string {
	var result strings.Builder

	for y := range mz.Height {
		for x := range mz.Width {
			result.WriteString(palette[mz.Cells[cells.Coordinates{X: x, Y: y}].Type])
		}

		result.WriteString("\n")
	}

	return result.String()
}
