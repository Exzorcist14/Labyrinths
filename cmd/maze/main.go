package main

import (
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application/session"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/renderers"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solvers"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/file"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/uis"
)

const pathToConfig = "./internal/infrastructure/files/config.json"

func main() {
	cfg := config.Config{}

	err := file.LoadData(pathToConfig, &cfg)
	if err != nil {
		os.Exit(1)
	}

	generator := generators.New(cfg.GeneratorType)
	solver := solvers.New(cfg.SolverType)

	renderer, err := renderers.New(cfg.RendererType)
	if err != nil {
		os.Exit(1)
	}

	ui := uis.New(cfg.UIType, renderer)

	s := session.New(generator, solver, ui)

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}
}
