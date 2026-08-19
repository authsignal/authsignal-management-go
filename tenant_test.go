package authsignal

import (
	"encoding/json"
	"testing"
)

func TestTenantSettingsOmitsUnsetFields(t *testing.T) {
	tenantSettings := TenantSettings{
		Name: SetValue("hello-world"),
	}

	jsonBody, err := json.Marshal(tenantSettings)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"name\":\"hello-world\"}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// A false setting has to reach the API. Plain bool fields with omitempty would drop it, leaving the
// setting on whatever it already was.
func TestTenantSettingsSendsFalseSettings(t *testing.T) {
	tenantSettings := TenantSettings{
		DisableRecoveryCodes:          SetValue(false),
		HideSuccessScreenOnEnrollment: SetValue(false),
	}

	jsonBody, err := json.Marshal(tenantSettings)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"disableRecoveryCodes\":false,\"hideSuccessScreenOnEnrollment\":false}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestTenantResponseDistinguishesUnsetFromFalse(t *testing.T) {
	var tenant TenantResponse

	err := json.Unmarshal([]byte("{\"tenantId\":\"abc\",\"disableRecoveryCodes\":false}"), &tenant)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if tenant.DisableRecoveryCodes == nil || *tenant.DisableRecoveryCodes {
		t.Fatalf("expected disableRecoveryCodes to be set to false")
	}

	if tenant.HideSuccessScreenOnEnrollment != nil {
		t.Fatalf("expected an absent hideSuccessScreenOnEnrollment to stay nil")
	}
}
