package authsignal

import "net/http"

type CreatePasskeyAuthenticatorConfigurationBody struct {
	IsActive                    *bool                  `json:"isActive,omitempty"`
	RelyingParty                string                 `json:"relyingParty"`
	ExpectedOrigins             []string               `json:"expectedOrigins"`
	PasskeyRegistrationHints    *ListJsonInput[string] `json:"passkeyRegistrationHints,omitempty"`
	UserVerificationRequirement *string                `json:"userVerificationRequirement,omitempty"`
	AuthenticatorAttachment     *string                `json:"authenticatorAttachment,omitempty"`
}

type UpdatePasskeyAuthenticatorConfigurationBody struct {
	IsActive                    NullableJsonInput[bool]   `json:"isActive,omitempty"`
	RelyingParty                NullableJsonInput[string] `json:"relyingParty,omitempty"`
	ExpectedOrigins             *ListJsonInput[string]    `json:"expectedOrigins,omitempty"`
	PasskeyRegistrationHints    *ListJsonInput[string]    `json:"passkeyRegistrationHints,omitempty"`
	UserVerificationRequirement NullableJsonInput[string] `json:"userVerificationRequirement,omitempty"`
	AuthenticatorAttachment     NullableJsonInput[string] `json:"authenticatorAttachment,omitempty"`
}

type PasskeyAuthenticatorConfiguration struct {
	AuthenticatorId             string    `json:"authenticatorId"`
	IsActive                    bool      `json:"isActive"`
	VerificationMethod          string    `json:"verificationMethod"`
	RelyingParty                *string   `json:"relyingParty,omitempty"`
	ExpectedOrigins             *[]string `json:"expectedOrigins,omitempty"`
	PasskeyRegistrationHints    *[]string `json:"passkeyRegistrationHints,omitempty"`
	UserVerificationRequirement *string   `json:"userVerificationRequirement,omitempty"`
	AuthenticatorAttachment     *string   `json:"authenticatorAttachment,omitempty"`
}

func (c Client) CreatePasskeyAuthenticatorConfiguration(configuration CreatePasskeyAuthenticatorConfigurationBody) (*PasskeyAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[PasskeyAuthenticatorConfiguration](c, http.MethodPost, passkeyAuthenticatorConfigurationSlug, configuration)
}

func (c Client) GetPasskeyAuthenticatorConfiguration() (*PasskeyAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[PasskeyAuthenticatorConfiguration](c, http.MethodGet, passkeyAuthenticatorConfigurationSlug, nil)
}

func (c Client) UpdatePasskeyAuthenticatorConfiguration(configuration UpdatePasskeyAuthenticatorConfigurationBody) (*PasskeyAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[PasskeyAuthenticatorConfiguration](c, http.MethodPatch, passkeyAuthenticatorConfigurationSlug, configuration)
}

func (c Client) DeletePasskeyAuthenticatorConfiguration() (*HttpStatusResponse, int, error) {
	return requestAuthenticatorConfiguration[HttpStatusResponse](c, http.MethodDelete, passkeyAuthenticatorConfigurationSlug, nil)
}
