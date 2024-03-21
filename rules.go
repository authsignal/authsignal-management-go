package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Condition struct {
}

type Rule struct {
	Name                              string    `json:"name"`
	Description                       string    `json:"description"`
	IsActive                          bool      `json:"isActive"`
	Priority                          int32     `json:"priority"`
	ActionCode                        string    `json:"actionCode"`
	RuleId                            string    `json:"ruleId"`
	TenantId                          string    `json:"tenantId"`
	Type                              string    `json:"type"`
	VerificationMethods               []string  `json:"verificationMethods"`
	PromptToEnrollVerificationMethods []string  `json:"promptToEnrollVerificationMethods"`
	DefaultVerificationMethod         []string  `json:"defaultVerificationMethod"`
	Conditions                        Condition `json:"conditions"`
}

func (c Client) CreateRule(actionCode string, rule Rule) (*Rule, error) {
	createBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/action-configurations/%s", c.Host, actionCode), bytes.NewReader(createBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var createdRule Rule
	err = json.Unmarshal(body, &createdRule)
	if err != nil {
		return nil, err
	}

	return &createdRule, nil
}

func (c Client) GetRule(actionCode string, ruleId string) (*Rule, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/action-configurations/%s/rules/%s", c.Host, actionCode, ruleId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var rule Rule
	err = json.Unmarshal(body, &rule)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func (c Client) UpdateRule(actionCode string, rule Rule) (*Rule, error) {
	updateBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/action-configurations/%s/rules/%s", c.Host, actionCode, rule.RuleId), bytes.NewReader(updateBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var updatedRule Rule
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
