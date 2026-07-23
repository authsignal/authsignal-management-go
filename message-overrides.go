package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// MessageOverridesBody is the request shape for updating a tenant's pre-built UI message overrides.
// Overrides are keyed by locale, then by message override ID (publicId). Writes are a full
// replacement: any locale or override not present is removed. An empty map clears all overrides.
type MessageOverridesBody struct {
	MessageOverrides map[string]map[string]string `json:"messageOverrides"`
}

// MessageOverridesBodyResponse is the response shape for a tenant's pre-built UI message overrides.
type MessageOverridesBodyResponse struct {
	MessageOverrides map[string]map[string]string `json:"messageOverrides"`
}

type MessageOverridesCatalogScreen struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Family string `json:"family"`
}

type MessageOverridesCatalogPoint struct {
	PublicId            string            `json:"publicId"`
	Screen              string            `json:"screen"`
	Role                string            `json:"role"`
	Item                string            `json:"item"`
	Label               string            `json:"label"`
	Products            []string          `json:"products"`
	MaxLength           int64             `json:"maxLength"`
	AllowedPlaceholders []string          `json:"allowedPlaceholders"`
	AllowedTags         []string          `json:"allowedTags"`
	DefaultCopy         map[string]string `json:"defaultCopy"`
}

type MessageOverridesCatalogResponse struct {
	CatalogVersion int64                           `json:"catalogVersion"`
	Screens        []MessageOverridesCatalogScreen `json:"screens"`
	Points         []MessageOverridesCatalogPoint  `json:"points"`
}

func (c Client) GetMessageOverrides() (*MessageOverridesBodyResponse, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/message-overrides", c.Host), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var messageOverrides MessageOverridesBodyResponse
	err = json.Unmarshal(body, &messageOverrides)
	if err != nil {
		return nil, statusCode, err
	}

	return &messageOverrides, statusCode, nil
}

func (c Client) UpdateMessageOverrides(messageOverrides MessageOverridesBody) (*MessageOverridesBodyResponse, int, error) {
	updateBody, err := json.Marshal(messageOverrides)
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/message-overrides", c.Host), bytes.NewReader(updateBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var updatedMessageOverrides MessageOverridesBodyResponse
	err = json.Unmarshal(body, &updatedMessageOverrides)
	if err != nil {
		return nil, statusCode, err
	}

	return &updatedMessageOverrides, statusCode, nil
}

func (c Client) GetMessageOverridesCatalog() (*MessageOverridesCatalogResponse, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/message-overrides/catalog", c.Host), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var catalog MessageOverridesCatalogResponse
	err = json.Unmarshal(body, &catalog)
	if err != nil {
		return nil, statusCode, err
	}

	return &catalog, statusCode, nil
}
