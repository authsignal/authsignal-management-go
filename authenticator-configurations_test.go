package authsignal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func ptr[T any](value T) *T {
	return &value
}

type recordedRequest struct {
	method string
	path   string
	body   string
}

func stubAuthsignal(t *testing.T, statusCode int, responseBody string) (Client, *recordedRequest) {
	t.Helper()

	recorded := &recordedRequest{}

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		recorded.method = request.Method
		recorded.path = request.URL.Path
		recorded.body = string(body)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(statusCode)
		_, _ = writer.Write([]byte(responseBody))
	}))

	t.Cleanup(server.Close)

	return NewClient(server.URL+"/v1/management", "tenant-1", "secret"), recorded
}

type authenticatorConfigurationCase struct {
	name             string
	call             func(Client) (any, int, error)
	responseBody     string
	expectedMethod   string
	expectedPath     string
	expectedBody     string
	expectedResponse any
}

func TestAuthenticatorConfigurationOperations(t *testing.T) {
	cases := []authenticatorConfigurationCase{
		{
			name: "create sms",
			call: func(client Client) (any, int, error) {
				return client.CreateSmsAuthenticatorConfiguration(CreateSmsAuthenticatorConfigurationBody{
					SmsProvider:                      SmsOtpProviderTwilio,
					SmsCountryCodes:                  SetList([]string{"AU", "NZ"}),
					DefaultCountryCode:               ptr("NZ"),
					SubmissionRateLimitConfiguration: &RateLimitConfiguration{RateLimit: 3, WindowInMinutes: 2.5},
					TwilioCredentials: &TwilioCredentialsCreate{
						AuthToken:           "token",
						MessagingServiceSid: "MG1",
						AccountSid:          "AC1",
					},
				})
			},
			responseBody:   `{"authenticatorId":"auth-1","isActive":true,"verificationMethod":"SMS","smsProvider":"TWILIO","twilioCredentials":{"messagingServiceSid":"MG1","accountSid":"AC1"}}`,
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/management/authenticator-configurations/sms",
			expectedBody:   `{"smsProvider":"TWILIO","smsCountryCodes":["AU","NZ"],"defaultCountryCode":"NZ","submissionRateLimitConfiguration":{"rateLimit":3,"windowInMinutes":2.5},"twilioCredentials":{"authToken":"token","messagingServiceSid":"MG1","accountSid":"AC1"}}`,
			expectedResponse: &SmsAuthenticatorConfiguration{
				AuthenticatorId:    "auth-1",
				IsActive:           true,
				VerificationMethod: "SMS",
				SmsProvider:        ptr(SmsOtpProviderTwilio),
				TwilioCredentials:  &TwilioCredentialsResponse{MessagingServiceSid: ptr("MG1"), AccountSid: ptr("AC1")},
			},
		},
		{
			name: "get sms",
			call: func(client Client) (any, int, error) {
				return client.GetSmsAuthenticatorConfiguration()
			},
			responseBody:   `{"authenticatorId":"auth-1","isActive":false,"verificationMethod":"SMS","smsProvider":"BIRD","prefixRateLimitConfigurations":[{"enabled":true,"blockSize":"country","rateLimit":5,"windowInMinutes":0.5}],"birdSmsCredentials":{"workspaceId":"workspace-1"}}`,
			expectedMethod: http.MethodGet,
			expectedPath:   "/v1/management/authenticator-configurations/sms",
			expectedResponse: &SmsAuthenticatorConfiguration{
				AuthenticatorId:    "auth-1",
				IsActive:           false,
				VerificationMethod: "SMS",
				SmsProvider:        ptr(SmsOtpProviderBird),
				PrefixRateLimitConfigurations: &[]PrefixRateLimitConfiguration{{
					Enabled:         ptr(true),
					BlockSize:       &BlockSize{Country: true},
					RateLimit:       ptr(int64(5)),
					WindowInMinutes: ptr(0.5),
				}},
				BirdSmsCredentials: &BirdSmsCredentialsResponse{WorkspaceId: ptr("workspace-1")},
			},
		},
		{
			name: "update sms clears the clearable fields and leaves the rest alone",
			call: func(client Client) (any, int, error) {
				return client.UpdateSmsAuthenticatorConfiguration(UpdateSmsAuthenticatorConfigurationBody{
					IsActive:                         SetValue(false),
					DefaultCountryCode:               SetNull(""),
					SubmissionRateLimitConfiguration: SetNull(RateLimitConfiguration{}),
					TwilioCredentials:                SetValue(TwilioCredentialsUpdate{AuthToken: ptr("rotated")}),
					BirdSmsCredentials:               SetNull(BirdSmsCredentialsUpdate{}),
				})
			},
			responseBody:   `{"authenticatorId":"auth-1","isActive":false,"verificationMethod":"SMS","smsProvider":"TWILIO"}`,
			expectedMethod: http.MethodPatch,
			expectedPath:   "/v1/management/authenticator-configurations/sms",
			expectedBody:   `{"isActive":false,"defaultCountryCode":null,"submissionRateLimitConfiguration":null,"twilioCredentials":{"authToken":"rotated"},"birdSmsCredentials":null}`,
			expectedResponse: &SmsAuthenticatorConfiguration{
				AuthenticatorId:    "auth-1",
				IsActive:           false,
				VerificationMethod: "SMS",
				SmsProvider:        ptr(SmsOtpProviderTwilio),
			},
		},
		{
			name: "delete sms",
			call: func(client Client) (any, int, error) {
				return client.DeleteSmsAuthenticatorConfiguration()
			},
			responseBody:     `{"success":true}`,
			expectedMethod:   http.MethodDelete,
			expectedPath:     "/v1/management/authenticator-configurations/sms",
			expectedResponse: &HttpStatusResponse{Success: true},
		},
		{
			name: "create email otp",
			call: func(client Client) (any, int, error) {
				return client.CreateEmailOtpAuthenticatorConfiguration(CreateEmailOtpAuthenticatorConfigurationBody{
					IsActive:                       ptr(false),
					EmailProvider:                  EmailOtpProviderSmtp,
					SendingRateLimitConfigurations: SetList([]RateLimitConfiguration{{RateLimit: 10, WindowInMinutes: 60}}),
					SmtpEmailCredentials: &SmtpEmailCredentialsCreate{
						Host:     "smtp.example.com",
						Port:     587,
						Secure:   false,
						User:     "user",
						Password: "super-secret",
						From:     "security@example.com",
					},
				})
			},
			responseBody:   `{"authenticatorId":"auth-2","isActive":false,"verificationMethod":"EMAIL_OTP","emailProvider":"SMTP","smtpEmailCredentials":{"host":"smtp.example.com","port":587,"secure":false,"user":"user","from":"security@example.com"}}`,
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/management/authenticator-configurations/email-otp",
			expectedBody:   `{"isActive":false,"emailProvider":"SMTP","sendingRateLimitConfigurations":[{"rateLimit":10,"windowInMinutes":60}],"smtpEmailCredentials":{"host":"smtp.example.com","port":587,"secure":false,"user":"user","password":"super-secret","from":"security@example.com"}}`,
			expectedResponse: &EmailOtpAuthenticatorConfiguration{
				AuthenticatorId:    "auth-2",
				IsActive:           false,
				VerificationMethod: "EMAIL_OTP",
				EmailProvider:      ptr(EmailOtpProviderSmtp),
				SmtpEmailCredentials: &SmtpEmailCredentialsResponse{
					Host:   ptr("smtp.example.com"),
					Port:   ptr(int64(587)),
					Secure: ptr(false),
					User:   ptr("user"),
					From:   ptr("security@example.com"),
				},
			},
		},
		{
			name: "get email otp",
			call: func(client Client) (any, int, error) {
				return client.GetEmailOtpAuthenticatorConfiguration()
			},
			responseBody:   `{"authenticatorId":"auth-2","isActive":true,"verificationMethod":"EMAIL_OTP","emailProvider":"MAILJET","allowedCustomEmailVariables":["firstName"],"mailjetEmailCredentials":{"publicKey":"public","templateId":42}}`,
			expectedMethod: http.MethodGet,
			expectedPath:   "/v1/management/authenticator-configurations/email-otp",
			expectedResponse: &EmailOtpAuthenticatorConfiguration{
				AuthenticatorId:             "auth-2",
				IsActive:                    true,
				VerificationMethod:          "EMAIL_OTP",
				EmailProvider:               ptr(EmailOtpProviderMailjet),
				AllowedCustomEmailVariables: &[]string{"firstName"},
				MailjetEmailCredentials:     &MailjetEmailCredentialsResponse{PublicKey: ptr("public"), TemplateId: ptr(int64(42))},
			},
		},
		{
			name: "update email otp",
			call: func(client Client) (any, int, error) {
				return client.UpdateEmailOtpAuthenticatorConfiguration(UpdateEmailOtpAuthenticatorConfigurationBody{
					EmailProvider:                  SetValue(EmailOtpProviderSendgrid),
					SendingRateLimitConfigurations: SetNull([]RateLimitConfiguration{}),
					AllowedCustomEmailVariables:    SetList([]string{}),
					SendgridEmailCredentials:       SetValue(SendgridEmailCredentialsUpdate{ApiKey: ptr("key"), TemplateId: ptr("d-1"), FromEmail: ptr("a@b.com")}),
				})
			},
			responseBody:   `{"authenticatorId":"auth-2","isActive":true,"verificationMethod":"EMAIL_OTP","emailProvider":"SENDGRID"}`,
			expectedMethod: http.MethodPatch,
			expectedPath:   "/v1/management/authenticator-configurations/email-otp",
			expectedBody:   `{"emailProvider":"SENDGRID","sendingRateLimitConfigurations":null,"allowedCustomEmailVariables":[],"sendgridEmailCredentials":{"apiKey":"key","templateId":"d-1","fromEmail":"a@b.com"}}`,
			expectedResponse: &EmailOtpAuthenticatorConfiguration{
				AuthenticatorId:    "auth-2",
				IsActive:           true,
				VerificationMethod: "EMAIL_OTP",
				EmailProvider:      ptr(EmailOtpProviderSendgrid),
			},
		},
		{
			name: "delete email otp",
			call: func(client Client) (any, int, error) {
				return client.DeleteEmailOtpAuthenticatorConfiguration()
			},
			responseBody:     `{"success":true}`,
			expectedMethod:   http.MethodDelete,
			expectedPath:     "/v1/management/authenticator-configurations/email-otp",
			expectedResponse: &HttpStatusResponse{Success: true},
		},
		{
			name: "create passkey",
			call: func(client Client) (any, int, error) {
				return client.CreatePasskeyAuthenticatorConfiguration(CreatePasskeyAuthenticatorConfigurationBody{
					RelyingParty:                "example.com",
					ExpectedOrigins:             []string{"https://example.com"},
					PasskeyRegistrationHints:    SetList([]string{PasskeyRegistrationHintClientDevice}),
					UserVerificationRequirement: ptr(UserVerificationRequirementRequired),
					AuthenticatorAttachment:     ptr(AuthenticatorAttachmentPlatform),
				})
			},
			responseBody:   `{"authenticatorId":"auth-3","isActive":true,"verificationMethod":"PASSKEY","relyingParty":"example.com","expectedOrigins":["https://example.com"],"passkeyRegistrationHints":["client-device"],"userVerificationRequirement":"required","authenticatorAttachment":"platform"}`,
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/management/authenticator-configurations/passkey",
			expectedBody:   `{"relyingParty":"example.com","expectedOrigins":["https://example.com"],"passkeyRegistrationHints":["client-device"],"userVerificationRequirement":"required","authenticatorAttachment":"platform"}`,
			expectedResponse: &PasskeyAuthenticatorConfiguration{
				AuthenticatorId:             "auth-3",
				IsActive:                    true,
				VerificationMethod:          "PASSKEY",
				RelyingParty:                ptr("example.com"),
				ExpectedOrigins:             &[]string{"https://example.com"},
				PasskeyRegistrationHints:    &[]string{PasskeyRegistrationHintClientDevice},
				UserVerificationRequirement: ptr(UserVerificationRequirementRequired),
				AuthenticatorAttachment:     ptr(AuthenticatorAttachmentPlatform),
			},
		},
		{
			name: "get passkey",
			call: func(client Client) (any, int, error) {
				return client.GetPasskeyAuthenticatorConfiguration()
			},
			responseBody:   `{"authenticatorId":"auth-3","isActive":true,"verificationMethod":"PASSKEY","relyingParty":"example.com","expectedOrigins":["android:apk-key-hash:abc"]}`,
			expectedMethod: http.MethodGet,
			expectedPath:   "/v1/management/authenticator-configurations/passkey",
			expectedResponse: &PasskeyAuthenticatorConfiguration{
				AuthenticatorId:    "auth-3",
				IsActive:           true,
				VerificationMethod: "PASSKEY",
				RelyingParty:       ptr("example.com"),
				ExpectedOrigins:    &[]string{"android:apk-key-hash:abc"},
			},
		},
		{
			name: "update passkey",
			call: func(client Client) (any, int, error) {
				return client.UpdatePasskeyAuthenticatorConfiguration(UpdatePasskeyAuthenticatorConfigurationBody{
					ExpectedOrigins:             SetList([]string{"https://app.example.com"}),
					UserVerificationRequirement: SetNull(""),
					AuthenticatorAttachment:     SetNull(""),
				})
			},
			responseBody:   `{"authenticatorId":"auth-3","isActive":true,"verificationMethod":"PASSKEY","relyingParty":"example.com","expectedOrigins":["https://app.example.com"]}`,
			expectedMethod: http.MethodPatch,
			expectedPath:   "/v1/management/authenticator-configurations/passkey",
			expectedBody:   `{"expectedOrigins":["https://app.example.com"],"userVerificationRequirement":null,"authenticatorAttachment":null}`,
			expectedResponse: &PasskeyAuthenticatorConfiguration{
				AuthenticatorId:    "auth-3",
				IsActive:           true,
				VerificationMethod: "PASSKEY",
				RelyingParty:       ptr("example.com"),
				ExpectedOrigins:    &[]string{"https://app.example.com"},
			},
		},
		{
			name: "delete passkey",
			call: func(client Client) (any, int, error) {
				return client.DeletePasskeyAuthenticatorConfiguration()
			},
			responseBody:     `{"success":true}`,
			expectedMethod:   http.MethodDelete,
			expectedPath:     "/v1/management/authenticator-configurations/passkey",
			expectedResponse: &HttpStatusResponse{Success: true},
		},
		{
			name: "create push",
			call: func(client Client) (any, int, error) {
				return client.CreatePushAuthenticatorConfiguration(CreatePushAuthenticatorConfigurationBody{
					PushProvider:                PushConfigurationProviderFirebase,
					CredentialLifetimeInMinutes: ptr(int64(15)),
					RequireAppAttestation:       ptr(true),
					AppAttestationFailureMode:   ptr(AppAttestationFailureModeBlock),
					FcmCredentials:              &FcmCredentialsCreate{ServiceAccountKey: `{"project_id":"project-1"}`},
				})
			},
			responseBody:   `{"authenticatorId":"auth-4","isActive":true,"verificationMethod":"PUSH","pushProvider":"FIREBASE","credentialLifetimeInMinutes":15,"requireAppAttestation":true,"appAttestationFailureMode":"BLOCK","fcmCredentials":{"projectId":"project-1","clientEmail":"sa@project-1.iam.gserviceaccount.com"}}`,
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/management/authenticator-configurations/push",
			expectedBody:   `{"pushProvider":"FIREBASE","credentialLifetimeInMinutes":15,"requireAppAttestation":true,"appAttestationFailureMode":"BLOCK","fcmCredentials":{"serviceAccountKey":"{\"project_id\":\"project-1\"}"}}`,
			expectedResponse: &PushAuthenticatorConfiguration{
				AuthenticatorId:             "auth-4",
				IsActive:                    true,
				VerificationMethod:          "PUSH",
				PushProvider:                ptr(PushConfigurationProviderFirebase),
				CredentialLifetimeInMinutes: ptr(int64(15)),
				RequireAppAttestation:       ptr(true),
				AppAttestationFailureMode:   ptr(AppAttestationFailureModeBlock),
				FcmCredentials:              &FcmCredentialsResponse{ProjectId: ptr("project-1"), ClientEmail: ptr("sa@project-1.iam.gserviceaccount.com")},
			},
		},
		{
			name: "get push",
			call: func(client Client) (any, int, error) {
				return client.GetPushAuthenticatorConfiguration()
			},
			responseBody:   `{"authenticatorId":"auth-4","isActive":true,"verificationMethod":"PUSH","pushProvider":"DEFAULT","apnsCredentials":{"teamId":"team-1","keyId":"key-1","bundleId":"com.example.app","sandbox":true}}`,
			expectedMethod: http.MethodGet,
			expectedPath:   "/v1/management/authenticator-configurations/push",
			expectedResponse: &PushAuthenticatorConfiguration{
				AuthenticatorId:    "auth-4",
				IsActive:           true,
				VerificationMethod: "PUSH",
				PushProvider:       ptr(PushConfigurationProviderDefault),
				ApnsCredentials: &ApnsCredentialsResponse{
					TeamId:   ptr("team-1"),
					KeyId:    ptr("key-1"),
					BundleId: ptr("com.example.app"),
					Sandbox:  ptr(true),
				},
			},
		},
		{
			name: "update push",
			call: func(client Client) (any, int, error) {
				return client.UpdatePushAuthenticatorConfiguration(UpdatePushAuthenticatorConfigurationBody{
					PushProvider:                SetValue(PushConfigurationProviderWebhook),
					WebhookUrl:                  SetValue("https://example.com/push"),
					CredentialLifetimeInMinutes: SetNull(int64(0)),
					ApnsCredentials:             SetNull(ApnsCredentialsUpdate{}),
					FcmCredentials:              SetNull(FcmCredentialsUpdate{}),
				})
			},
			responseBody:   `{"authenticatorId":"auth-4","isActive":true,"verificationMethod":"PUSH","pushProvider":"WEBHOOK","webhookUrl":"https://example.com/push"}`,
			expectedMethod: http.MethodPatch,
			expectedPath:   "/v1/management/authenticator-configurations/push",
			expectedBody:   `{"pushProvider":"WEBHOOK","webhookUrl":"https://example.com/push","credentialLifetimeInMinutes":null,"apnsCredentials":null,"fcmCredentials":null}`,
			expectedResponse: &PushAuthenticatorConfiguration{
				AuthenticatorId:    "auth-4",
				IsActive:           true,
				VerificationMethod: "PUSH",
				PushProvider:       ptr(PushConfigurationProviderWebhook),
				WebhookUrl:         ptr("https://example.com/push"),
			},
		},
		{
			name: "delete push",
			call: func(client Client) (any, int, error) {
				return client.DeletePushAuthenticatorConfiguration()
			},
			responseBody:     `{"success":true}`,
			expectedMethod:   http.MethodDelete,
			expectedPath:     "/v1/management/authenticator-configurations/push",
			expectedResponse: &HttpStatusResponse{Success: true},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			client, recorded := stubAuthsignal(t, http.StatusOK, testCase.responseBody)

			response, statusCode, err := testCase.call(client)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if statusCode != http.StatusOK {
				t.Fatalf("bad status code. expected: 200. got: %v", statusCode)
			}

			if recorded.method != testCase.expectedMethod {
				t.Fatalf("bad method. expected: %v. got: %v", testCase.expectedMethod, recorded.method)
			}

			if recorded.path != testCase.expectedPath {
				t.Fatalf("bad path. expected: %v. got: %v", testCase.expectedPath, recorded.path)
			}

			if recorded.body != testCase.expectedBody {
				t.Fatalf("bad request body. expected: %v. got: %v", testCase.expectedBody, recorded.body)
			}

			if !reflect.DeepEqual(response, testCase.expectedResponse) {
				t.Fatalf("bad response. expected: %+v. got: %+v", testCase.expectedResponse, response)
			}
		})
	}
}

