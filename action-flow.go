package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ActionFlowRule struct {
	RuleId     string    `json:"ruleId"`
	Name       string    `json:"name"`
	Conditions Condition `json:"conditions,omitempty"`
}

type ActionFlow struct {
	ActionNodes         []ActionNode     `json:"actionNodes"`
	Rules               []ActionFlowRule `json:"rules"`
	ExpectedFlowVersion *int64           `json:"expectedFlowVersion,omitempty"`
}

type ActionFlowResponse struct {
	TenantId    string       `json:"tenantId"`
	ActionCode  string       `json:"actionCode"`
	ActionType  string       `json:"actionType"`
	ActionNodes []ActionNode `json:"actionNodes"`
	FlowVersion int64        `json:"flowVersion"`
}

// The API's schema requires the rules array, so a nil slice has to marshal as [] not null.
func marshalActionFlow(flow ActionFlow) ([]byte, error) {
	if flow.Rules == nil {
		flow.Rules = []ActionFlowRule{}
	}

	return json.Marshal(flow)
}

func (c Client) UpdateActionFlow(actionCode string, flow ActionFlow) (*ActionFlowResponse, int, error) {
	updateBody, err := marshalActionFlow(flow)
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/action-configurations/%s/flow", c.Host, actionCode), bytes.NewReader(updateBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var updatedActionFlow ActionFlowResponse

	// Numbers in the returned document have to re-marshal exactly as the API sent them.
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()

	err = decoder.Decode(&updatedActionFlow)
	if err != nil {
		return nil, statusCode, err
	}

	return &updatedActionFlow, statusCode, nil
}
