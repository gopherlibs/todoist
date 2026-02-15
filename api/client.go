package api

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
	token      string
}

// This takes a full URL complete with endpoint and params.
func (c *Client) get(url *url.URL) (*http.Response, error) {

	slog.Debug("Client making a GET request.", "url", url.String())

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return c.httpClient.Do(req)
}

// This takes a full URL complete with endpoint, params, and an optional body.
func (c *Client) post(url *url.URL, body io.Reader) (*http.Response, error) {

	slog.Debug("Client making a POST request.", "url", url.String())

	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return c.httpClient.Do(req)
}

// This takes a full URL complete with endpoint, params, and an optional body.
func (c *Client) delete(url *url.URL) (*http.Response, error) {

	slog.Debug("Client making a DELETE request.", "url", url.String())

	req, err := http.NewRequest("DELETE", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return c.httpClient.Do(req)
}

// Returns a new client, a must to use this package.
func New(token string) *Client {

	baseURL, err := url.Parse("https://api.todoist.com")
	if err != nil {
		return nil
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   baseURL,
		userAgent: "Gopherlibs/Todoist",
		token:     token,
	}
}
