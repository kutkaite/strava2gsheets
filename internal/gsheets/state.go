package gsheets

import (
	"fmt"
	"log/slog"
	"time"

	"google.golang.org/api/sheets/v4"
)

const (
	layout       = "2006-01-02"
	fallbackDate = "2024-12-15"
)

func (g *Service) GetStartDateForNextRun() (time.Time, error) {
	slog.Info("fetching last ingested row..")
	allRows, err := g.listAllRows()
	if err != nil {
		return time.Time{}, fmt.Errorf("find latest row: %w", err)
	}
	startDate := getLastRunDate(allRows)

	parsedDate, err := time.Parse(layout, startDate)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse date %s as %s: %w", startDate, layout, err)
	}

	nextStartDate := parsedDate.Add(24 * time.Hour)
	slog.Info("successfully determined the date for the next run", slog.Time("nextRunStartDate", nextStartDate))
	return nextStartDate, nil
}

func getLastRunDate(latestRow *sheets.ValueRange) string {
	if len(latestRow.Values) == 0 || len(latestRow.Values[0]) == 0 {
		slog.Warn("no rows found in sheet, using fallback date", "fallbackDate", fallbackDate)
		return fallbackDate
	}

	// Get the value from the last row
	value := latestRow.Values[len(latestRow.Values)-1][0]
	str, ok := value.(string)
	if !ok {
		slog.Warn("value in last row is not a string, using fallback date", "fallbackDate", fallbackDate)
		return fallbackDate
	}
	return str
}
