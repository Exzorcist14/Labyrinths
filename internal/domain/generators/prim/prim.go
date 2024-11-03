package prim

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators/gutils"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// Generator - структура генератора по алгоритму Прима.
type Generator struct {
	border map[cells.Coordinates]struct{} // множество координат пограничных клеток
	mz     maze.Maze
}

// NewGenerator возвращает указатель на новый primGenerator.
func NewGenerator() *Generator {
	return &Generator{
		border: make(map[cells.Coordinates]struct{}),
	}
}

// Generate генерирует лабиринт заданной высоты и ширины.
func (g *Generator) Generate(height, width int) (maze.Maze, error) {
	g.prepare(height, width)

	err := g.prim()
	if err != nil {
		return maze.Maze{}, fmt.Errorf("can`t generate using Prim`s algorithm: %w", err)
	}

	return g.mz, nil
}

// prim генерирует лабиринт по алгоритму Прима.
func (g *Generator) prim() error {
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
	current, err := gutils.GetRandomCoords(g.mz.Height, g.mz.Width) // Выбираем случайную клетку.
	if err != nil {
		return fmt.Errorf("can`t get random coordinates: %w", err)
	}

	g.border[current] = struct{}{} // Клетка становится пограничной.

	for len(g.border) != 0 { // Пока есть пограничные клетки:
		current, err = gutils.GetRandomCoordsFrom(g.border) // Получаем случайную координаты пограничной клетки.
		if err != nil {
			return fmt.Errorf("can`get random available border coordinates: %w", err)
		}

		g.mz.Cells[current].Type, err = gutils.GetRandomSignificantType() // Клетка по координатам получает тип.
		if err != nil {
			return fmt.Errorf("can`t get random significant type: %w", err)
		}

		err = g.linkToPassage(current) // Добавляем её в лабиринт.
		if err != nil {
			return fmt.Errorf("can`t link to mz: %w", err)
		}

		g.updateBorder(current) // Обновляем множество пограничных клеток.
	}

	return nil
}

// prepare подготавливает Generator для исполнения Generate.
func (g *Generator) prepare(height, width int) {
	g.mz = maze.New(height, width)
}

// linkToPassage связывает клетку со случайным смежным проходом лабиринта.
func (g *Generator) linkToPassage(newPassage cells.Coordinates) error {
	previousPassage, err := g.getRandomAdjacentPassageCoords(newPassage)
	if err != nil {
		return fmt.Errorf("can`t get random adjacent passage current: %w", err)
	}

	if previousPassage.X != -1 {
		g.mz.Cells[newPassage].Transitions = append(
			g.mz.Cells[newPassage].Transitions,
			previousPassage,
		)

		g.mz.Cells[previousPassage].Transitions = append(
			g.mz.Cells[previousPassage].Transitions,
			newPassage,
		)
	}

	return nil
}

// updateBorder обновляет множество пограничных клеток, добавляя новые и удаляя текущую.
func (g *Generator) updateBorder(coords cells.Coordinates) {
	for i := 0; i < len(gutils.Dx); i++ {
		newCoords := cells.Coordinates{X: coords.X + gutils.Dx[i], Y: coords.Y + gutils.Dy[i]}

		if gutils.IsInside(newCoords, g.mz.Height, g.mz.Width) &&
			g.mz.Cells[newCoords].Type == cells.Wall { // Если не является проходом.
			g.border[newCoords] = struct{}{}
		}
	}

	delete(g.border, coords)
}

// getRandomAdjacentPassageCoords возвращает координаты случайной смежной клетки, являющейся проходом.
func (g *Generator) getRandomAdjacentPassageCoords(coords cells.Coordinates) (cells.Coordinates, error) {
	adjacentPassagesCoords := []cells.Coordinates{}

	for i := range gutils.Dx { // Заполняем слайс смежных проходов.
		adjacentCoords := cells.Coordinates{X: coords.X + gutils.Dx[i], Y: coords.Y + gutils.Dy[i]}

		if gutils.IsInside(adjacentCoords, g.mz.Height, g.mz.Width) &&
			g.mz.Cells[adjacentCoords].Type != cells.Wall {
			adjacentPassagesCoords = append(adjacentPassagesCoords, adjacentCoords)
		}
	}

	if len(adjacentPassagesCoords) != 0 {
		number, err := gutils.GetRandomInt(len(adjacentPassagesCoords)) // Получаем случайный номер прохода.
		if err != nil {
			return cells.Coordinates{}, fmt.Errorf("can`t generate random number of adjacent passage: %w", err)
		}

		return adjacentPassagesCoords[number], nil // Возвращаем проход по случайному номеру.
	}

	return cells.Coordinates{X: -1, Y: -1}, nil
}
