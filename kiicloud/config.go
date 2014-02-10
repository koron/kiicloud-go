package kiicloud

import (
	"io"
	"os"
)

// Config hold all configurations of the app.
type Config struct {
	EndPoint     string `json:"site"`
	AppId        string `json:"app_id"`
	AppKey       string `json:"app_key"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// DefaultConfig load config from "kiicloud.json"
func DefaultConfig() (c *Config, err error) {
	return LoadConfig("kiicloud.json")
}

// LoadConfig load a Config from file.
func LoadConfig(path string) (c *Config, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readConfig(f)
}

// readConfig read and parse Config from io.Reader
func readConfig(r io.Reader) (c *Config, err error) {
	var conf Config
	err = parseJson(r, &conf)
	if err != nil {
		return nil, err
	}
	// Check and arrange config.
	conf.EndPoint, err = expandSite(conf.EndPoint)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

// NewClient create a client using Config.
func (c *Config) NewClient() (*Client, error) {
	return NewClient(c.EndPoint, c.AppId, c.AppKey)
}

// NewAdminClient create a admin client using Config.
func (c *Config) NewAdminClient() (*AdminClient, error) {
	return NewAdminClient(c.EndPoint, c.AppId, c.AppKey,
		c.ClientId, c.ClientSecret)
}
