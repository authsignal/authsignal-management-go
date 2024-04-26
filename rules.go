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
	Name                              string    `json:"name,omitempty"`
	Description                       string    `json:"description,omitempty"`
	IsActive                          bool      `json:"isActive,omitempty"`
	Priority                          int64     `json:"priority,omitempty"`
	ActionCode                        string    `json:"actionCode,omitempty"`
	RuleId                            string    `json:"ruleId,omitempty"`
	TenantId                          string    `json:"tenantId,omitempty"`
	Type                              string    `json:"type,omitempty"`
	VerificationMethods               []string  `json:"verificationMethods,omitempty"`
	PromptToEnrollVerificationMethods []string  `json:"promptToEnrollVerificationMethods,omitempty"`
	DefaultVerificationMethod         string    `json:"defaultVerificationMethod,omitempty"`
	Conditions                        Condition `json:"conditions,omitempty"`
}

func (c Client) CreateRule(actionCode string, rule Rule) (*Rule, error) {
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

func (c Client) UpdateRule(actionCode string, ruleId string, rule Rule) (*Rule, error) {
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
