package gsheets

import (
	"fmt"
	"log/slog"
	"net/http"

	"google.golang.org/api/sheets/v4"
	"strava2gsheets/internal/strava"
)

func (g *Service) Write(activities []strava.ActivityAgg) error {
	allRows, err := g.listAllRows()
	if err != nil {
		return fmt.Errorf("find latest row: %w", err)
	}

	// Format the stats from the activities
	formattedStats := Format(activities)
	startRowIdx := len(allRows.Values) + 1
	noOfDaysToUpdate := len(formattedStats)
	noOfRowsToUpdate := startRowIdx + noOfDaysToUpdate

	// Check if the row range is empty and available for writing
	if err := g.rowRangeCheck(startRowIdx, noOfRowsToUpdate); err != nil {
		return fmt.Errorf("row range check: %w", err)
	}

	if err := g.writeRows(startRowIdx, noOfRowsToUpdate, formattedStats); err != nil {
		return err
	}
	return nil
}

func (g *Service) writeRows(rangeStart, rangeEnd int, stats []Stats) error {
	rowRange := getRange(rangeStart, rangeEnd)

	vals := make([][]interface{}, len(stats))
	for i, val := range stats {
		vals[i] = []interface{}{
			val.TrainingDateStr,
			val.TotalDur,
			val.TotalDistance,
		}
	}

	data := []*sheets.ValueRange{
		{
			MajorDimension: "ROWS",
			Range:          rowRange,
			Values:         vals,
		},
	}
	resp, err := g.values.BatchUpdate(g.spreadsheetID, &sheets.BatchUpdateValuesRequest{
		Data:             data,
		ValueInputOption: "USER_ENTERED",
	}).Do()
	if err != nil {
		return fmt.Errorf("batch update: %w", err)
	}

	if resp.HTTPStatusCode != http.StatusOK {
		return fmt.Errorf("couldn't update rows (status code: %d): %w", resp.HTTPStatusCode, err)
	}
	slog.Info("successfully updated the rows", "rowRange", rowRange)
	return nil
}
