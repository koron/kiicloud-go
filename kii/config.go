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
	AppSite        string `json:"app_site"`
	AppId          string `json:"app_id"`
	AppKey         string `json:"app_key"`
	ClientId       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	CustomEndpoint string `json:"custom_endpoint"`
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
	case "CUSTOM":
		if len(cj.CustomEndpoint) == 0 {
			return nil, errors.New("No custom_endpoint")
		}
		site = CUSTOM
	default:
		return nil, errors.New("Unknown app_site: " + cj.AppSite)
	}
	return &Config{
		&AppInfo{site, cj.AppId, cj.AppKey, cj.CustomEndpoint},
		&AdminInfo{cj.ClientId, cj.ClientSecret},
	}, err
}
