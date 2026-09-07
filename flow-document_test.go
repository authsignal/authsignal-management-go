package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// A flow document is passed through untouched, so these values have to survive a round trip: an
// integer above 2^53, a trailing zero on a decimal, a null and a field the SDK does not know.
const awkwardNode = `{"nodeId":"challenge-abc","nodeType":"CHALLENGE","attempts":9007199254740993,` +
	`"score":1.10,"label":null,"somethingNew":{"nested":[1,2,3]}}`

const awkwardCondition = `{"and":[{"==":[{"var":"attempts"},9007199254740993]},` +
	`{"==":[{"var":"score"},1.10]},{"==":[{"var":"label"},null]},{"somethingNew":true}]}`

func compact(t *testing.T, jsonBody []byte) string {
	t.Helper()

	var compacted bytes.Buffer

	err := json.Compact(&compacted, jsonBody)
	if err != nil {
		t.Fatalf("failed to compact json")
	}

	return compacted.String()
}

func newMockServer(t *testing.T, responseBody string, capturedRequestBody *string) Client {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestBody, err := io.ReadAll(request.Body)
		if err != nil {
			t.Errorf("failed to read the request body")
		}

		if capturedRequestBody != nil {
			*capturedRequestBody = string(requestBody)
		}

		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprint(writer, responseBody)
	}))

	t.Cleanup(server.Close)

	return NewClient(server.URL, "tenant-abc", "secret-abc")
}

func TestUpdateActionFlowSendsTheDocumentVerbatim(t *testing.T) {
	var requestBody string

	expectedFlowVersion := int64(2)

	client := newMockServer(t, "{}", &requestBody)

	_, _, err := client.UpdateActionFlow("sign-in", ActionFlow{
		ActionNodes: []ActionNode{ActionNode(awkwardNode)},
		Rules: []ActionFlowRule{
			{RuleId: "rule-abc", Name: "Awkward", Conditions: json.RawMessage(awkwardCondition)},
		},
		ExpectedFlowVersion: &expectedFlowVersion,
	})
	if err != nil {
		t.Fatalf("failed to update the action flow: %v", err)
	}

	expectedJson := fmt.Sprintf("{\"actionNodes\":[%s],\"rules\":[{\"ruleId\":\"rule-abc\","+
		"\"name\":\"Awkward\",\"conditions\":%s}],\"expectedFlowVersion\":2}", awkwardNode, awkwardCondition)

	if requestBody != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, requestBody)
	}
}

func TestUpdateActionFlowReturnsTheDocumentVerbatim(t *testing.T) {
	responseBody := fmt.Sprintf("{\"tenantId\":\"tenant-abc\",\"actionCode\":\"sign-in\","+
		"\"actionType\":\"FLOW\",\"actionNodes\":[%s],\"flowVersion\":2}", awkwardNode)

	client := newMockServer(t, responseBody, nil)

	flow, _, err := client.UpdateActionFlow("sign-in", ActionFlow{ActionNodes: []ActionNode{}})
	if err != nil {
		t.Fatalf("failed to update the action flow: %v", err)
	}

	if len(flow.ActionNodes) != 1 {
		t.Fatalf("expected one action node. got : %+v", flow.ActionNodes)
	}

	if compact(t, flow.ActionNodes[0]) != awkwardNode {
		t.Fatalf("bad json. expected: %v. got : %v", awkwardNode, compact(t, flow.ActionNodes[0]))
	}
}

func TestListRulesReturnsConditionsVerbatim(t *testing.T) {
	responseBody := fmt.Sprintf("[{\"ruleId\":\"rule-abc\",\"name\":\"Awkward\",\"conditions\":%s}]", awkwardCondition)

	client := newMockServer(t, responseBody, nil)

	rules, _, err := client.ListRules("sign-in")
	if err != nil {
		t.Fatalf("failed to list the rules: %v", err)
	}

	if len(rules) != 1 {
		t.Fatalf("expected one rule. got : %+v", rules)
	}

	conditionsJson, err := json.Marshal(rules[0].Conditions)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	if string(conditionsJson) != awkwardCondition {
		t.Fatalf("bad json. expected: %v. got : %v", awkwardCondition, string(conditionsJson))
	}
}

func TestGetActionConfigurationReturnsActionNodesVerbatim(t *testing.T) {
	responseBody := fmt.Sprintf("{\"actionCode\":\"sign-in\",\"actionType\":\"FLOW\","+
		"\"actionNodes\":[%s],\"flowVersion\":2}", awkwardNode)

	client := newMockServer(t, responseBody, nil)

	actionConfiguration, _, err := client.GetActionConfiguration("sign-in")
	if err != nil {
		t.Fatalf("failed to get the action configuration: %v", err)
	}

	if len(actionConfiguration.ActionNodes) != 1 {
		t.Fatalf("expected one action node. got : %+v", actionConfiguration.ActionNodes)
	}

	if compact(t, actionConfiguration.ActionNodes[0]) != awkwardNode {
		t.Fatalf("bad json. expected: %v. got : %v", awkwardNode, compact(t, actionConfiguration.ActionNodes[0]))
	}
}
