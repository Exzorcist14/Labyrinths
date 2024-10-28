package session

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generator"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solver"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/file"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/ui"
)

const pathToConfig = "./internal/infrastructure/files/config.json"

// Session хранит генератор, решатель и пользовательский интерфейс.
type Session struct {
	generator generator.Generator
	solver    solver.Solver
	ui        ui.UserInterface
}

// New возвращает инициализированную структуру Session.
func New() (*Session, error) {
	cfg := config.Config{}

	err := file.LoadData(pathToConfig, &cfg)
	if err != nil {
		return nil, fmt.Errorf("can`t load config: %w", err)
	}

	s := Session{
		generator: generator.New(cfg.GeneratorType),
		solver:    solver.New(cfg.SolverType),
	}

	s.ui, err = ui.New(cfg.UIType, cfg.RendererType)
	if err != nil {
		return nil, fmt.Errorf("can`t initialize ui: %w", err)
	}

	return &s, nil
}

// Run запускает проигрывание Session.
func (s *Session) Run() error {
	height, width := s.ui.AskMazeDimensions() // Спрашиваем размеры лабиринта.

	maze, err := s.generator.Generate(height, width) // Генерируем лабиринт.
	if err != nil {
		return fmt.Errorf("can`t generate maze: %w", err)
	}

	start, end := s.ui.AskCoordinates(height, width) // Спрашиваем координаты начала и конца.

	path := s.solver.Solve(maze, start, end) // Ищем путь между началом и концом.

	s.ui.DisplayMaze(maze)               // Отображаем лабиринта на пользовательском интерфейсе.
	s.ui.DisplayMazeWithPath(maze, path) // Отображеем лабиринт и путь на пользовательском интерфейсе.

	return nil
}
