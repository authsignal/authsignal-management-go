package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ValueListItem interface{}

type ValueList struct {
	Name           NullableJsonInput[string]          `json:"name,omitempty"`
	Alias          NullableJsonInput[string]          `json:"alias,omitempty"`
	ItemType       NullableJsonInput[string]          `json:"itemType,omitempty"`
	ValueListItems NullableJsonInput[[]ValueListItem] `json:"valueListItems,omitempty"`
	IsActive       NullableJsonInput[bool]            `json:"isActive,omitempty"`
}

type ValueListResponse struct {
	Name           string          `json:"name"`
	Alias          string          `json:"alias"`
	ItemType       string          `json:"itemType"`
	ValueListItems []ValueListItem `json:"valueListItems"`
	IsActive       bool            `json:"isActive"`
}

func (c Client) CreateValueList(valueList ValueList) (*ValueListResponse, error) {
	createBody, err := json.Marshal(valueList)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/value-lists", c.Host), bytes.NewReader(createBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var createdValueList ValueListResponse
	err = json.Unmarshal(body, &createdValueList)
	if err != nil {
		return nil, err
	}

	return &createdValueList, nil
}

func (c Client) GetValueList(alias string) (*ValueListResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/value-lists/%s", c.Host, alias), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var valueList ValueListResponse

	err = json.Unmarshal(body, &valueList)
	if err != nil {
		return nil, err
	}

	return &valueList, nil
}

func (c Client) UpdateValueList(alias string, valueList ValueList) (*ValueListResponse, error) {
	updateBody, err := json.Marshal(valueList)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/value-lists/%s", c.Host, alias), bytes.NewReader(updateBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var updatedValueList ValueListResponse
	err = json.Unmarshal(body, &updatedValueList)
	if err != nil {
		return nil, err
	}

	return &updatedValueList, nil
}

func (c Client) DeleteValueList(alias string) (*HttpStatusResponse, error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/value-lists/%s", c.Host, alias), nil)
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
