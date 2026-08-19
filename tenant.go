package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TenantSettings is the request shape for updating a tenant's settings. A tenant always exists, so
// this is a partial update: fields left unset are omitted from the request and keep their current
// value. Use SetValue to send a setting, including a false or zero one.
type TenantSettings struct {
	Name                          NullableJsonInput[string] `json:"name,omitempty"`
	TokenDurationInMinutes        NullableJsonInput[int64]  `json:"tokenDurationInMinutes,omitempty"`
	DisableRecoveryCodes          NullableJsonInput[bool]   `json:"disableRecoveryCodes,omitempty"`
	HideSuccessScreenOnEnrollment NullableJsonInput[bool]   `json:"hideSuccessScreenOnEnrollment,omitempty"`
}

// TenantResponse is the response shape for a tenant. It carries settings that are readable but not
// writable through this API, so it is wider than TenantSettings. Fields are pointers because the
// API omits a setting the tenant has never had set, which is not the same as one set to false.
type TenantResponse struct {
	TenantId                                  string  `json:"tenantId"`
	Name                                      *string `json:"name,omitempty"`
	CustomDomain                              *string `json:"customDomain,omitempty"`
	TokenDurationInMinutes                    *int64  `json:"tokenDurationInMinutes,omitempty"`
	AllowDisablingMfa                         *bool   `json:"allowDisablingMfa,omitempty"`
	HideAuthsignalLogo                        *bool   `json:"hideAuthsignalLogo,omitempty"`
	RedirectOnSessionExpiry                   *bool   `json:"redirectOnSessionExpiry,omitempty"`
	DisableRecoveryCodes                      *bool   `json:"disableRecoveryCodes,omitempty"`
	SkipRecoveryCodesOnProgrammaticEnrollment *bool   `json:"skipRecoveryCodesOnProgrammaticEnrollment,omitempty"`
	HideRecoveryCodesOnEnrollment             *bool   `json:"hideRecoveryCodesOnEnrollment,omitempty"`
	HideSuccessScreenOnEnrollment             *bool   `json:"hideSuccessScreenOnEnrollment,omitempty"`
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
