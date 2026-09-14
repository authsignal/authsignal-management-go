package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	smsAuthenticatorConfigurationSlug      = "sms"
	emailOtpAuthenticatorConfigurationSlug = "email-otp"
	passkeyAuthenticatorConfigurationSlug  = "passkey"
	pushAuthenticatorConfigurationSlug     = "push"
)

const (
	SmsOtpProviderBird         = "BIRD"
	SmsOtpProviderTwilio       = "TWILIO"
	SmsOtpProviderMessageMedia = "MESSAGE_MEDIA"
	SmsOtpProviderModicaGroup  = "MODICA_GROUP"
	SmsOtpProviderTnz          = "TNZ"
)

const (
	EmailOtpProviderBird     = "BIRD"
	EmailOtpProviderMailjet  = "MAILJET"
	EmailOtpProviderMailgun  = "MAILGUN"
	EmailOtpProviderMandrill = "MANDRILL"
	EmailOtpProviderSendgrid = "SENDGRID"
	EmailOtpProviderSmtp     = "SMTP"
)

const (
	PushConfigurationProviderDefault  = "DEFAULT"
	PushConfigurationProviderFirebase = "FIREBASE"
	PushConfigurationProviderWebhook  = "WEBHOOK"
)

const (
	AppAttestationFailureModeBlock            = "BLOCK"
	AppAttestationFailureModeAllowWithWarning = "ALLOW_WITH_WARNING"
)

const (
	UserVerificationRequirementPreferred = "preferred"
	UserVerificationRequirementRequired  = "required"
)

const (
	AuthenticatorAttachmentCrossPlatform = "cross-platform"
	AuthenticatorAttachmentPlatform      = "platform"
	AuthenticatorAttachmentAllSupported  = "all-supported"
)

const (
	PasskeyRegistrationHintSecurityKey  = "security-key"
	PasskeyRegistrationHintClientDevice = "client-device"
	PasskeyRegistrationHintHybrid       = "hybrid"
)

type RateLimitConfiguration struct {
	RateLimit       int64   `json:"rateLimit"`
	WindowInMinutes float64 `json:"windowInMinutes"`
}

type BlockSize struct {
	Country bool
	Numeric int64
}

func (b BlockSize) MarshalJSON() ([]byte, error) {
	if b.Country {
		return json.Marshal("country")
	}

	if b.Numeric <= 0 {
		return nil, fmt.Errorf("block size must be the string country or a positive number, got %d", b.Numeric)
	}

	return json.Marshal(b.Numeric)
}

func (b *BlockSize) UnmarshalJSON(data []byte) error {
	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		if asString != "country" {
			return fmt.Errorf("block size must be the string country or a positive number, got %s", data)
		}

		*b = BlockSize{Country: true}

		return nil
	}

	var asNumber int64
	if err := json.Unmarshal(data, &asNumber); err != nil {
		return fmt.Errorf("block size must be the string country or a positive number, got %s", data)
	}

	if asNumber <= 0 {
		return fmt.Errorf("block size must be the string country or a positive number, got %s", data)
	}

	*b = BlockSize{Numeric: asNumber}

	return nil
}

type PrefixRateLimitConfiguration struct {
	Enabled         *bool      `json:"enabled,omitempty"`
	BlockSize       *BlockSize `json:"blockSize,omitempty"`
	RateLimit       *int64     `json:"rateLimit,omitempty"`
	WindowInMinutes *float64   `json:"windowInMinutes,omitempty"`
	CountryCodes    *[]string  `json:"countryCodes,omitempty"`
}

func requestAuthenticatorConfiguration[T any](c Client, method string, slug string, requestBody any) (*T, int, error) {
	var reader io.Reader

	if requestBody != nil {
		encodedBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil, 0, err
		}

		reader = bytes.NewReader(encodedBody)
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s/authenticator-configurations/%s", c.Host, slug), reader)
	if err != nil {
		return nil, 0, err
	}

	if requestBody != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var decoded T

	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return nil, statusCode, err
	}

	return &decoded, statusCode, nil
}
