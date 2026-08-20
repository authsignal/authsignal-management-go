package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TenantSettings is the request shape for updating a tenant's settings. It covers the settings this
// client can write, not every field the endpoint accepts. A tenant always exists, so an update is
// partial: a setting left unset keeps its current value. Use SetValue to send one, including false.
type TenantSettings struct {
	HideSuccessScreenOnEnrollment NullableJsonInput[bool] `json:"hideSuccessScreenOnEnrollment,omitempty"`
}

// TenantResponse is the response shape for a tenant, covering the same settings TenantSettings can
// write. A nil field means the setting has never been set on the tenant, which is not the same as
// it being set to false.
type TenantResponse struct {
	HideSuccessScreenOnEnrollment *bool `json:"hideSuccessScreenOnEnrollment,omitempty"`
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
