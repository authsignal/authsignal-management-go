package authsignal

import "net/http"

type CreateNumberVerifyAuthenticatorConfigurationBody struct {
	IsActive    *bool   `json:"isActive,omitempty"`
	AppId       string  `json:"appId"`
	Environment *string `json:"environment,omitempty"`
}

type UpdateNumberVerifyAuthenticatorConfigurationBody struct {
	IsActive    NullableJsonInput[bool]   `json:"isActive,omitempty"`
	AppId       NullableJsonInput[string] `json:"appId,omitempty"`
	Environment NullableJsonInput[string] `json:"environment,omitempty"`
}

type NumberVerifyAuthenticatorConfiguration struct {
	AuthenticatorId    string  `json:"authenticatorId"`
	IsActive           bool    `json:"isActive"`
	VerificationMethod string  `json:"verificationMethod"`
	AppId              *string `json:"appId,omitempty"`
	Environment        *string `json:"environment,omitempty"`
}

func (c Client) CreateNumberVerifyAuthenticatorConfiguration(configuration CreateNumberVerifyAuthenticatorConfigurationBody) (*NumberVerifyAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[NumberVerifyAuthenticatorConfiguration](c, http.MethodPost, numberVerifyAuthenticatorConfigurationSlug, configuration)
}

func (c Client) GetNumberVerifyAuthenticatorConfiguration() (*NumberVerifyAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[NumberVerifyAuthenticatorConfiguration](c, http.MethodGet, numberVerifyAuthenticatorConfigurationSlug, nil)
}

func (c Client) UpdateNumberVerifyAuthenticatorConfiguration(configuration UpdateNumberVerifyAuthenticatorConfigurationBody) (*NumberVerifyAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[NumberVerifyAuthenticatorConfiguration](c, http.MethodPatch, numberVerifyAuthenticatorConfigurationSlug, configuration)
}

func (c Client) DeleteNumberVerifyAuthenticatorConfiguration() (*HttpStatusResponse, int, error) {
	return requestAuthenticatorConfiguration[HttpStatusResponse](c, http.MethodDelete, numberVerifyAuthenticatorConfigurationSlug, nil)
}
