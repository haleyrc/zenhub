package zenhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

const BaseURL = "https://api.zenhub.com/"

func New(token string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
		Token:      token,
	}
}

type Client struct {
	Debug      bool
	HTTPClient *http.Client
	Token      string
}

func (c *Client) get(path string, data interface{}) error {
	req, err := c.makeRequest(http.MethodGet, path, data)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if err := c.decode(resp, data); err != nil {
		return err
	}

	return nil
}

func (c *Client) makeRequest(method, path string, body interface{}) (*http.Request, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}

	url := BaseURL + path
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Authentication-Token", c.Token)

	if c.Debug {
		dumpRequest(req)
	}

	return req, nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		dumpResponse(resp)
	}

	return resp, nil
}

func (c *Client) decode(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

func dumpResponse(resp *http.Response) {
	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))
}

func dumpRequest(req *http.Request) {
	b, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(b))
}
