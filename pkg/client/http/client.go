package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HttpClient interface {
	Get(path string) (*http.Response, error)
	Post(path string, body interface{}) (*http.Response, error)
	Put(path string, body interface{}) (*http.Response, error)
	Delete(path string) (*http.Response, error)
}

type httpClient struct {
	Host    string
	Headers map[string]string
}

func NewHttpClient(host string, headers map[string]string) HttpClient {
	return &httpClient{
		Host:    host,
		Headers: headers,
	}
}

func (c *httpClient) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.Host+path, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

func (c *httpClient) Post(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.Host+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

func (c *httpClient) Put(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", c.Host+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

func (c *httpClient) Delete(path string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.Host+path, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

func (c *httpClient) doRequest(req *http.Request) (*http.Response, error) {
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(body))

	return resp, nil
}
