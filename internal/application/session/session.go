package session

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generator"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/renderer"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solver"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/file"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/ui"
)

// Session хранит генератор, решатель и пользовательский интерфейс.
type Session struct {
	generator generator.Generator
	solver    solver.Solver
	ui        ui.UserInterface
}

// New возвращает инициализированную структуру Session.
func New() (*Session, error) {
	cfg := config.Config{}
	palette := renderer.Palette{}

	err := file.LoadData("./internal/infrastructure/files/config.json", &cfg)
	if err != nil {
		return nil, fmt.Errorf("can`t load config: %w", err)
	}

	err = file.LoadData("./internal/infrastructure/files/palette.json", &palette)
	if err != nil {
		return nil, fmt.Errorf("can`t load palette: %w", err)
	}

	return &Session{
		generator: generator.New(cfg.GeneratorType),
		solver:    solver.New(cfg.SolverType),
		ui:        ui.New(cfg.UIType, cfg.RendererType, palette),
	}, nil
}

// Run запускает проигрывание Session.
func (s *Session) Run() error {
	height, width, err := s.ui.AskMazeDimensions() // Спрашиваем размеры лабиринта.
	if err != nil {
		return fmt.Errorf("can`t ask maze dimensions: %w", err)
	}

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
