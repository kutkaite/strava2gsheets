package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"strava2gsheets/config"
)

type Service struct {
	httpClient http.Client
	cfg        *config.StravaConfig
}

func New(cfg *config.StravaConfig) *Service {
	client := http.Client{}
	return &Service{cfg: cfg, httpClient: client}
}

func (s *Service) GetActivities(ctx context.Context, startTime, endTime time.Time) ([]ActivityAgg, error) {
	logger := slog.With(
		slog.Time("startTime", startTime),
		slog.Time("endTime", endTime),
	)
	logger.Info("fetching activities from Strava...")
	activities, err := s.fetchActivities(ctx, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("fetch strava activities")
	}
	if len(activities) == 0 {
		logger.Info("no new activities")
		return nil, nil
	}

	logger.Info("aggregating found activities", slog.Int("noOfActivities", len(activities)))
	aggStats := AggregatePerDay(activities)

	logger.Info("successfully fetched activities")
	return aggStats, nil
}

func (s *Service) fetchActivities(ctx context.Context, startDate, endDate time.Time) ([]ActivitiesResponse, error) {
	token, err := s.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("get access token: %w", err)
	}
	url := fmt.Sprintf(
		"https://www.strava.com/api/v3/athlete/activities?after=%d&before=%d", startDate.Unix(), endDate.Unix(),
	)
	req, err := buildRequest(ctx, http.MethodGet, url, nil, token)
	if err != nil {
		return nil, err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	activityResp, err := validateActivitiesResponse(resp)
	if err != nil {
		return nil, err
	}
	return activityResp, nil
}

func validateActivitiesResponse(resp *http.Response) ([]ActivitiesResponse, error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got %s from backend: %q (err: %v)", resp.Status, b, err)
	}

	var activityResp []ActivitiesResponse
	if err := json.Unmarshal(b, &activityResp); err != nil {
		return nil, fmt.Errorf("invalid activity response: %v", err)
	}
	return activityResp, nil
}

func buildRequest(
	ctx context.Context,
	httpMethod, url string,
	params io.Reader,
	auth *AuthResponse,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		httpMethod,
		url,
		params,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if auth != nil {
		req.Header.Add("Authorization", "Bearer "+auth.AccessToken)
	}

	return req, nil
}
