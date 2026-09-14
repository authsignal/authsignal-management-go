package authsignal

import "net/http"

type TwilioCredentialsCreate struct {
	AuthToken           string `json:"authToken"`
	MessagingServiceSid string `json:"messagingServiceSid"`
	AccountSid          string `json:"accountSid"`
}

type TwilioCredentialsUpdate struct {
	AuthToken           *string `json:"authToken,omitempty"`
	MessagingServiceSid *string `json:"messagingServiceSid,omitempty"`
	AccountSid          *string `json:"accountSid,omitempty"`
}

type TwilioCredentialsResponse struct {
	MessagingServiceSid *string `json:"messagingServiceSid,omitempty"`
	AccountSid          *string `json:"accountSid,omitempty"`
}

type BirdSmsCredentialsCreate struct {
	AccessKey              string  `json:"accessKey"`
	WorkspaceId            string  `json:"workspaceId"`
	ChannelId              *string `json:"channelId,omitempty"`
	NavigatorId            *string `json:"navigatorId,omitempty"`
	ProjectId              *string `json:"projectId,omitempty"`
	Locale                 *string `json:"locale,omitempty"`
	EnableMessageTemplates *bool   `json:"enableMessageTemplates,omitempty"`
}

type BirdSmsCredentialsUpdate struct {
	AccessKey              *string `json:"accessKey,omitempty"`
	WorkspaceId            *string `json:"workspaceId,omitempty"`
	ChannelId              *string `json:"channelId,omitempty"`
	NavigatorId            *string `json:"navigatorId,omitempty"`
	ProjectId              *string `json:"projectId,omitempty"`
	Locale                 *string `json:"locale,omitempty"`
	EnableMessageTemplates *bool   `json:"enableMessageTemplates,omitempty"`
}

type BirdSmsCredentialsResponse struct {
	WorkspaceId            *string `json:"workspaceId,omitempty"`
	ChannelId              *string `json:"channelId,omitempty"`
	NavigatorId            *string `json:"navigatorId,omitempty"`
	ProjectId              *string `json:"projectId,omitempty"`
	Locale                 *string `json:"locale,omitempty"`
	EnableMessageTemplates *bool   `json:"enableMessageTemplates,omitempty"`
}

type MessageMediaCredentialsCreate struct {
	ApiKey       string  `json:"apiKey"`
	ApiSecret    string  `json:"apiSecret"`
	SourceNumber *string `json:"sourceNumber,omitempty"`
}

type MessageMediaCredentialsUpdate struct {
	ApiKey       *string `json:"apiKey,omitempty"`
	ApiSecret    *string `json:"apiSecret,omitempty"`
	SourceNumber *string `json:"sourceNumber,omitempty"`
}

type MessageMediaCredentialsResponse struct {
	SourceNumber *string `json:"sourceNumber,omitempty"`
}

type ModicaGroupCredentialsCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ModicaGroupCredentialsUpdate struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

type ModicaGroupCredentialsResponse struct {
	Username *string `json:"username,omitempty"`
}

type TnzCredentialsCreate struct {
	ApiKey string `json:"apiKey"`
}

type TnzCredentialsUpdate struct {
	ApiKey *string `json:"apiKey,omitempty"`
}

type TnzCredentialsResponse struct{}

type CreateSmsAuthenticatorConfigurationBody struct {
	IsActive                         *bool                                        `json:"isActive,omitempty"`
	SmsProvider                      string                                       `json:"smsProvider"`
	SmsCountryCodes                  *ListJsonInput[string]                       `json:"smsCountryCodes,omitempty"`
	DefaultCountryCode               *string                                      `json:"defaultCountryCode,omitempty"`
	SubmissionRateLimitConfiguration *RateLimitConfiguration                      `json:"submissionRateLimitConfiguration,omitempty"`
	SendingRateLimitConfigurations   *ListJsonInput[RateLimitConfiguration]       `json:"sendingRateLimitConfigurations,omitempty"`
	PrefixRateLimitConfigurations    *ListJsonInput[PrefixRateLimitConfiguration] `json:"prefixRateLimitConfigurations,omitempty"`
	AllowedCustomSmsVariables        *ListJsonInput[string]                       `json:"allowedCustomSmsVariables,omitempty"`
	TwilioCredentials                *TwilioCredentialsCreate                     `json:"twilioCredentials,omitempty"`
	BirdSmsCredentials               *BirdSmsCredentialsCreate                    `json:"birdSmsCredentials,omitempty"`
	MessageMediaCredentials          *MessageMediaCredentialsCreate               `json:"messageMediaCredentials,omitempty"`
	ModicaGroupCredentials           *ModicaGroupCredentialsCreate                `json:"modicaGroupCredentials,omitempty"`
	TnzCredentials                   *TnzCredentialsCreate                        `json:"tnzCredentials,omitempty"`
}

