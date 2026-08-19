package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TenantSettings is the request shape for updating a tenant's settings. It covers the settings the
// SDK manages rather than every setting the endpoint accepts, so more are added here as they are
// needed. A tenant always exists, so this is a partial update: fields left unset are omitted from
// the request and keep their current value. Use SetValue to send a setting, including a false one.
type TenantSettings struct {
	HideSuccessScreenOnEnrollment NullableJsonInput[bool] `json:"hideSuccessScreenOnEnrollment,omitempty"`
}

// TenantResponse is the response shape for a tenant, narrowed to the settings TenantSettings can
// write. Fields are pointers because the API omits a setting the tenant has never had set, which is
// not the same as one set to false.
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
