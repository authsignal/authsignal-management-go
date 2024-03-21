package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ActionConfiguration struct {
	LastActionCreatedAt     string `json:"lastActionCreatedAt"`
	DefaultUserActionResult string `json:"defaultUserActionResult"`
	TenantId                string `json:"tenantId"`
	ActionCode              string `json:"actionCode"`
}

func (c Client) CreateActionConfiguration(actionConfiguration ActionConfiguration) (*ActionConfiguration, error) {
	createBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/action-configurations", c.Host), bytes.NewReader(createBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var createdActionConfiguration ActionConfiguration
	err = json.Unmarshal(body, &createdActionConfiguration)
	if err != nil {
		return nil, err
	}

	return &createdActionConfiguration, nil
}

func (c Client) GetActionConfiguration(actionCode string) (*ActionConfiguration, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionCode), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var actionConfiguration ActionConfiguration
	err = json.Unmarshal(body, &actionConfiguration)
	if err != nil {
		return nil, err
	}

	return &actionConfiguration, nil
}

func (c Client) UpdateActionConfiguration(actionConfiguration ActionConfiguration) (*ActionConfiguration, error) {
	updateBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionConfiguration.ActionCode), bytes.NewReader(updateBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var updatedActionConfiguration ActionConfiguration
	err = json.Unmarshal(body, &updatedActionConfiguration)
	if err != nil {
		return nil, err
	}

	return &updatedActionConfiguration, nil
}

func (c Client) DeleteActionConfiguration(actionCode string) (*HttpStatusResponse, error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionCode), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var httpStatusResponse HttpStatusResponse

	err = json.Unmarshal(body, &httpStatusResponse)
	if err != nil {
		return nil, err
	}

	return &httpStatusResponse, nil
}
