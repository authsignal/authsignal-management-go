package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MessagingTemplates interface {
}

type ActionConfiguration struct {
	DefaultUserActionResult           NullableJsonInput[string]             `json:"defaultUserActionResult,omitempty"`
	ActionCode                        NullableJsonInput[string]             `json:"actionCode,omitempty"`
	MessagingTemplates                NullableJsonInput[MessagingTemplates] `json:"messagingTemplates,omitempty"`
	VerificationMethods               NullableJsonInput[[]string]           `json:"verificationMethods,omitempty"`
	PromptToEnrollVerificationMethods NullableJsonInput[[]string]           `json:"promptToEnrollVerificationMethods,omitempty"`
	DefaultVerificationMethod         NullableJsonInput[string]             `json:"defaultVerificationMethod,omitempty"`
}

type ActionConfigurationResponse struct {
	LastActionCreatedAt               string             `json:"lastActionCreatedAt"`
	DefaultUserActionResult           string             `json:"defaultUserActionResult"`
	TenantId                          string             `json:"tenantId"`
	ActionCode                        string             `json:"actionCode"`
	MessagingTemplates                MessagingTemplates `json:"messagingTemplates"`
	VerificationMethods               []string           `json:"verificationMethods"`
	PromptToEnrollVerificationMethods []string           `json:"promptToEnrollVerificationMethods"`
	DefaultVerificationMethod         string             `json:"defaultVerificationMethod"`
}

func (c Client) CreateActionConfiguration(actionConfiguration ActionConfiguration) (*ActionConfigurationResponse, int, error) {
	createBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/action-configurations", c.Host), bytes.NewReader(createBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var createdActionConfiguration ActionConfigurationResponse
	err = json.Unmarshal(body, &createdActionConfiguration)
	if err != nil {
		return nil, statusCode, err
	}

	return &createdActionConfiguration, statusCode, nil
}

func (c Client) GetActionConfiguration(actionCode string) (*ActionConfigurationResponse, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionCode), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var actionConfiguration ActionConfigurationResponse
	err = json.Unmarshal(body, &actionConfiguration)
	if err != nil {
		return nil, statusCode, err
	}

	return &actionConfiguration, statusCode, nil
}

func (c Client) UpdateActionConfiguration(actionCode string, actionConfiguration ActionConfiguration) (*ActionConfigurationResponse, int, error) {
	updateBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionCode), bytes.NewReader(updateBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var updatedActionConfiguration ActionConfigurationResponse
	err = json.Unmarshal(body, &updatedActionConfiguration)
	if err != nil {
		return nil, statusCode, err
	}

	return &updatedActionConfiguration, statusCode, nil
}

func (c Client) DeleteActionConfiguration(actionCode string) (*HttpStatusResponse, int, error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionCode), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var httpStatusResponse HttpStatusResponse

	err = json.Unmarshal(body, &httpStatusResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &httpStatusResponse, statusCode, nil
}
