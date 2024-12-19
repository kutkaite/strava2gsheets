package config

import (
	"encoding/base64"
	"fmt"
	"os"
)

type Config struct {
	GSheetsService GSheetsConfig
	StravaService  StravaConfig
}

type GSheetsConfig struct {
	Email         string
	PrivateKeyID  string
	PrivateKey    string
	SpreadsheetID string
}

type StravaConfig struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

func BuildConfig() (Config, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(os.Getenv("GSHEETS_SERVICE_PRIVATE_KEY"))
	if err != nil {
		return Config{}, fmt.Errorf("decode base64 string: %w", err)
	}
	pem := fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s-----END PRIVATE KEY-----\n", string(decodedKey))

	cfg := Config{
		GSheetsService: GSheetsConfig{
			Email:         os.Getenv("GSHEETS_SERVICE_EMAIL"),
			PrivateKeyID:  os.Getenv("GSHEETS_SERVICE_PRIVATE_KEY_ID"),
			PrivateKey:    pem,
			SpreadsheetID: os.Getenv("GSHEETS_SERVICE_SPREADSHEET_ID"),
		}, StravaService: StravaConfig{
			ClientID:     os.Getenv("STRAVA_SERVICE_CLIENT_ID"),
			ClientSecret: os.Getenv("STRAVA_SERVICE_CLIENT_SECRET"),
			RefreshToken: os.Getenv("STRAVA_SERVICE_REFRESH_TOKEN"),
		},
	}

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.GSheetsService.Email == "" {
		return fmt.Errorf("missing required environment variable: GSHEETS_SERVICE_EMAIL")
	}
	if c.GSheetsService.PrivateKeyID == "" {
		return fmt.Errorf("missing required environment variable: GSHEETS_SERVICE_PRIVATE_KEY_ID")
	}
	if c.GSheetsService.PrivateKey == "" {
		return fmt.Errorf("missing required environment variable: GSHEETS_SERVICE_PRIVATE_KEY")
	}
	if c.GSheetsService.SpreadsheetID == "" {
		return fmt.Errorf("missing required environment variable: GSHEETS_SERVICE_SPREADSHEET_ID")
	}
	if c.StravaService.ClientID == "" {
		return fmt.Errorf("missing required environment variable: STRAVA_SERVICE_CLIENT_ID")
	}
	if c.StravaService.ClientSecret == "" {
		return fmt.Errorf("missing required environment variable: STRAVA_SERVICE_CLIENT_SECRET")
	}
	if c.StravaService.RefreshToken == "" {
		return fmt.Errorf("missing required environment variable: STRAVA_SERVICE_REFRESH_TOKEN")
	}
	return nil
}