type UpdateSmsAuthenticatorConfigurationBody struct {
	IsActive                         NullableJsonInput[bool]                          `json:"isActive,omitempty"`
	SmsProvider                      NullableJsonInput[string]                        `json:"smsProvider,omitempty"`
	SmsCountryCodes                  *ListJsonInput[string]                           `json:"smsCountryCodes,omitempty"`
	DefaultCountryCode               NullableJsonInput[string]                        `json:"defaultCountryCode,omitempty"`
	SubmissionRateLimitConfiguration NullableJsonInput[RateLimitConfiguration]        `json:"submissionRateLimitConfiguration,omitempty"`
	SendingRateLimitConfigurations   NullableJsonInput[[]RateLimitConfiguration]      `json:"sendingRateLimitConfigurations,omitempty"`
	PrefixRateLimitConfigurations    *ListJsonInput[PrefixRateLimitConfiguration]     `json:"prefixRateLimitConfigurations,omitempty"`
	AllowedCustomSmsVariables        *ListJsonInput[string]                           `json:"allowedCustomSmsVariables,omitempty"`
	TwilioCredentials                NullableJsonInput[TwilioCredentialsUpdate]       `json:"twilioCredentials,omitempty"`
	BirdSmsCredentials               NullableJsonInput[BirdSmsCredentialsUpdate]      `json:"birdSmsCredentials,omitempty"`
	MessageMediaCredentials          NullableJsonInput[MessageMediaCredentialsUpdate] `json:"messageMediaCredentials,omitempty"`
	ModicaGroupCredentials           NullableJsonInput[ModicaGroupCredentialsUpdate]  `json:"modicaGroupCredentials,omitempty"`
	TnzCredentials                   NullableJsonInput[TnzCredentialsUpdate]          `json:"tnzCredentials,omitempty"`
}

type SmsAuthenticatorConfiguration struct {
	AuthenticatorId                  string                           `json:"authenticatorId"`
	IsActive                         bool                             `json:"isActive"`
	VerificationMethod               string                           `json:"verificationMethod"`
	SmsProvider                      *string                          `json:"smsProvider,omitempty"`
	SmsCountryCodes                  *[]string                        `json:"smsCountryCodes,omitempty"`
	DefaultCountryCode               *string                          `json:"defaultCountryCode,omitempty"`
	SubmissionRateLimitConfiguration *RateLimitConfiguration          `json:"submissionRateLimitConfiguration,omitempty"`
	SendingRateLimitConfigurations   *[]RateLimitConfiguration        `json:"sendingRateLimitConfigurations,omitempty"`
	PrefixRateLimitConfigurations    *[]PrefixRateLimitConfiguration  `json:"prefixRateLimitConfigurations,omitempty"`
	AllowedCustomSmsVariables        *[]string                        `json:"allowedCustomSmsVariables,omitempty"`
	TwilioCredentials                *TwilioCredentialsResponse       `json:"twilioCredentials,omitempty"`
	BirdSmsCredentials               *BirdSmsCredentialsResponse      `json:"birdSmsCredentials,omitempty"`
	MessageMediaCredentials          *MessageMediaCredentialsResponse `json:"messageMediaCredentials,omitempty"`
	ModicaGroupCredentials           *ModicaGroupCredentialsResponse  `json:"modicaGroupCredentials,omitempty"`
	TnzCredentials                   *TnzCredentialsResponse          `json:"tnzCredentials,omitempty"`
}

func (c Client) CreateSmsAuthenticatorConfiguration(configuration CreateSmsAuthenticatorConfigurationBody) (*SmsAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[SmsAuthenticatorConfiguration](c, http.MethodPost, smsAuthenticatorConfigurationSlug, configuration)
}

func (c Client) GetSmsAuthenticatorConfiguration() (*SmsAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[SmsAuthenticatorConfiguration](c, http.MethodGet, smsAuthenticatorConfigurationSlug, nil)
}

func (c Client) UpdateSmsAuthenticatorConfiguration(configuration UpdateSmsAuthenticatorConfigurationBody) (*SmsAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[SmsAuthenticatorConfiguration](c, http.MethodPatch, smsAuthenticatorConfigurationSlug, configuration)
}

func (c Client) DeleteSmsAuthenticatorConfiguration() (*HttpStatusResponse, int, error) {
	return requestAuthenticatorConfiguration[HttpStatusResponse](c, http.MethodDelete, smsAuthenticatorConfigurationSlug, nil)
}
