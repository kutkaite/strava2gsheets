package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	stravaAPIURL = "https://www.strava.com/api/v3/oauth/token"
	grantType    = "refresh_token"
)

func (s *Service) getAccessToken(ctx context.Context) (*AuthResponse, error) {
	params := url.Values{
		"client_id":     {s.cfg.ClientID},
		"grant_type":    {grantType},
		"refresh_token": {s.cfg.RefreshToken},
		"client_secret": {s.cfg.ClientSecret},
	}
	reqBody := strings.NewReader(params.Encode())

	req, err := buildRequest(ctx, http.MethodPost, stravaAPIURL, reqBody, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	authResp, err := validateAuthResponse(resp)
	if err != nil {
		return nil, err
	}
	return authResp, nil
}

func validateAuthResponse(resp *http.Response) (*AuthResponse, error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got %s from backend: %q (err: %v)", resp.Status, b, err)
	}

	var authResp AuthResponse
	if err := json.Unmarshal(b, &authResp); err != nil {
		return nil, fmt.Errorf("invalid response: %v", err)
	}
	return &authResp, nil
}
