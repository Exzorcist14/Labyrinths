package wilson

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators/gutils"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// Generator - структура генератора по алгоритму Уилсона.
type Generator struct {
	unvisited map[cells.Coordinates]struct{}            // Множество непосещённых координат.
	wandering map[cells.Coordinates][]cells.Coordinates // Словарь координаты - координаты, к которым есть переход.
	mz        maze.Maze
}

// NewWilsonGenerator возвращает указатель на новый Generator.
func NewWilsonGenerator() *Generator {
	return &Generator{
		unvisited: make(map[cells.Coordinates]struct{}),
		wandering: make(map[cells.Coordinates][]cells.Coordinates),
	}
}

// Generate генерирует лабиринт заданной высоты и ширины.
func (g *Generator) Generate(height, width int) (maze.Maze, error) {
	g.prepare(height, width)

	err := g.wilson()
	if err != nil {
		return maze.Maze{}, fmt.Errorf("can`t generate using Wilson`s algorithm: %w", err)
	}

	return g.mz, nil
}

// wilson генерирует лабиринт по алгоритму Уилсона.
func (g *Generator) wilson() error {
	// Алгоритм Уилсона отличается тем, что генерирует несмещенную выборку из равномерного распределения
	// по всем лабиринтам, используя случайные блуждания с удалением петель,
	// хотя и имеет значительно более высокую временную сложность.
	//
	// Суть алгоритма Уилсона:
	//
	// Изначально ни одна клетка не принадлежит лабиринту.
	//
	// Алгоритм:
	// 1) Выбираются случайная клетка и становится частью лабиринта.
	// 2) Выбирается случайная клетка.
	// 3) Из этой клетки начинается блуждание:
	//   3.1) Выбирается случайная смежная клетка.
	//   3.2) Если выбранная клетка ещё не появлялась в блуждании (цикл отсутствует), добавляем её в блуждание.
	//        Иначе - сбрасываем блуждание, удаляя цикл.
	//   Действия 3.1 и 3.2 повторяются до тех пор, пока не будет достигнут лабиринт.
	// 4) Посещённые во время блуждания клетки становится частью лабиринта.
	//
	// Действия 2, 3, 4 повторяются до тех пор, пока существуют непосещённые клетки.
	//
	// Получаемый лабиринт идеален.
	//
	// Итак, изначально все клетки лабиринта являются стенами ("лабиринт" пуст), которые будут заменяться проходами,
	// В дальнейшем под лабиринтом будет пониматься именно множество имеющихся проходов.
	err := g.processRandomStartingCoords() // Обрабатываем первую случайную клетку.
	if err != nil {
		return fmt.Errorf("can`t processing random starting coordinates: %w", err)
	}

	for len(g.unvisited) > 0 { // Пока есть непосещённые клетки.
		err = g.randomlyWander() // Случайно блуждаем.
		if err != nil {
			return fmt.Errorf("can`t randomly wander: %w", err)
		}

		err = g.addWanderingToMaze() // Добавляем блуждание к лабиринту.
		if err != nil {
			return fmt.Errorf("can`t add wandering result to mz: %w", err)
		}
	}

	return nil
}

// prepare подготавливает Generator для исполнения Generate.
func (g *Generator) prepare(height, width int) {
	g.mz = maze.New(height, width)

	for coords := range g.mz.Cells {
		g.unvisited[coords] = struct{}{}
	}
}

// processRandomStartingCoords выбирает и обрабатывает случайные стартовые координаты.
func (g *Generator) processRandomStartingCoords() error {
	coords, err := gutils.GetRandomCoords(g.mz.Height, g.mz.Width) // Выбираем случайные координаты.
	if err != nil {
		return fmt.Errorf("can`t get random coordinates: %w", err)
	}

	g.mz.Cells[coords].Type, err = gutils.GetRandomSignificantType() // Клетка по координатам получает тип.
	if err != nil {
		return fmt.Errorf("can`t get random significant type: %w", err)
	}

	delete(g.unvisited, coords) // Удаляем координаты из списка непосещённых.

	return nil
}

// randomlyWander случайно блуждает, пока не встретит часть лабиринта.
func (g *Generator) randomlyWander() error {
	var (
		start    cells.Coordinates
		current  cells.Coordinates
		previous cells.Coordinates
		err      error
	)

	start, err = gutils.GetRandomCoordsFrom(g.unvisited) // Получаем координаты начала блуждания случайной непосещённый клетки.
	if err != nil {
		return fmt.Errorf("can`t get random coordinates: %w", err)
	}

	previous = start
	_, notMaze := g.unvisited[start]

	for notMaze { // Пока полученные координаты не относятся к лабиринту.
		current, err = gutils.GetRandomAdjacentCoords(previous, g.mz.Height, g.mz.Width)
		if err != nil {
			return fmt.Errorf("can`t get random coordinates: %w", err)
		}

		if _, isCycle := g.wandering[current]; !isCycle {
			g.addCoordsToWandering(current, previous) // Добавляем в блуждание.
			previous = current
		} else { // Если цикл обнаружен:
			previous = g.resetWandering(start) // Сбрасываем блуждание, откатывая previous до start.
		}

		_, notMaze = g.unvisited[current]
	}

	return nil
}

// addCoordsToWandering добавляет координаты в блуждание, обновляя слайсы переходов wandering генератора.
func (g *Generator) addCoordsToWandering(current, previous cells.Coordinates) {
	if current != previous {
		g.wandering[current] = append(g.wandering[current], previous)
		g.wandering[previous] = append(g.wandering[previous], current)
	}
}

// resetWandering сбрасывает блуждание, возвращая координаты, к которым надо сброситься.
func (g *Generator) resetWandering(start cells.Coordinates) (previous cells.Coordinates) {
	clear(g.wandering) // Очищаем блуждание.
	return start
}

// addWanderingToMaze добавляет результат блуждания к лабиринту.
func (g *Generator) addWanderingToMaze() error {
	var err error

	for coords, transitions := range g.wandering {
		g.mz.Cells[coords].Type, err = gutils.GetRandomSignificantType() // Клетка по координатам получает тип.
		if err != nil {
			return fmt.Errorf("can`t get random significant type: %w", err)
		}

		g.mz.Cells[coords].Transitions = append( // Добавляем переходы, полученные в результате блуждания.
			g.mz.Cells[coords].Transitions,
			transitions...,
		)

		delete(g.unvisited, coords) // Удаляем координаты из непосещённых.
	}

	return nil
}
