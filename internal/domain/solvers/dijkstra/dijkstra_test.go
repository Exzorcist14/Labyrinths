package dijkstra_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solvers/dijkstra"
	"github.com/stretchr/testify/assert"
)

func TestDijkstraSolverSolve(t *testing.T) {
	type args struct {
		mz    maze.Maze
		start cells.Coordinates
		end   cells.Coordinates
	}

	tests := []struct {
		name     string
		args     args
		expected []cells.Coordinates
	}{
		{
			name: "path doesn`t exist",
			args: args{
				mz:    maze.New(3, 3),
				start: cells.Coordinates{X: 0, Y: 0},
				end:   cells.Coordinates{X: 2, Y: 2},
			},
			expected: []cells.Coordinates{},
		},
		{
			name: "path to self",
			args: args{
				mz:    maze.New(3, 3),
				start: cells.Coordinates{X: 0, Y: 0},
				end:   cells.Coordinates{X: 0, Y: 0},
			},
			expected: []cells.Coordinates{},
		},
		{
			name: "there is only one path",
			args: args{
				mz:    newOnePathMaze(),
				start: cells.Coordinates{X: 0, Y: 0},
				end:   cells.Coordinates{X: 2, Y: 2},
			},
			expected: []cells.Coordinates{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 1, Y: 2},
				{X: 2, Y: 2},
			},
		},
		{
			name: "path is the shortest of several",
			args: args{
				mz:    newSeveralPathMaze(),
				start: cells.Coordinates{X: 0, Y: 0},
				end:   cells.Coordinates{X: 2, Y: 2},
			},
			expected: []cells.Coordinates{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 1},
				{X: 2, Y: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := dijkstra.NewDijkstraSolver()

			path := s.Solve(tt.args.mz, tt.args.start, tt.args.end)

			assert.Equal(t, tt.expected, path)
		})
	}
}

func newOnePathMaze() maze.Maze {
	OnePathMaze := maze.New(3, 3)

	OnePathMaze.Cells[cells.Coordinates{X: 0, Y: 0}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 1},
		},
	}

	OnePathMaze.Cells[cells.Coordinates{X: 0, Y: 1}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 0},
			{X: 0, Y: 2}},
	}

	OnePathMaze.Cells[cells.Coordinates{X: 0, Y: 2}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 1},
			{X: 1, Y: 2}},
	}

	OnePathMaze.Cells[cells.Coordinates{X: 1, Y: 2}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 2},
			{X: 2, Y: 2}},
	}

	OnePathMaze.Cells[cells.Coordinates{X: 2, Y: 2}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 1, Y: 2},
		},
	}

	return OnePathMaze
}

func newSeveralPathMaze() maze.Maze {
	mz := maze.New(3, 3)

	mz.Cells[cells.Coordinates{X: 0, Y: 0}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 1},
			{X: 1, Y: 0},
		},
	}

	mz.Cells[cells.Coordinates{X: 2, Y: 2}] = &cells.Cell{
		Type: cells.LightedPass,
		Transitions: []cells.Coordinates{
			{X: 1, Y: 2},
			{X: 2, Y: 1},
		},
	}

	// Образуем первый путь, вес которого будет равен 8.
	// (cells.Pass + cells.LightedPass + cells.Pass + cells.Pass + cells.LightedPass)
	// 2           + 1                 + 2          +  2         + 1

	mz.Cells[cells.Coordinates{X: 0, Y: 1}] = &cells.Cell{
		Type: cells.LightedPass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 0},
			{X: 0, Y: 2}},
	}

	mz.Cells[cells.Coordinates{X: 0, Y: 2}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 1},
			{X: 1, Y: 2}},
	}

	mz.Cells[cells.Coordinates{X: 1, Y: 2}] = &cells.Cell{
		Type: cells.Pass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 2},
			{X: 2, Y: 2}},
	}

	// Образуем второй путь, вес которого будет равен 6.
	// (cells.Pass + cells.LightedPass + cells.Pass + cells.Pass + cells.LightedPass)
	// 2           + 1                 + 1          +  1         + 1

	mz.Cells[cells.Coordinates{X: 1, Y: 0}] = &cells.Cell{
		Type: cells.LightedPass,
		Transitions: []cells.Coordinates{
			{X: 0, Y: 0},
			{X: 2, Y: 0}},
	}

	mz.Cells[cells.Coordinates{X: 2, Y: 0}] = &cells.Cell{
		Type: cells.LightedPass,
		Transitions: []cells.Coordinates{
			{X: 1, Y: 0},
			{X: 2, Y: 1}},
	}

	mz.Cells[cells.Coordinates{X: 2, Y: 1}] = &cells.Cell{
		Type: cells.LightedPass,
		Transitions: []cells.Coordinates{
			{X: 2, Y: 0},
			{X: 2, Y: 2}},
	}

	return mz
}
