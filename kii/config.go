package kii

import (
	"errors"
	"os"
)

type Config struct {
	AppInfo   *AppInfo
	AdminInfo *AdminInfo
}

type configJson struct {
	AppSite      string `json:"app_site"`
	AppId        string `json:"app_id"`
	AppKey       string `json:"app_key"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func DefaultConfig() (c *Config, err error) {
	return LoadConfig("kiicloud.json")
}

func LoadConfig(path string) (c *Config, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cj configJson
	err = parseJson(f, &cj)
	if err != nil {
		return nil, err
	}
	var site Site
	switch cj.AppSite {
	case "US":
		site = US
	case "JP":
		site = JP
	case "CN":
		site = CN
	default:
		return nil, errors.New("Unknown app_site: " + cj.AppSite)
	}
	return &Config{
		&AppInfo{site, cj.AppId, cj.AppKey},
		&AdminInfo{cj.ClientId, cj.ClientSecret},
	}, err
}
