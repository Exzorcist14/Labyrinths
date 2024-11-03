package dijkstra

import (
	"math"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solvers/sutils"
)

// INF обозначает ненайденную дистанцию.
const INF = math.MaxInt

// Solver - структура Solver по алгоритму Дейкстры.
type Solver struct {
	dist         map[cells.Coordinates]cells.Type // Хранит для каждой вершины информацию об её оценке пути.
	heap         sutils.Heap                      // Куча минимумов, содержащая вершины и их оценку пути.
	predecessors sutils.Predecessors              // Хранит для каждой вершины информацию о её предшественниках.
}

// NewSolver возвращает указатель на инициализированный Solver.
func NewSolver() *Solver {
	ds := Solver{
		dist: make(map[cells.Coordinates]cells.Type),
		heap: sutils.New(),
	}

	return &ds
}

// Solve находит и возвращает путь от start до end в mz в виде []cells.Coordinates.
func (s *Solver) Solve(mz maze.Maze, start, end cells.Coordinates) []cells.Coordinates {
	s.prepare(mz.Height, mz.Width)

	s.dijkstra(mz, start, end)

	return sutils.RestorePath(start, end, s.predecessors)
}

// dijkstra находит кратчаший путь согласно алгоритму Дейкстры, записывая предшественника для каждой вершины.
// в predecessors для последующего восстановления пути.
func (s *Solver) dijkstra(mz maze.Maze, start, end cells.Coordinates) {
	// Суть алгоритма Дейкстры (в текущей реализации):
	//
	// Изначально оценка пути до каждой вершины равна INF.
	//
	// Алгоритм:
	// 1) Оценка пути до начальной вершины становится равной её весу, начало с оценкой кладётся в кучу минимумов.
	// 2) Достаётся вершина A с наименьшой оценкой пути из кучи.
	// 3) Если полученная вершина является end, алгоритм прерывает своё выполнение.
	// 4) Рассматривается каждая смежная с ней вершина, оценка до которой равна INF:
	//   3.1) Обновляется оценка её пути.
	//   3.2) Добавляется в кучу вместе с новой оценкой.
	//   3.3) Записывается координата вершины A (необходимо для восстановления пути по предшественникам).
	//
	// Пункты 2, 3, 4 повторяются, пока в куче существуют вершины, которые необходимо рассмотреть.
	weight := mz.Cells[start].Type

	s.dist[start] = weight
	s.heap.Push(sutils.Item{Vertex: start, Weight: weight})

	for s.heap.Len() != 0 {
		vertex1 := s.heap.Pop().Vertex // Получаем вершину с наименьшой оценкой пути из кучи.

		if vertex1 == end {
			break
		}

		for _, vertex2 := range mz.Cells[vertex1].Transitions { // Рассматриваем смежные вершины.
			if s.dist[vertex2] == INF { // Если оценка пути равна INF.
				s.dist[vertex2] = s.dist[vertex1] + mz.Cells[vertex2].Type         // Обновляем оценку пути.
				s.heap.Push(sutils.Item{Vertex: vertex2, Weight: s.dist[vertex2]}) // Добавляем в кучу.
				s.predecessors[vertex2] = vertex1                                  // Записываем предшественника для vertex2.
			}
		}
	}
}

// prepare подготавливает Solver для исполнения Solve.
func (s *Solver) prepare(height, width int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			s.dist[cells.Coordinates{X: x, Y: y}] = INF // Изначально оценка пути до каждой вершины равна INF.
		}
	}

	s.predecessors = sutils.NewPredecessors(height, width)
}
