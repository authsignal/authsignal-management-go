package authsignal

import (
	"encoding/json"
	"testing"
)

func TestActionFlowMarshalsNilRulesAsEmptyArray(t *testing.T) {
	flow := ActionFlow{ActionNodes: []ActionNode{}}

	jsonBody, err := marshalActionFlow(flow)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"actionNodes\":[],\"rules\":[]}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestActionFlowMarshalsExpectedVersionAndOmitsNilConditions(t *testing.T) {
	expectedFlowVersion := int64(3)

	flow := ActionFlow{
		ActionNodes: []ActionNode{
			ActionNode(`{"nodeId":"complete-ghi","nodeType":"COMPLETE"}`),
		},
		Rules: []ActionFlowRule{
			{RuleId: "rule-nz", Name: "From New Zealand", Conditions: map[string]any{"and": []any{}}},
			{RuleId: "rule-any", Name: "Always"},
		},
		ExpectedFlowVersion: &expectedFlowVersion,
	}

	jsonBody, err := marshalActionFlow(flow)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"actionNodes\":[{\"nodeId\":\"complete-ghi\",\"nodeType\":\"COMPLETE\"}]," +
		"\"rules\":[{\"ruleId\":\"rule-nz\",\"name\":\"From New Zealand\",\"conditions\":{\"and\":[]}}," +
		"{\"ruleId\":\"rule-any\",\"name\":\"Always\"}]," +
		"\"expectedFlowVersion\":3}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestActionFlowOmitsUnsetExpectedVersion(t *testing.T) {
	flow := ActionFlow{ActionNodes: []ActionNode{}, Rules: []ActionFlowRule{}}

	jsonBody, err := marshalActionFlow(flow)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"actionNodes\":[],\"rules\":[]}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestActionFlowResponseUnmarshal(t *testing.T) {
	var response ActionFlowResponse

	responseJson := "{\"tenantId\":\"abc\",\"actionCode\":\"sign-in\",\"actionType\":\"FLOW\"," +
		"\"actionNodes\":[{\"nodeId\":\"complete-ghi\",\"nodeType\":\"COMPLETE\"}],\"flowVersion\":2}"

	err := json.Unmarshal([]byte(responseJson), &response)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if response.ActionType != "FLOW" || response.FlowVersion != 2 {
		t.Fatalf("bad response. got : %+v", response)
	}

	if len(response.ActionNodes) != 1 {
		t.Fatalf("expected one action node. got : %+v", response.ActionNodes)
	}

	nodeJson, err := json.Marshal(response.ActionNodes)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedNodeJson := "[{\"nodeId\":\"complete-ghi\",\"nodeType\":\"COMPLETE\"}]"

	if string(nodeJson) != expectedNodeJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedNodeJson, string(nodeJson))
	}
}

func TestActionConfigurationResponseDistinguishesUnpublishedFlow(t *testing.T) {
	var unpublished ActionConfigurationResponse

	err := json.Unmarshal([]byte("{\"actionCode\":\"sign-in\",\"actionType\":\"FLOW\"}"), &unpublished)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if unpublished.FlowVersion != nil {
		t.Fatalf("expected an absent flowVersion to stay nil")
	}

	var published ActionConfigurationResponse

	err = json.Unmarshal([]byte("{\"actionCode\":\"sign-in\",\"actionType\":\"FLOW\",\"flowVersion\":1}"), &published)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if published.FlowVersion == nil || *published.FlowVersion != 1 {
		t.Fatalf("expected flowVersion 1. got : %+v", published.FlowVersion)
	}
}
