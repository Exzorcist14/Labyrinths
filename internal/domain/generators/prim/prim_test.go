package prim_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators/prim"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
	"github.com/stretchr/testify/assert"
)

func TestPrimGeneratorGenerate(t *testing.T) {
	type args struct {
		height int
		width  int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "height & width: 1x1",
			args: args{
				height: 1,
				width:  1,
			},
		},
		{
			name: "height & width: 8x8",
			args: args{
				height: 8,
				width:  8,
			},
		},
		{
			name: "height & width: 64x64",
			args: args{
				height: 64,
				width:  64,
			},
		},
		{
			name: "height & width: 512x512",
			args: args{
				height: 512,
				width:  512,
			},
		},
		{
			name: "height & width: 1024x512",
			args: args{
				height: 1024,
				width:  512,
			},
		},
		{
			name: "height & width: 512x1024",
			args: args{
				height: 512,
				width:  1024,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := prim.NewGenerator()

			mz, err := g.Generate(tt.args.height, tt.args.width)

			assert.NoError(t, err)
			assert.True(t, isComponentOnlyOne(mz))
		})
	}
}

// isComponentOnlyOne проверяет, что в лабиринте одна компонента связности.
func isComponentOnlyOne(mz maze.Maze) bool {
	number := 0

	visited := make(map[cells.Coordinates]struct{})

	// Каждый поиск в глубину охватывает ровно одну компоненту связности.
	for cell := range mz.Cells {
		if _, ok := visited[cell]; !ok {
			dfs(cell, mz, visited)

			number++
		}

		if number > 1 {
			return false
		}
	}

	return true
}

// dfs отмечает пройденные вершины.
func dfs(current cells.Coordinates, mz maze.Maze, visited map[cells.Coordinates]struct{}) {
	visited[current] = struct{}{}

	for _, next := range mz.Cells[current].Transitions {
		if _, ok := visited[next]; !ok {
			dfs(next, mz, visited)
		}
	}
}
