package authsignal

import (
	"encoding/json"
	"testing"
)

func TestActionConfigurationJsonMarshal1(t *testing.T) {
	actionConfiguration := ActionConfiguration{
		ActionCode:              SetValue("hello-world"),
		DefaultUserActionResult: SetValue("ALLOW"),
	}

	jsonBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"defaultUserActionResult\":\"ALLOW\",\"actionCode\":\"hello-world\"}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestActionConfigurationJsonMarshal2(t *testing.T) {
	actionConfiguration := ActionConfiguration{
		ActionCode:              SetValue("hello-world"),
		DefaultUserActionResult: SetNull("ALLOW"),
	}

	jsonBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"defaultUserActionResult\":null,\"actionCode\":\"hello-world\"}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestActionConfigurationJsonMarshal3(t *testing.T) {
	actionConfiguration := ActionConfiguration{
		ActionCode: SetValue("hello-world"),
	}

	jsonBody, err := json.Marshal(actionConfiguration)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"actionCode\":\"hello-world\"}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestRuleJsonMarshal1(t *testing.T) {
	rule := Rule{
		Name:                SetValue("helloworld"),
		Priority:            SetValue(int64(29)),
		VerificationMethods: SetValue([]string{"SMS"}),
		IsActive:            SetValue(true),
	}

	jsonBody, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"name\":\"helloworld\",\"isActive\":true,\"priority\":29,\"verificationMethods\":[\"SMS\"]}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestRuleJsonMarshal2(t *testing.T) {
	var conditions Condition
	err := json.Unmarshal([]byte("{\"hello\":\"world\"}"), &conditions)

	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	rule := Rule{
		Name:                SetValue("helloworld"),
		Priority:            SetValue(int64(0)),
		VerificationMethods: SetNull([]string{"SMS"}),
		IsActive:            SetValue(false),
		Conditions:          SetValue(conditions),
	}

	jsonBody, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"name\":\"helloworld\",\"isActive\":false,\"priority\":0,\"verificationMethods\":null,\"conditions\":{\"hello\":\"world\"}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestRuleJsonMarshal3(t *testing.T) {
	var conditions Condition
	err := json.Unmarshal([]byte("{\"hello\":\"world\"}"), &conditions)

	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	rule := Rule{
		Name:                SetValue("helloworld"),
		Priority:            SetNull(int64(0)),
		VerificationMethods: SetNull([]string{"SMS"}),
		IsActive:            SetNull(false),
		Conditions:          SetNull(conditions),
	}

	jsonBody, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"name\":\"helloworld\",\"isActive\":null,\"priority\":null,\"verificationMethods\":null,\"conditions\":null}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}