func TestAuthenticatorConfigurationReadsSendNoBody(t *testing.T) {
	var contentType string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contentType = request.Header.Get("Content-Type")

		_, _ = writer.Write([]byte(`{"success":true}`))
	}))

	t.Cleanup(server.Close)

	client := NewClient(server.URL, "tenant-1", "secret")

	if _, _, err := client.DeletePasskeyAuthenticatorConfiguration(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if contentType != "" {
		t.Fatalf("expected no content type on a request without a body, got %v", contentType)
	}
}

func TestUpdateBodyOmitsUnsetFields(t *testing.T) {
	jsonBody, err := json.Marshal(UpdateSmsAuthenticatorConfigurationBody{})
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	if string(jsonBody) != "{}" {
		t.Fatalf("bad json. expected: {}. got: %v", string(jsonBody))
	}
}

func TestUpdateBodySendsFalseAndZeroValues(t *testing.T) {
	jsonBody, err := json.Marshal(UpdatePushAuthenticatorConfigurationBody{
		IsActive:                    SetValue(false),
		CredentialLifetimeInMinutes: SetValue(int64(0)),
		RequireAppAttestation:       SetValue(false),
	})
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := `{"isActive":false,"credentialLifetimeInMinutes":0,"requireAppAttestation":false}`

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got: %v", expectedJson, string(jsonBody))
	}
}

