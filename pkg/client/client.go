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
	GetConfig(app, profile string) (config map[string]interface{}, err error)
}

type client struct {
	httpClient *http.Client
	url        string
}

func DefaultGeoffreyClient() GeoffreyClient {
	cfg := config.GetConfig()
	t, err := time.ParseDuration(cfg.ClientTimeout)
	if err != nil {
		panic(err)
	}
	httpClient := &http.Client{Timeout: t}

	return &client{url: cfg.Url, httpClient: httpClient}
}

func NewGeoffreyClient(url string, httpClient *http.Client) GeoffreyClient {
	if url == "" {
		url = config.GetConfig().Url
	}
	return &client{url: url, httpClient: httpClient}
}

func (c *client) GetConfig(app, profile string) (map[string]interface{}, error) {
	url := c.getFormattedUrl(c.url, app, profile)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var config map[string]interface{}
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *client) getFormattedUrl(urlString, app, profile string) string {
	if !strings.HasSuffix(urlString, "/") {
		urlString += "/"
	}
	urlString += app + "/" + profile
	return urlString
}
