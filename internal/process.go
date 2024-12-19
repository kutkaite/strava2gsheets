package strava2gsheets

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"strava2gsheets/config"
	"strava2gsheets/internal/gsheets"
	"strava2gsheets/internal/strava"
)

func Run(cfg config.Config) error {
	ctx := context.Background()

	slog.Info("creating Google Sheets client...")
	gSheetsSrv, err := gsheets.NewService(ctx, &cfg.GSheetsService)
	if err != nil {
		return fmt.Errorf("init new sheets gSheets service")
	}

	nextRun, err := gSheetsSrv.GetStartDateForNextRun()
	if err != nil {
		return fmt.Errorf("determine the start date for the next run: %w", err)
	}
	if nextRun.After(time.Now()) {
		slog.Info("no new activities", slog.Time("nextRun", nextRun))
		return nil
	}

	stravaSrv := strava.New(&cfg.StravaService)
	endDate := time.Now()
	activities, err := stravaSrv.GetActivities(ctx, nextRun, endDate)
	if err != nil {
		return fmt.Errorf("get aggregated activities: %w", err)
	}

	if err := gSheetsSrv.Write(activities); err != nil {
		return fmt.Errorf("write activities to a Google sheet: %w", err)
	}
	return nil
}
