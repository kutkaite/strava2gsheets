package main

import (
	"log/slog"
	"os"

	"strava2gsheets/config"
	strava2gsheets "strava2gsheets/internal"
)

func main() {
	cfg, err := config.BuildConfig()
	if err != nil {
		slog.Error("build config", "err", err)
		os.Exit(1)
	}

	slog.Info("starting the application...")
	if err := strava2gsheets.Run(cfg); err != nil {
		slog.Error("run application", "err", err)
		os.Exit(1)
	}
}
