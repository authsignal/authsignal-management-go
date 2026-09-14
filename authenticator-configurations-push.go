package authsignal

import "net/http"

type ApnsCredentialsCreate struct {
	TeamId     string `json:"teamId"`
	KeyId      string `json:"keyId"`
	BundleId   string `json:"bundleId"`
	PrivateKey string `json:"privateKey"`
	Sandbox    *bool  `json:"sandbox,omitempty"`
}

type ApnsCredentialsUpdate struct {
	TeamId     *string `json:"teamId,omitempty"`
	KeyId      *string `json:"keyId,omitempty"`
	BundleId   *string `json:"bundleId,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
	Sandbox    *bool   `json:"sandbox,omitempty"`
}

type ApnsCredentialsResponse struct {
	TeamId   *string `json:"teamId,omitempty"`
	KeyId    *string `json:"keyId,omitempty"`
	BundleId *string `json:"bundleId,omitempty"`
	Sandbox  *bool   `json:"sandbox,omitempty"`
}

type FcmCredentialsCreate struct {
	ServiceAccountKey string `json:"serviceAccountKey"`
}

type FcmCredentialsUpdate struct {
	ServiceAccountKey *string `json:"serviceAccountKey,omitempty"`
}

type FcmCredentialsResponse struct {
	ProjectId   *string `json:"projectId,omitempty"`
	ClientEmail *string `json:"clientEmail,omitempty"`
}

type CreatePushAuthenticatorConfigurationBody struct {
	IsActive                       *bool                                  `json:"isActive,omitempty"`
	PushProvider                   string                                 `json:"pushProvider"`
	WebhookUrl                     *string                                `json:"webhookUrl,omitempty"`
	CredentialLifetimeInMinutes    *int64                                 `json:"credentialLifetimeInMinutes,omitempty"`
	RequireAppAttestation          *bool                                  `json:"requireAppAttestation,omitempty"`
	AppAttestationFailureMode      *string                                `json:"appAttestationFailureMode,omitempty"`
	SendingRateLimitConfigurations *ListJsonInput[RateLimitConfiguration] `json:"sendingRateLimitConfigurations,omitempty"`
	ApnsCredentials                *ApnsCredentialsCreate                 `json:"apnsCredentials,omitempty"`
	FcmCredentials                 *FcmCredentialsCreate                  `json:"fcmCredentials,omitempty"`
}

type UpdatePushAuthenticatorConfigurationBody struct {
	IsActive                       NullableJsonInput[bool]                     `json:"isActive,omitempty"`
	PushProvider                   NullableJsonInput[string]                   `json:"pushProvider,omitempty"`
	WebhookUrl                     NullableJsonInput[string]                   `json:"webhookUrl,omitempty"`
	CredentialLifetimeInMinutes    NullableJsonInput[int64]                    `json:"credentialLifetimeInMinutes,omitempty"`
	RequireAppAttestation          NullableJsonInput[bool]                     `json:"requireAppAttestation,omitempty"`
	AppAttestationFailureMode      NullableJsonInput[string]                   `json:"appAttestationFailureMode,omitempty"`
	SendingRateLimitConfigurations NullableJsonInput[[]RateLimitConfiguration] `json:"sendingRateLimitConfigurations,omitempty"`
	ApnsCredentials                NullableJsonInput[ApnsCredentialsUpdate]    `json:"apnsCredentials,omitempty"`
	FcmCredentials                 NullableJsonInput[FcmCredentialsUpdate]     `json:"fcmCredentials,omitempty"`
}

type PushAuthenticatorConfiguration struct {
	AuthenticatorId                string                    `json:"authenticatorId"`
	IsActive                       bool                      `json:"isActive"`
	VerificationMethod             string                    `json:"verificationMethod"`
	PushProvider                   *string                   `json:"pushProvider,omitempty"`
	WebhookUrl                     *string                   `json:"webhookUrl,omitempty"`
	CredentialLifetimeInMinutes    *int64                    `json:"credentialLifetimeInMinutes,omitempty"`
	RequireAppAttestation          *bool                     `json:"requireAppAttestation,omitempty"`
	AppAttestationFailureMode      *string                   `json:"appAttestationFailureMode,omitempty"`
	SendingRateLimitConfigurations *[]RateLimitConfiguration `json:"sendingRateLimitConfigurations,omitempty"`
	ApnsCredentials                *ApnsCredentialsResponse  `json:"apnsCredentials,omitempty"`
	FcmCredentials                 *FcmCredentialsResponse   `json:"fcmCredentials,omitempty"`
}

func (c Client) CreatePushAuthenticatorConfiguration(configuration CreatePushAuthenticatorConfigurationBody) (*PushAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[PushAuthenticatorConfiguration](c, http.MethodPost, pushAuthenticatorConfigurationSlug, configuration)
}

func (c Client) GetPushAuthenticatorConfiguration() (*PushAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[PushAuthenticatorConfiguration](c, http.MethodGet, pushAuthenticatorConfigurationSlug, nil)
}

func (c Client) UpdatePushAuthenticatorConfiguration(configuration UpdatePushAuthenticatorConfigurationBody) (*PushAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[PushAuthenticatorConfiguration](c, http.MethodPatch, pushAuthenticatorConfigurationSlug, configuration)
}

func (c Client) DeletePushAuthenticatorConfiguration() (*HttpStatusResponse, int, error) {
	return requestAuthenticatorConfiguration[HttpStatusResponse](c, http.MethodDelete, pushAuthenticatorConfigurationSlug, nil)
}
