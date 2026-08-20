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
