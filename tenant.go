package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticatorEventsWebhookConfig struct {
	Url                        *string `json:"url,omitempty"`
	IncludeCredentialPublicKey *bool   `json:"includeCredentialPublicKey,omitempty"`
}

type AuthenticatorEventsWebhookConfigResponse struct {
	Url                        *string `json:"url,omitempty"`
	IncludeCredentialPublicKey *bool   `json:"includeCredentialPublicKey,omitempty"`
}

type LogEventsWebhookConfig struct {
	EndpointUrl *string `json:"endpointUrl,omitempty"`
}

type LogEventsWebhookConfigResponse struct {
	EndpointUrl *string `json:"endpointUrl,omitempty"`
}

// TenantSettings is the request shape for updating a tenant's settings. It covers the settings this
// client can write, not every field the endpoint accepts. A tenant always exists, so an update is
// partial: a setting left unset keeps its current value. Use SetValue to send one, including false.
type TenantSettings struct {
	HideSuccessScreenOnEnrollment    NullableJsonInput[bool]                             `json:"hideSuccessScreenOnEnrollment,omitempty"`
	TokenDurationInMinutes           NullableJsonInput[int64]                            `json:"tokenDurationInMinutes,omitempty"`
	AuthenticatorEventsWebhookConfig NullableJsonInput[AuthenticatorEventsWebhookConfig] `json:"authenticatorEventsWebhookConfig,omitempty"`
	LogEventsWebhookConfig           NullableJsonInput[LogEventsWebhookConfig]           `json:"logEventsWebhookConfig,omitempty"`
	IpWhitelist                      *ListJsonInput[string]                              `json:"ipWhitelist,omitempty"`
}

// TenantResponse is the response shape for a tenant, covering the same settings TenantSettings can
// write. A nil field means the setting has never been set on the tenant, which is not the same as
// it being set to false.
type TenantResponse struct {
	HideSuccessScreenOnEnrollment    *bool                                     `json:"hideSuccessScreenOnEnrollment,omitempty"`
	TokenDurationInMinutes           *int64                                    `json:"tokenDurationInMinutes,omitempty"`
	AuthenticatorEventsWebhookConfig *AuthenticatorEventsWebhookConfigResponse `json:"authenticatorEventsWebhookConfig,omitempty"`
	LogEventsWebhookConfig           *LogEventsWebhookConfigResponse           `json:"logEventsWebhookConfig,omitempty"`
	IpWhitelist                      *[]string                                 `json:"ipWhitelist,omitempty"`
}

func (c Client) GetTenant() (*TenantResponse, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/tenant", c.Host), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var tenant TenantResponse
	err = json.Unmarshal(body, &tenant)
	if err != nil {
		return nil, statusCode, err
	}

	return &tenant, statusCode, nil
}

func (c Client) UpdateTenant(tenant TenantSettings) (*TenantResponse, int, error) {
	updateBody, err := json.Marshal(tenant)
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/tenant", c.Host), bytes.NewReader(updateBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var updatedTenant TenantResponse
	err = json.Unmarshal(body, &updatedTenant)
	if err != nil {
		return nil, statusCode, err
	}

	return &updatedTenant, statusCode, nil
}
