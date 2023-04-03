package dingtalk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type ClientConfig struct {
	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error

	Jar     http.CookieJar
	Timeout time.Duration
}

type Client struct {
	HTTPClient *http.Client
}

func newClient(config ClientConfig) *Client {
	return &Client{HTTPClient: &http.Client{
		Transport:     config.Transport,
		CheckRedirect: config.CheckRedirect,
		Jar:           config.Jar,
		Timeout:       config.Timeout,
	}}
}

func (c *Client) request(method string, urlString string, params map[string]string, body map[string]any, v any) error {
	var reqBytes []byte
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Create request
	req, err := http.NewRequest(method, urlString, bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	// Add params to URL
	if params != nil {
		values := url.Values{}
		for key, value := range params {
			values.Add(key, value)
		}
		req.URL.RawQuery = values.Encode()
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Send request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
