package authsignal

import (
	"encoding/json"
	"testing"
)

func TestTenantResponseReadsDefaults(t *testing.T) {
	var tenant TenantResponse

	err := json.Unmarshal([]byte(`{"defaults":{"hideSuccessScreenOnEnrollment":false,"tokenDurationInMinutes":10}}`), &tenant)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	var tokenDuration int64
	if err := json.Unmarshal(tenant.Defaults["tokenDurationInMinutes"], &tokenDuration); err != nil || tokenDuration != 10 {
		t.Fatalf("expected a tokenDurationInMinutes default of 10. got: %s", tenant.Defaults["tokenDurationInMinutes"])
	}
}

func TestAuthenticatorConfigurationReadsDefaults(t *testing.T) {
	var passkey PasskeyAuthenticatorConfiguration

	err := json.Unmarshal([]byte(`{"authenticatorId":"abc","isActive":true,"defaults":{"userVerificationRequirement":"preferred"}}`), &passkey)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	var userVerification string
	if err := json.Unmarshal(passkey.Defaults["userVerificationRequirement"], &userVerification); err != nil || userVerification != UserVerificationRequirementPreferred {
		t.Fatalf("expected a userVerificationRequirement default of preferred. got: %s", passkey.Defaults["userVerificationRequirement"])
	}
}

func TestDefaultsAreNilWhenTheApiReportsNone(t *testing.T) {
	var tenant TenantResponse

	err := json.Unmarshal([]byte(`{"tokenDurationInMinutes":10}`), &tenant)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if tenant.Defaults != nil {
		t.Fatalf("expected no defaults. got: %v", tenant.Defaults)
	}
}
