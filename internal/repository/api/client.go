package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ApiCLient struct {
	baseURL      string
	client       *http.Client
	serviceToken string
}

type ApiClientResponse struct {
	StatusCode int
	Body       interface{}
}

func NewClient(baseURL string, serviceToken string) *ApiCLient {
	return &ApiCLient{
		baseURL:      baseURL,
		client:       &http.Client{},
		serviceToken: serviceToken,
	}
}

func (c *ApiCLient) DoRequest(method, path string, headers http.Header, payload interface{}) (*ApiClientResponse, error) {
	var jsonData []byte
	var err error
	if payload != nil {
		jsonData, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Service-Token", c.serviceToken)

	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseBody interface{}

	if err := json.Unmarshal(bodyData, &responseBody); err != nil {
		return nil, err
	}

	return &ApiClientResponse{
		StatusCode: resp.StatusCode,
		Body:       responseBody,
	}, nil
}

func (c *ApiCLient) Get(endpoint string, headers http.Header) (*ApiClientResponse, error) {
	return c.DoRequest(http.MethodGet, endpoint, headers, nil)
}

func (c *ApiCLient) Post(endpoint string, headers http.Header, payload interface{}) (*ApiClientResponse, error) {
	return c.DoRequest(http.MethodPost, endpoint, headers, payload)
}

func (c *ApiCLient) Put(endpoint string, headers http.Header, payload interface{}) (*ApiClientResponse, error) {
	return c.DoRequest(http.MethodPut, endpoint, headers, payload)
}

func (c *ApiCLient) Patch(endpoint string, headers http.Header, payload interface{}) (*ApiClientResponse, error) {
	return c.DoRequest(http.MethodPatch, endpoint, headers, payload)
}

func (c *ApiCLient) Delete(endpoint string, headers http.Header) (*ApiClientResponse, error) {
	return c.DoRequest(http.MethodDelete, endpoint, headers, nil)
}
