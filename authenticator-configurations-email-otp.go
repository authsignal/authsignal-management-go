package authsignal

import "net/http"

type BirdEmailCredentialsCreate struct {
	AccessKey   string  `json:"accessKey"`
	WorkspaceId string  `json:"workspaceId"`
	ChannelId   string  `json:"channelId"`
	ProjectId   string  `json:"projectId"`
	VersionId   *string `json:"versionId,omitempty"`
	Locale      *string `json:"locale,omitempty"`
	SenderEmail *string `json:"senderEmail,omitempty"`
	SenderName  *string `json:"senderName,omitempty"`
}

type BirdEmailCredentialsUpdate struct {
	AccessKey   *string `json:"accessKey,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
	ChannelId   *string `json:"channelId,omitempty"`
	ProjectId   *string `json:"projectId,omitempty"`
	VersionId   *string `json:"versionId,omitempty"`
	Locale      *string `json:"locale,omitempty"`
	SenderEmail *string `json:"senderEmail,omitempty"`
	SenderName  *string `json:"senderName,omitempty"`
}

type BirdEmailCredentialsResponse struct {
	WorkspaceId *string `json:"workspaceId,omitempty"`
	ChannelId   *string `json:"channelId,omitempty"`
	ProjectId   *string `json:"projectId,omitempty"`
	VersionId   *string `json:"versionId,omitempty"`
	Locale      *string `json:"locale,omitempty"`
	SenderEmail *string `json:"senderEmail,omitempty"`
	SenderName  *string `json:"senderName,omitempty"`
}

type MailjetEmailCredentialsCreate struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	TemplateId int64  `json:"templateId"`
}

type MailjetEmailCredentialsUpdate struct {
	PrivateKey *string `json:"privateKey,omitempty"`
	PublicKey  *string `json:"publicKey,omitempty"`
	TemplateId *int64  `json:"templateId,omitempty"`
}

type MailjetEmailCredentialsResponse struct {
	PublicKey  *string `json:"publicKey,omitempty"`
	TemplateId *int64  `json:"templateId,omitempty"`
}

type MailgunEmailCredentialsCreate struct {
	ApiKey       string  `json:"apiKey"`
	Url          string  `json:"url"`
	TemplateName string  `json:"templateName"`
	Domain       string  `json:"domain"`
	From         *string `json:"from,omitempty"`
}

type MailgunEmailCredentialsUpdate struct {
	ApiKey       *string `json:"apiKey,omitempty"`
	Url          *string `json:"url,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	Domain       *string `json:"domain,omitempty"`
	From         *string `json:"from,omitempty"`
}

type MailgunEmailCredentialsResponse struct {
	Url          *string `json:"url,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	Domain       *string `json:"domain,omitempty"`
	From         *string `json:"from,omitempty"`
}

type MandrillEmailCredentialsCreate struct {
	ApiKey       string `json:"apiKey"`
	TemplateName string `json:"templateName"`
}

type MandrillEmailCredentialsUpdate struct {
	ApiKey       *string `json:"apiKey,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
}

type MandrillEmailCredentialsResponse struct {
	TemplateName *string `json:"templateName,omitempty"`
}

type SendgridEmailCredentialsCreate struct {
	ApiKey     string  `json:"apiKey"`
	TemplateId string  `json:"templateId"`
	FromEmail  string  `json:"fromEmail"`
	FromName   *string `json:"fromName,omitempty"`
}

