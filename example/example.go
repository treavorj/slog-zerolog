package main

import (
	"fmt"
	"os"
	"time"

	"log/slog"

	slogzerolog "github.com/treavorj/slog-zerolog"
	"github.com/treavorj/zerolog"
)

func main() {
	zerologLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	logger := slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &zerologLogger}.NewZerologHandler())
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