func TestBlockSizeRoundTrips(t *testing.T) {
	cases := []struct {
		name      string
		blockSize BlockSize
		json      string
	}{
		{name: "country", blockSize: BlockSize{Country: true}, json: `"country"`},
		{name: "numeric", blockSize: BlockSize{Numeric: 4}, json: `4`},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(testCase.blockSize)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			if string(jsonBody) != testCase.json {
				t.Fatalf("bad json. expected: %v. got: %v", testCase.json, string(jsonBody))
			}

			var decoded BlockSize
			if err := json.Unmarshal([]byte(testCase.json), &decoded); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			if decoded != testCase.blockSize {
				t.Fatalf("bad block size. expected: %+v. got: %+v", testCase.blockSize, decoded)
			}
		})
	}
}

func TestBlockSizeRejectsAnythingElse(t *testing.T) {
	var decoded BlockSize

	if err := json.Unmarshal([]byte(`"prefix"`), &decoded); err == nil {
		t.Fatalf("expected a string other than country to be rejected")
	}

	if err := json.Unmarshal([]byte(`0`), &decoded); err == nil {
		t.Fatalf("expected a non positive block size to be rejected")
	}
}

func TestAuthenticatorConfigurationErrorResponses(t *testing.T) {
	cases := []struct {
		name               string
		statusCode         int
		responseBody       string
		call               func(Client) (int, error)
		expectedInError    string
		expectedStatusCode int
	}{
		{
			name:         "unknown field is a bad request",
			statusCode:   http.StatusBadRequest,
			responseBody: `{"error":"invalid_request","errorDescription":"unknownField is not a known field."}`,
			call: func(client Client) (int, error) {
				_, statusCode, err := client.CreateSmsAuthenticatorConfiguration(CreateSmsAuthenticatorConfigurationBody{SmsProvider: SmsOtpProviderTnz})

				return statusCode, err
			},
			expectedInError:    "unknownField is not a known field.",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:         "creating a second configuration is a conflict",
			statusCode:   http.StatusConflict,
			responseBody: `{"error":"resource_already_exists","errorDescription":"A SMS OTP configuration already exists."}`,
			call: func(client Client) (int, error) {
				_, statusCode, err := client.CreateSmsAuthenticatorConfiguration(CreateSmsAuthenticatorConfigurationBody{SmsProvider: SmsOtpProviderTnz})

				return statusCode, err
			},
			expectedInError:    "A SMS OTP configuration already exists.",
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:         "reading a configuration the tenant does not have is a not found",
			statusCode:   http.StatusNotFound,
			responseBody: `{"error":"authenticator_configuration_not_found","errorDescription":"The authenticator configuration was not found."}`,
			call: func(client Client) (int, error) {
				_, statusCode, err := client.GetPushAuthenticatorConfiguration()

				return statusCode, err
			},
			expectedInError:    "The authenticator configuration was not found.",
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			client, _ := stubAuthsignal(t, testCase.statusCode, testCase.responseBody)

			statusCode, err := testCase.call(client)

			if err == nil {
				t.Fatalf("expected an error")
			}

			if statusCode != testCase.expectedStatusCode {
				t.Fatalf("bad status code. expected: %v. got: %v", testCase.expectedStatusCode, statusCode)
			}

			if !strings.Contains(err.Error(), testCase.expectedInError) {
				t.Fatalf("bad error. expected it to contain: %v. got: %v", testCase.expectedInError, err.Error())
			}
		})
	}
}