type SendgridEmailCredentialsUpdate struct {
	ApiKey     *string `json:"apiKey,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
	FromEmail  *string `json:"fromEmail,omitempty"`
	FromName   *string `json:"fromName,omitempty"`
}

type SendgridEmailCredentialsResponse struct {
	TemplateId *string `json:"templateId,omitempty"`
	FromEmail  *string `json:"fromEmail,omitempty"`
	FromName   *string `json:"fromName,omitempty"`
}

type SmtpEmailCredentialsCreate struct {
	Host        string  `json:"host"`
	Port        int64   `json:"port"`
	Secure      bool    `json:"secure"`
	User        string  `json:"user"`
	Password    string  `json:"password"`
	From        string  `json:"from"`
	FromName    *string `json:"fromName,omitempty"`
	ReplyTo     *string `json:"replyTo,omitempty"`
	ReplyToName *string `json:"replyToName,omitempty"`
}

type SmtpEmailCredentialsUpdate struct {
	Host        *string `json:"host,omitempty"`
	Port        *int64  `json:"port,omitempty"`
	Secure      *bool   `json:"secure,omitempty"`
	User        *string `json:"user,omitempty"`
	Password    *string `json:"password,omitempty"`
	From        *string `json:"from,omitempty"`
	FromName    *string `json:"fromName,omitempty"`
	ReplyTo     *string `json:"replyTo,omitempty"`
	ReplyToName *string `json:"replyToName,omitempty"`
}

type SmtpEmailCredentialsResponse struct {
	Host        *string `json:"host,omitempty"`
	Port        *int64  `json:"port,omitempty"`
	Secure      *bool   `json:"secure,omitempty"`
	User        *string `json:"user,omitempty"`
	From        *string `json:"from,omitempty"`
	FromName    *string `json:"fromName,omitempty"`
	ReplyTo     *string `json:"replyTo,omitempty"`
	ReplyToName *string `json:"replyToName,omitempty"`
}

type CreateEmailOtpAuthenticatorConfigurationBody struct {
	IsActive                         *bool                                  `json:"isActive,omitempty"`
	EmailProvider                    string                                 `json:"emailProvider"`
	SubmissionRateLimitConfiguration *RateLimitConfiguration                `json:"submissionRateLimitConfiguration,omitempty"`
	SendingRateLimitConfigurations   *ListJsonInput[RateLimitConfiguration] `json:"sendingRateLimitConfigurations,omitempty"`
	AllowedCustomEmailVariables      *ListJsonInput[string]                 `json:"allowedCustomEmailVariables,omitempty"`
	BirdEmailCredentials             *BirdEmailCredentialsCreate            `json:"birdEmailCredentials,omitempty"`
	MailjetEmailCredentials          *MailjetEmailCredentialsCreate         `json:"mailjetEmailCredentials,omitempty"`
	MailgunEmailCredentials          *MailgunEmailCredentialsCreate         `json:"mailgunEmailCredentials,omitempty"`
	MandrillEmailCredentials         *MandrillEmailCredentialsCreate        `json:"mandrillEmailCredentials,omitempty"`
	SendgridEmailCredentials         *SendgridEmailCredentialsCreate        `json:"sendgridEmailCredentials,omitempty"`
	SmtpEmailCredentials             *SmtpEmailCredentialsCreate            `json:"smtpEmailCredentials,omitempty"`
}

type UpdateEmailOtpAuthenticatorConfigurationBody struct {
	IsActive                         NullableJsonInput[bool]                           `json:"isActive,omitempty"`
	EmailProvider                    NullableJsonInput[string]                         `json:"emailProvider,omitempty"`
	SubmissionRateLimitConfiguration NullableJsonInput[RateLimitConfiguration]         `json:"submissionRateLimitConfiguration,omitempty"`
	SendingRateLimitConfigurations   NullableJsonInput[[]RateLimitConfiguration]       `json:"sendingRateLimitConfigurations,omitempty"`
	AllowedCustomEmailVariables      *ListJsonInput[string]                            `json:"allowedCustomEmailVariables,omitempty"`
	BirdEmailCredentials             NullableJsonInput[BirdEmailCredentialsUpdate]     `json:"birdEmailCredentials,omitempty"`
	MailjetEmailCredentials          NullableJsonInput[MailjetEmailCredentialsUpdate]  `json:"mailjetEmailCredentials,omitempty"`
	MailgunEmailCredentials          NullableJsonInput[MailgunEmailCredentialsUpdate]  `json:"mailgunEmailCredentials,omitempty"`
	MandrillEmailCredentials         NullableJsonInput[MandrillEmailCredentialsUpdate] `json:"mandrillEmailCredentials,omitempty"`
	SendgridEmailCredentials         NullableJsonInput[SendgridEmailCredentialsUpdate] `json:"sendgridEmailCredentials,omitempty"`
	SmtpEmailCredentials             NullableJsonInput[SmtpEmailCredentialsUpdate]     `json:"smtpEmailCredentials,omitempty"`
}

type EmailOtpAuthenticatorConfiguration struct {
	AuthenticatorId                  string                            `json:"authenticatorId"`
	IsActive                         bool                              `json:"isActive"`
	VerificationMethod               string                            `json:"verificationMethod"`
	EmailProvider                    *string                           `json:"emailProvider,omitempty"`
	SubmissionRateLimitConfiguration *RateLimitConfiguration           `json:"submissionRateLimitConfiguration,omitempty"`
	SendingRateLimitConfigurations   *[]RateLimitConfiguration         `json:"sendingRateLimitConfigurations,omitempty"`
	AllowedCustomEmailVariables      *[]string                         `json:"allowedCustomEmailVariables,omitempty"`
	BirdEmailCredentials             *BirdEmailCredentialsResponse     `json:"birdEmailCredentials,omitempty"`
	MailjetEmailCredentials          *MailjetEmailCredentialsResponse  `json:"mailjetEmailCredentials,omitempty"`
	MailgunEmailCredentials          *MailgunEmailCredentialsResponse  `json:"mailgunEmailCredentials,omitempty"`
	MandrillEmailCredentials         *MandrillEmailCredentialsResponse `json:"mandrillEmailCredentials,omitempty"`
	SendgridEmailCredentials         *SendgridEmailCredentialsResponse `json:"sendgridEmailCredentials,omitempty"`
	SmtpEmailCredentials             *SmtpEmailCredentialsResponse     `json:"smtpEmailCredentials,omitempty"`
}

func (c Client) CreateEmailOtpAuthenticatorConfiguration(configuration CreateEmailOtpAuthenticatorConfigurationBody) (*EmailOtpAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[EmailOtpAuthenticatorConfiguration](c, http.MethodPost, emailOtpAuthenticatorConfigurationSlug, configuration)
}

func (c Client) GetEmailOtpAuthenticatorConfiguration() (*EmailOtpAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[EmailOtpAuthenticatorConfiguration](c, http.MethodGet, emailOtpAuthenticatorConfigurationSlug, nil)
}

func (c Client) UpdateEmailOtpAuthenticatorConfiguration(configuration UpdateEmailOtpAuthenticatorConfigurationBody) (*EmailOtpAuthenticatorConfiguration, int, error) {
	return requestAuthenticatorConfiguration[EmailOtpAuthenticatorConfiguration](c, http.MethodPatch, emailOtpAuthenticatorConfigurationSlug, configuration)
}

func (c Client) DeleteEmailOtpAuthenticatorConfiguration() (*HttpStatusResponse, int, error) {
	return requestAuthenticatorConfiguration[HttpStatusResponse](c, http.MethodDelete, emailOtpAuthenticatorConfigurationSlug, nil)
}
