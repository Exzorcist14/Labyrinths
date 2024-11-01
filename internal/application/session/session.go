package session

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

type generator interface {
	Generate(height, width int) (maze.Maze, error) // Возвращает сгенерированный лабиринт.
}

type solver interface {
	Solve(mz maze.Maze, start, end cells.Coordinates) []cells.Coordinates // Ищет путь от start до end.
}

type userInterface interface {
	AskMazeDimensions() (height, width int)                          // Спрашивает ширину и высоту.
	AskCoordinates(height, width int) (start, end cells.Coordinates) // Спрашивает координаты start и end.
	DisplayMaze(mz maze.Maze)                                        // Отображает лабиринт.
	DisplayMazeWithPath(mz maze.Maze, path []cells.Coordinates)      // Отображает лабиринт и путь на нём.
}

// Session хранит генератор, решатель и пользовательский интерфейс.
type Session struct {
	generator generator
	solver    solver
	ui        userInterface
}

// New возвращает инициализированную структуру Session.
func New(generator generator, solver solver, ui userInterface) *Session {
	return &Session{
		generator: generator,
		solver:    solver,
		ui:        ui,
	}
}

// Run запускает проигрывание Session.
func (s *Session) Run() error {
	height, width := s.ui.AskMazeDimensions() // Спрашиваем размеры лабиринта.

	mz, err := s.generator.Generate(height, width) // Генерируем лабиринт.
	if err != nil {
		return fmt.Errorf("can`t generate maze: %w", err)
	}

	start, end := s.ui.AskCoordinates(height, width) // Спрашиваем координаты начала и конца.

	path := s.solver.Solve(mz, start, end) // Ищем путь между началом и концом.

	s.ui.DisplayMaze(mz)               // Отображаем лабиринта на пользовательском интерфейсе.
	s.ui.DisplayMazeWithPath(mz, path) // Отображеем лабиринт и путь на пользовательском интерфейсе.

	return nil
}