func TestUpdateBodySendsEmptyListForNilSlice(t *testing.T) {
	cases := []struct {
		name         string
		body         any
		expectedJson string
	}{
		{
			name:         "smsCountryCodes",
			body:         UpdateSmsAuthenticatorConfigurationBody{SmsCountryCodes: SetList[string](nil)},
			expectedJson: `{"smsCountryCodes":[]}`,
		},
		{
			name:         "prefixRateLimitConfigurations",
			body:         UpdateSmsAuthenticatorConfigurationBody{PrefixRateLimitConfigurations: SetList[PrefixRateLimitConfiguration](nil)},
			expectedJson: `{"prefixRateLimitConfigurations":[]}`,
		},
		{
			name:         "allowedCustomSmsVariables",
			body:         UpdateSmsAuthenticatorConfigurationBody{AllowedCustomSmsVariables: SetList[string](nil)},
			expectedJson: `{"allowedCustomSmsVariables":[]}`,
		},
		{
			name:         "allowedCustomEmailVariables",
			body:         UpdateEmailOtpAuthenticatorConfigurationBody{AllowedCustomEmailVariables: SetList[string](nil)},
			expectedJson: `{"allowedCustomEmailVariables":[]}`,
		},
		{
			name:         "expectedOrigins",
			body:         UpdatePasskeyAuthenticatorConfigurationBody{ExpectedOrigins: SetList[string](nil)},
			expectedJson: `{"expectedOrigins":[]}`,
		},
		{
			name:         "passkeyRegistrationHints",
			body:         UpdatePasskeyAuthenticatorConfigurationBody{PasskeyRegistrationHints: SetList[string](nil)},
			expectedJson: `{"passkeyRegistrationHints":[]}`,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(testCase.body)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			if string(jsonBody) != testCase.expectedJson {
				t.Fatalf("bad json. expected: %v. got: %v", testCase.expectedJson, string(jsonBody))
			}
		})
	}
}

