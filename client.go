package zenhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"
)

const DefaultURL = "https://api.zenhub.com"

func New(token string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Token:      token,
		URL:        DefaultURL,
	}
}

type Client struct {
	Debug      bool
	HTTPClient *http.Client
	Token      string
	URL        string
}

func (c *Client) get(path string, data interface{}) error {
	req, err := c.makeRequest(http.MethodGet, path, nil)
	if err != nil {
		return err
	}

start:
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 403 {
		waitUntil(resp.Header.Get("Date"), resp.Header.Get("X-RateLimit-Reset"))
		goto start
	}

	if err := c.decode(resp, data); err != nil {
		return err
	}

	return nil
}

func waitUntil(date, resetTime string) {
	curr, _ := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", date)
	curr = curr.Local()
	reset, _ := strconv.ParseInt(resetTime, 10, 64)
	d := time.Unix(reset, 0).Local().Sub(curr)
	fmt.Println("Waiting", d)
	<-time.After(d)
}

func (c *Client) makeRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	url := c.URL + path
	req, err := http.NewRequest(method, url, &buf)
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
	fmt.Fprintln(os.Stderr, string(b))
}

func dumpRequest(req *http.Request) {
	b, _ := httputil.DumpRequest(req, true)
	fmt.Fprintln(os.Stderr, string(b))
}
