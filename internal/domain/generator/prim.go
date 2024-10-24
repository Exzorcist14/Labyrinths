package generator

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// primGenerator - структура генератора по алгоритму Прима.
type primGenerator struct {
	border map[cells.Coordinates]struct{} // множество координат пограничных клеток
	maze   maze.Maze
}

// newPrimGenerator возвращает указатель на новый primGenerator.
func newPrimGenerator() *primGenerator {
	return &primGenerator{
		border: make(map[cells.Coordinates]struct{}),
	}
}

// Generate генерирует лабиринт заданной высоты и ширины.
func (g *primGenerator) Generate(height int, width int) (maze.Maze, error) {
	g.prepare(height, width)

	err := g.prim()
	if err != nil {
		return maze.Maze{}, fmt.Errorf("can`t generate using Prim`s algorithm: %w", err)
	}

	return g.maze, nil
}

// prim генерирует лабиринт по алгоритму Прима.
func (g *primGenerator) prim() error {
	// Суть алгоритма Прима (в текущей реализации):
	//
	// Изначально ни одна клетка не принадлежит лабиринту.
	//
	// Алгоритм:
	// 1) Выбирается случайная клетка лабиринта и становится "пограничной".
	// 2) Выбирается случайная пограничная клетка.
	// 3) Выбранная пограничная клетка становится частью имеющегося лабиринта.
	// 4) Смежные с ней клетки, не относящиеся к лабиринту, становятся пограничными, а она перестаёт ей быть.
	//
	// Действия 2, 3, 4 повторяются до тех пор, пока есть пограничные клетки.
	//
	// Получаемый лабиринт идеален.
	//
	// Итак, изначально все клетки лабиринта являются стенами ("лабиринт" пуст), которые будут заменяться проходами.
	// В дальнейшем под лабиринтом будет пониматься именно множество проходов.

	current, err := GetRandomCoords(g.maze.Height, g.maze.Width) // Выбираем случайную клетку.
	if err != nil {
		return fmt.Errorf("can`t get random coordinates: %w", err)
	}

	g.border[current] = struct{}{} //Клетка становится пограничной.

	for len(g.border) != 0 { // Пока есть пограничные клетки:
		current, err = GetRandomCoordsFrom(g.border) // Получаем случайную координаты пограничной клетки.
		if err != nil {
			return fmt.Errorf("can`get random available border coordinates: %w", err)
		}

		g.maze.Cells[current.Y][current.X].Type, err = GetRandomSignificantType() // Клетка по координатам получает тип.
		if err != nil {
			return fmt.Errorf("can`t get random significant type: %w", err)
		}

		err = g.linkToPassage(current) // Добавляем её в лабиринт.
		if err != nil {
			return fmt.Errorf("can`t link to maze: %w", err)
		}

		g.updateBorder(current) // Обновляем множество пограничных клеток.
	}

	return nil
}

// prepare подготавливает primGenerator для исполнения Generate.
func (g *primGenerator) prepare(height, width int) {
	g.maze = maze.New(height, width)

	g.maze.Cells = make([][]cells.Cell, height)
	for y := range height {
		g.maze.Cells[y] = make([]cells.Cell, width)
	}
}

// linkToPassage связывает клетку со случайным смежным проходом лабиринта.
func (g *primGenerator) linkToPassage(newPassage cells.Coordinates) error {
	previousPassage, err := g.getRandomAdjacentPassageCoords(newPassage)
	if err != nil {
		return fmt.Errorf("can`t get random adjacent passage current: %w", err)
	}

	if previousPassage.X != -1 {
		g.maze.Cells[newPassage.Y][newPassage.X].Transitions = append(
			g.maze.Cells[newPassage.Y][newPassage.X].Transitions,
			previousPassage,
		)

		g.maze.Cells[previousPassage.Y][previousPassage.X].Transitions = append(
			g.maze.Cells[previousPassage.Y][previousPassage.X].Transitions,
			newPassage,
		)
	}

	return nil
}

// updateBorder обновляет множество пограничных клеток, добавляя новые и удаляя текущую.
func (g *primGenerator) updateBorder(coords cells.Coordinates) {
	for i := range len(dx) {
		newCoords := cells.Coordinates{coords.X + dx[i], coords.Y + dy[i]}

		if IsInside(newCoords, g.maze.Height, g.maze.Width) &&
			g.maze.Cells[newCoords.Y][newCoords.X].Type == cells.Wall { //Если не является проходом.

			g.border[newCoords] = struct{}{}
		}
	}

	delete(g.border, coords)
}

// getRandomAdjacentPassageCoords возвращает координаты случайной смежной клетки, являющейся проходом.
func (g *primGenerator) getRandomAdjacentPassageCoords(coords cells.Coordinates) (cells.Coordinates, error) {
	adjacentPassagesCoords := []cells.Coordinates{}

	for i := range dx { // Заполняем слайс смежных проходов.
		adjacentCoords := cells.Coordinates{coords.X + dx[i], coords.Y + dy[i]}

		if IsInside(adjacentCoords, g.maze.Height, g.maze.Width) &&
			g.maze.Cells[adjacentCoords.Y][adjacentCoords.X].Type != cells.Wall {

			adjacentPassagesCoords = append(adjacentPassagesCoords, adjacentCoords)
		}
	}

	if len(adjacentPassagesCoords) != 0 {
		number, err := GetRandomInt(len(adjacentPassagesCoords)) // Получаем случайный номер прохода.
		if err != nil {
			return cells.Coordinates{}, fmt.Errorf("can`t generate random number of adjacent passage: %w", err)
		}

		return adjacentPassagesCoords[number], nil // Возвращаем проход по случайному номеру.
	} else {
		return cells.Coordinates{-1, -1}, nil
	}
}
