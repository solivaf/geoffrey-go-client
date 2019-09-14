package client

import (
	"encoding/json"
	"github.com/solivaf/geoffrey-go-client/internal/config"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type GeoffreyClient interface {
	GetConfig(app, profile string, model interface{}) error
}

type client struct {
	httpClient *http.Client
	url        string
}

func DefaultGeoffreyClient() GeoffreyClient {
	cfg := config.GetConfig()
	httpClient := &http.Client{Timeout: 10 * time.Second}

	return &client{url: cfg.Url, httpClient: httpClient}
}

func NewGeoffreyClient(url string, httpClient *http.Client) GeoffreyClient {
	if url == "" {
		url = config.GetConfig().Url
	}
	return &client{url: url, httpClient: httpClient}
}

func (c *client) GetConfig(app, profile string, model interface{}) error {
	url := c.getFormattedUrl(c.url, app, profile)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(b, &model); err != nil {
		return err
	}

	return nil
}

func (c *client) getFormattedUrl(urlString, app, profile string) string {
	if !strings.HasSuffix(urlString, "/") {
		urlString += "/"
	}
	urlString += app + "/" + profile
	return urlString
}
