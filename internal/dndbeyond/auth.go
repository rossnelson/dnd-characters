package dndbeyond

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	authServiceURL = "https://auth-service.dndbeyond.com/v1/cobalt-token"
)

type TokenResponse struct {
	Token string `json:"token"`
	TTL   int    `json:"ttl"`
}

type Auth struct {
	httpClient    *http.Client
	cobaltSession string
	token         string
	expiresAt     time.Time
}

func NewAuth(cobaltSession string) *Auth {
	return &Auth{
		httpClient:    &http.Client{Timeout: 30 * time.Second},
		cobaltSession: cobaltSession,
	}
}

func (a *Auth) GetToken() (string, error) {
	if a.token != "" && time.Now().Before(a.expiresAt) {
		return a.token, nil
	}

	return a.RefreshToken()
}

func (a *Auth) RefreshToken() (string, error) {
	req, err := http.NewRequest("POST", authServiceURL, nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://www.dndbeyond.com")
	req.Header.Set("Referer", "https://www.dndbeyond.com/")
	req.AddCookie(&http.Cookie{
		Name:  "CobaltSession",
		Value: a.cobaltSession,
	})

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("refreshing token: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("decoding token response: %w", err)
	}

	a.token = tokenResp.Token
	a.expiresAt = time.Now().Add(time.Duration(tokenResp.TTL-30) * time.Second)

	return a.token, nil
}
