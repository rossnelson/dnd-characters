package dndbeyond

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	characterServiceURL = "https://character-service.dndbeyond.com/character/v5/character"
)

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(token string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		token:      token,
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) FetchCharacter(characterID string) (*Character, error) {
	url := fmt.Sprintf("%s/%s", characterServiceURL, characterID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching character: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var response CharacterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("API returned success=false: %s", response.Message)
	}

	return &response.Data, nil
}

func (c *Client) FetchParty(characterID string) ([]PartyMember, error) {
	char, err := c.FetchCharacter(characterID)
	if err != nil {
		return nil, err
	}

	if char.Campaign == nil {
		return nil, fmt.Errorf("character is not in a campaign")
	}

	return char.Campaign.Characters, nil
}

func (c *Client) FetchAllPartyCharacters(characterID string) ([]*Character, error) {
	party, err := c.FetchParty(characterID)
	if err != nil {
		return nil, err
	}

	var characters []*Character
	for _, member := range party {
		char, err := c.FetchCharacter(fmt.Sprintf("%d", member.CharacterID))
		if err != nil {
			fmt.Printf("Warning: could not fetch %s (ID: %d): %v\n", member.CharacterName, member.CharacterID, err)
			continue
		}
		characters = append(characters, char)
	}

	return characters, nil
}
