package solver

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solver/sutils"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solver/sutils/heaps"
)

// dfsSolver - структура решателя по модифицированному поиску в глубину (DFS).
type dfsSolver struct {
	visited      map[cells.Coordinates]struct{} // Хранит множество посещённых вершин.
	predecessors sutils.Predecessors            // Хранит для каждой вершины информацию о её предшественниках.
	mz           maze.Maze
}

// newDfsSolver возвращает указатель на инициализированный dfsSolver.
func newDfsSolver() *dfsSolver {
	return &dfsSolver{
		visited: make(map[cells.Coordinates]struct{}),
	}
}

// Solve находит и возвращает путь от start до end в mz в виде []cells.Coordinates.
func (s *dfsSolver) Solve(mz maze.Maze, start, end cells.Coordinates) []cells.Coordinates {
	s.prepare(mz)

	s.dfs(start, cells.Coordinates{X: sutils.MissingX, Y: sutils.MissingY}, end)

	return sutils.RestorePath(start, end, s.predecessors)
}

// dfs находит путь с помощью поиска в глубину, записывая предшественника для каждой вершины.
func (s *dfsSolver) dfs(current, previous, end cells.Coordinates) {
	// В текущей модификации поиск в глубину предпочитает при очередном выборе следующей вершины
	// сначала углубляться в более "дешёвые клетки", используя для этого локальную кучу минимумов.
	//
	// В некотором смысле, алгоритм пытается "здесь и сейчас" избежать обычного прохода cells.Pass и
	// вместо него сначала пойти в более освещённый cells.LightedPass; хотя в действительности
	// он не мыслит такими категориями, не ограничаясь только данными типами клеток, а рассуждая в плоскости их весов.
	s.predecessors[current] = previous
	s.visited[current] = struct{}{}

	if current == end {
		return
	}

	localHeap := heaps.New() // Локальная куча минимумов.

	for _, next := range s.mz.Cells[current].Transitions {
		if _, ok := s.visited[next]; !ok {
			localHeap.Push(heaps.Item{
				Vertex: next,
				Weight: s.mz.Cells[next].Type,
			})
		}
	}

	for localHeap.Len() != 0 { // Запускаем поиск в глубину для вершин в порядке увеличения их веса.
		s.dfs(localHeap.Pop().Vertex, current, end)
	}
}

// prepare подготавливает dfsSolver для исполнения Solve.
func (s *dfsSolver) prepare(mz maze.Maze) {
	s.mz = mz
	s.predecessors = sutils.NewPredecessors(s.mz.Height, s.mz.Width)
}
