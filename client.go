package authsignal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Host       string
	ApiSecret  string
	TenantId   string
	HttpClient *http.Client
}

type HttpStatusResponse struct {
	Success          bool
	StatusCode       int64
	Message          string
	Error            string
	ErrorDescription string
}

func NewClient(host string, tenantId string, apiSecret string) Client {
	return Client{
		Host:       host,
		TenantId:   tenantId,
		ApiSecret:  apiSecret,
		HttpClient: &http.Client{},
	}
}

func (c Client) makeRequest(request *http.Request, apiSecret string) ([]byte, int, error) {
	request.SetBasicAuth(apiSecret, "")

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, 0, err
	}

	defer response.Body.Close()

	if response.StatusCode > 499 {
		return nil, response.StatusCode, fmt.Errorf("bad request to %s, http status code of %d, status was: %s", request.URL, response.StatusCode, response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	if response.StatusCode > 299 {
		var responseBodyToPrint HttpStatusResponse
		json.Unmarshal(responseBody, &responseBodyToPrint)
		return nil, response.StatusCode, fmt.Errorf("request to %s failed.\n    status: %s\n    error: %s\n    error description: %s", request.URL, response.Status, responseBodyToPrint.Error, responseBodyToPrint.ErrorDescription)
	}

	return responseBody, response.StatusCode, nil
}
