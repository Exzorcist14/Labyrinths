package main

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/application/session"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	session, err := session.New()
	if err != nil {
		logger.Error("failed to create session", slog.Any("error", err))
		os.Exit(1)
	}

	err = session.Run()
	if err != nil {
		logger.Error("failed to run session", slog.Any("error", err))
		os.Exit(1)
	}
}
