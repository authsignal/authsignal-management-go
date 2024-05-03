package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Condition interface {
}

type Rule struct {
	Name                              NullableJsonInput[string]    `json:"name,omitempty"`
	Description                       NullableJsonInput[string]    `json:"description,omitempty"`
	IsActive                          NullableJsonInput[bool]      `json:"isActive,omitempty"`
	Priority                          NullableJsonInput[int64]     `json:"priority,omitempty"`
	Type                              NullableJsonInput[string]    `json:"type,omitempty"`
	VerificationMethods               NullableJsonInput[[]string]  `json:"verificationMethods,omitempty"`
	PromptToEnrollVerificationMethods NullableJsonInput[[]string]  `json:"promptToEnrollVerificationMethods,omitempty"`
	DefaultVerificationMethod         NullableJsonInput[string]    `json:"defaultVerificationMethod,omitempty"`
	Conditions                        NullableJsonInput[Condition] `json:"conditions,omitempty"`
}

type RuleResponse struct {
	Name                              string    `json:"name"`
	Description                       string    `json:"description"`
	IsActive                          bool      `json:"isActive"`
	Priority                          int64     `json:"priority"`
	ActionCode                        string    `json:"actionCode"`
	RuleId                            string    `json:"ruleId"`
	TenantId                          string    `json:"tenantId"`
	Type                              string    `json:"type"`
	VerificationMethods               []string  `json:"verificationMethods"`
	PromptToEnrollVerificationMethods []string  `json:"promptToEnrollVerificationMethods"`
	DefaultVerificationMethod         string    `json:"defaultVerificationMethod"`
	Conditions                        Condition `json:"conditions"`
}

func (c Client) CreateRule(actionCode string, rule Rule) (*RuleResponse, error) {

	createBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/action-configurations/%s/rules", c.Host, actionCode), bytes.NewReader(createBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var createdRule RuleResponse
	err = json.Unmarshal(body, &createdRule)
	if err != nil {
		return nil, err
	}

	return &createdRule, nil
}

func (c Client) GetRule(actionCode string, ruleId string) (*RuleResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/action-configurations/%s/rules/%s", c.Host, actionCode, ruleId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var rule RuleResponse
	err = json.Unmarshal(body, &rule)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func (c Client) UpdateRule(actionCode string, ruleId string, rule Rule) (*RuleResponse, error) {
	updateBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/action-configurations/%s/rules/%s", c.Host, actionCode, ruleId), bytes.NewReader(updateBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var updatedRule RuleResponse
	err = json.Unmarshal(body, &updatedRule)
	if err != nil {
		return nil, err
	}

	return &updatedRule, nil
}

func (c Client) DeleteRule(actionCode string, ruleId string) (*HttpStatusResponse, error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/action-configurations/%s/rules/%s", c.Host, actionCode, ruleId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var httpStatusResponse HttpStatusResponse

	err = json.Unmarshal(body, &httpStatusResponse)
	if err != nil {
		return nil, err
	}

	return &httpStatusResponse, nil
}