func TestUpdateBodySendsNullForNullableListFields(t *testing.T) {
	cases := []struct {
		name         string
		body         any
		expectedJson string
	}{
		{
			name:         "sms sending rate limits cleared with SetNull",
			body:         UpdateSmsAuthenticatorConfigurationBody{SendingRateLimitConfigurations: SetNull([]RateLimitConfiguration{})},
			expectedJson: `{"sendingRateLimitConfigurations":null}`,
		},
		{
			name:         "email otp sending rate limits cleared with SetNull",
			body:         UpdateEmailOtpAuthenticatorConfigurationBody{SendingRateLimitConfigurations: SetNull([]RateLimitConfiguration{})},
			expectedJson: `{"sendingRateLimitConfigurations":null}`,
		},
		{
			name:         "push sending rate limits cleared with SetNull",
			body:         UpdatePushAuthenticatorConfigurationBody{SendingRateLimitConfigurations: SetNull([]RateLimitConfiguration{})},
			expectedJson: `{"sendingRateLimitConfigurations":null}`,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(testCase.body)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			if string(jsonBody) != testCase.expectedJson {
				t.Fatalf("bad json. expected: %v. got: %v", testCase.expectedJson, string(jsonBody))
			}
		})
	}
}

func TestListFieldIsOmittedWhenUnset(t *testing.T) {
	jsonBody, err := json.Marshal(UpdatePasskeyAuthenticatorConfigurationBody{IsActive: SetValue(true)})
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}

	expectedJson := `{"isActive":true}`

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got: %v", expectedJson, string(jsonBody))
	}
}

