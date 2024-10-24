package main

import (
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application/session"
)

func main() {
	s, err := session.New()
	if err != nil {
		os.Exit(1)
	}

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}
}
