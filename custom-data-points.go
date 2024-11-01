package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomDataPoint struct {
	Id          NullableJsonInput[string] `json:"id,omitempty"`
	Name        NullableJsonInput[string] `json:"name,omitempty"`
	DataType    NullableJsonInput[string] `json:"dataType,omitempty"`
	ModelType   NullableJsonInput[string] `json:"modelType,omitempty"`
	Description NullableJsonInput[string] `json:"description,omitempty"`
}

type CustomDataPointResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DataType    string `json:"dataType"`
	ModelType   string `json:"modelType"`
	Description string `json:"description"`
}

func (c Client) CreateCustomDataPoint(customDataPoint CustomDataPoint) (*CustomDataPointResponse, int, error) {
	createBody, err := json.Marshal(customDataPoint)
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/custom-data-points", c.Host), bytes.NewReader(createBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var createdCustomDataPoint CustomDataPointResponse
	err = json.Unmarshal(body, &createdCustomDataPoint)
	if err != nil {
		return nil, statusCode, err
	}

	return &createdCustomDataPoint, statusCode, nil
}

func (c Client) GetCustomDataPoint(id string) (*CustomDataPointResponse, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/custom-data-points/%s", c.Host, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var customDataPoint CustomDataPointResponse

	err = json.Unmarshal(body, &customDataPoint)
	if err != nil {
		return nil, statusCode, err
	}

	return &customDataPoint, statusCode, nil
}

func (c Client) DeleteCustomDataPoint(id string) (*HttpStatusResponse, int, error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/custom-data-points/%s", c.Host, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var httpStatusResponse HttpStatusResponse

	err = json.Unmarshal(body, &httpStatusResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &httpStatusResponse, statusCode, nil
}
