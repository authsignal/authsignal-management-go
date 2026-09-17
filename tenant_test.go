package authsignal

import (
	"encoding/json"
	"testing"
)

func TestTenantSettingsOmitsUnsetFields(t *testing.T) {
	tenantSettings := TenantSettings{}

	jsonBody, err := json.Marshal(tenantSettings)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// A false setting has to reach the API. A plain bool field with omitempty would drop it, leaving the
// setting on whatever it already was.
func TestTenantSettingsSendsFalseSettings(t *testing.T) {
	tenantSettings := TenantSettings{
		HideSuccessScreenOnEnrollment: SetValue(false),
	}

	jsonBody, err := json.Marshal(tenantSettings)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"hideSuccessScreenOnEnrollment\":false}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestTenantResponseDistinguishesUnsetFromFalse(t *testing.T) {
	var setToFalse TenantResponse

	err := json.Unmarshal([]byte("{\"hideSuccessScreenOnEnrollment\":false}"), &setToFalse)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if setToFalse.HideSuccessScreenOnEnrollment == nil || *setToFalse.HideSuccessScreenOnEnrollment {
		t.Fatalf("expected hideSuccessScreenOnEnrollment to be set to false")
	}

	var unset TenantResponse

	err = json.Unmarshal([]byte("{\"tenantId\":\"abc\"}"), &unset)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if unset.HideSuccessScreenOnEnrollment != nil {
		t.Fatalf("expected an absent hideSuccessScreenOnEnrollment to stay nil")
	}
}

func assertJson(t *testing.T, value any, expected string) {
	t.Helper()

	jsonBody, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	if string(jsonBody) != expected {
		t.Fatalf("bad json. expected: %v. got : %v", expected, string(jsonBody))
	}
}

func TestTenantSettingsSendsWebhookConfigurations(t *testing.T) {
	settings := TenantSettings{
		AuthenticatorEventsWebhookConfig: SetValue(AuthenticatorEventsWebhookConfig{
			Url:                        ptr("https://example.com/events"),
			IncludeCredentialPublicKey: ptr(true),
		}),
		LogEventsWebhookConfig: SetValue(LogEventsWebhookConfig{
			EndpointUrl: ptr("https://example.com/logs"),
		}),
	}

	assertJson(t, settings, "{\"authenticatorEventsWebhookConfig\":{\"url\":\"https://example.com/events\",\"includeCredentialPublicKey\":true},\"logEventsWebhookConfig\":{\"endpointUrl\":\"https://example.com/logs\"}}")
}

func TestTenantSettingsSendsOneWebhookFieldOnItsOwn(t *testing.T) {
	settings := TenantSettings{
		AuthenticatorEventsWebhookConfig: SetValue(AuthenticatorEventsWebhookConfig{
			IncludeCredentialPublicKey: ptr(false),
		}),
	}

	assertJson(t, settings, "{\"authenticatorEventsWebhookConfig\":{\"includeCredentialPublicKey\":false}}")
}

func TestTenantSettingsClearsWebhookConfigurations(t *testing.T) {
	settings := TenantSettings{
		AuthenticatorEventsWebhookConfig: SetNull(AuthenticatorEventsWebhookConfig{}),
		LogEventsWebhookConfig:           SetNull(LogEventsWebhookConfig{}),
	}

	assertJson(t, settings, "{\"authenticatorEventsWebhookConfig\":null,\"logEventsWebhookConfig\":null}")
}

func TestTenantSettingsSendsAnEmptyIpAllowlist(t *testing.T) {
	assertJson(t, TenantSettings{IpWhitelist: SetList([]string{})}, "{\"ipWhitelist\":[]}")
	assertJson(t, TenantSettings{IpWhitelist: SetList[string](nil)}, "{\"ipWhitelist\":[]}")
}

func TestTenantSettingsSendsAPopulatedIpAllowlist(t *testing.T) {
	settings := TenantSettings{IpWhitelist: SetList([]string{"203.0.113.0/24", "198.51.100.7/32"})}

	assertJson(t, settings, "{\"ipWhitelist\":[\"203.0.113.0/24\",\"198.51.100.7/32\"]}")
}

func TestTenantSettingsSendsTokenDuration(t *testing.T) {
	assertJson(t, TenantSettings{TokenDurationInMinutes: SetValue(int64(15))}, "{\"tokenDurationInMinutes\":15}")
}

func TestTenantResponseReadsWebhookConfigurations(t *testing.T) {
	var tenant TenantResponse

	body := "{\"tokenDurationInMinutes\":15,\"authenticatorEventsWebhookConfig\":{\"url\":\"https://example.com/events\",\"includeCredentialPublicKey\":false},\"logEventsWebhookConfig\":{\"endpointUrl\":\"https://example.com/logs\"},\"ipWhitelist\":[]}"

	if err := json.Unmarshal([]byte(body), &tenant); err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if tenant.TokenDurationInMinutes == nil || *tenant.TokenDurationInMinutes != 15 {
		t.Fatalf("expected tokenDurationInMinutes to be 15")
	}

	if tenant.AuthenticatorEventsWebhookConfig == nil || *tenant.AuthenticatorEventsWebhookConfig.Url != "https://example.com/events" {
		t.Fatalf("expected the authenticator events endpoint to be read back")
	}

	if tenant.AuthenticatorEventsWebhookConfig.IncludeCredentialPublicKey == nil || *tenant.AuthenticatorEventsWebhookConfig.IncludeCredentialPublicKey {
		t.Fatalf("expected includeCredentialPublicKey to be set to false")
	}

	if tenant.LogEventsWebhookConfig == nil || *tenant.LogEventsWebhookConfig.EndpointUrl != "https://example.com/logs" {
		t.Fatalf("expected the log events endpoint to be read back")
	}

	if tenant.IpWhitelist == nil || len(*tenant.IpWhitelist) != 0 {
		t.Fatalf("expected an empty ipWhitelist to be read back as empty rather than nil")
	}
}

func TestTenantResponseLeavesUnconfiguredWebhooksNil(t *testing.T) {
	var tenant TenantResponse

	if err := json.Unmarshal([]byte("{\"tenantId\":\"abc\"}"), &tenant); err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if tenant.AuthenticatorEventsWebhookConfig != nil || tenant.LogEventsWebhookConfig != nil || tenant.IpWhitelist != nil {
		t.Fatalf("expected unconfigured settings to stay nil")
	}
}