func TestListJsonInputMarshalsOmittedAndListStates(t *testing.T) {
	cases := []struct {
		name         string
		body         UpdatePasskeyAuthenticatorConfigurationBody
		expectedJson string
	}{
		{
			name:         "omitted when unset",
			body:         UpdatePasskeyAuthenticatorConfigurationBody{},
			expectedJson: `{}`,
		},
		{
			name:         "the list when set",
			body:         UpdatePasskeyAuthenticatorConfigurationBody{PasskeyRegistrationHints: SetList([]string{PasskeyRegistrationHintHybrid})},
			expectedJson: `{"passkeyRegistrationHints":["hybrid"]}`,
		},
		{
			name:         "an empty list when set to a nil slice",
			body:         UpdatePasskeyAuthenticatorConfigurationBody{PasskeyRegistrationHints: SetList[string](nil)},
			expectedJson: `{"passkeyRegistrationHints":[]}`,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(testCase.body)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			if string(jsonBody) != testCase.expectedJson {
				t.Fatalf("bad json. expected: %v. got: %v", testCase.expectedJson, string(jsonBody))
			}
		})
	}
}

func TestListJsonInputZeroValueMarshalsAsEmptyList(t *testing.T) {
	jsonBody, err := json.Marshal(UpdatePasskeyAuthenticatorConfigurationBody{ExpectedOrigins: &ListJsonInput[string]{}})
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}

	expectedJson := `{"expectedOrigins":[]}`

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got: %v", expectedJson, string(jsonBody))
	}
}
