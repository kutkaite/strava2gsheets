package gsheets

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"strava2gsheets/config"
)

const (
	authURL = "https://oauth2.googleapis.com/token"
	scopes  = "https://www.googleapis.com/auth/spreadsheets"
)

type Service struct {
	ctx           context.Context
	sheet         *sheets.SpreadsheetsService
	values        *sheets.SpreadsheetsValuesService
	spreadsheetID string
}

func NewService(ctx context.Context, cfg *config.GSheetsConfig) (Service, error) {
	jwtCfg := &jwt.Config{
		Email:        cfg.Email,
		PrivateKeyID: cfg.PrivateKeyID,
		PrivateKey:   []byte(cfg.PrivateKey),
		TokenURL:     authURL,
		Scopes:       []string{scopes},
	}
	client := jwtCfg.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return Service{}, fmt.Errorf("retrieve Google Sheets service: %v", err)
	}
	return Service{
		ctx:           ctx,
		sheet:         srv.Spreadsheets,
		values:        srv.Spreadsheets.Values,
		spreadsheetID: cfg.SpreadsheetID,
	}, nil
}
