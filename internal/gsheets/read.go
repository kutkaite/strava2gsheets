package gsheets

import (
	"errors"
	"fmt"

	"google.golang.org/api/sheets/v4"
)

const (
	sheet       = "data"
	startColumn = "A"
	endColumn   = "C"
)

var ErrRowsNotEmpty = errors.New("gsheet rows are not empty")

func (g *Service) rowRangeCheck(rangeStart, rangeEnd int) error {
	resp, err := g.values.Get(g.spreadsheetID, getRange(rangeStart, rangeEnd)).Do()
	if err != nil {
		return fmt.Errorf("retrieve data from the sheet %s: %v", g.spreadsheetID, err)
	}

	if len(resp.Values) > 0 {
		return fmt.Errorf("%w: %s:%d-%d", ErrRowsNotEmpty, startColumn, rangeStart, rangeEnd)
	}
	return nil
}

func (g *Service) listAllRows() (*sheets.ValueRange, error) {
	rangeStr := fmt.Sprintf("%s!%s:%s", sheet, startColumn, endColumn)

	resp, err := g.values.Get(g.spreadsheetID, rangeStr).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve rows: %v", err)
	}
	if resp.Values == nil {
		return nil, errors.New("response values are empty")
	}
	return resp, nil
}

func getRange(rangeStart, rangeEnd int) string {
	return fmt.Sprintf("%s!%s%d:%s%d", sheet, startColumn, rangeStart, endColumn, rangeEnd)
}
