package kii

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site int

const (
	US Site = 0
	JP      = 1
	CN      = 2
)

type AppInfo struct {
	Site Site
	Id   string
	Key  string
}

type AdminInfo struct {
	Id     string
	Secret string
}

type Context struct {
	app AppInfo
	// TODO:
}

func NewContext(info *AppInfo) (cx *Context, err error) {
	switch (info.Site) {
	case US:
	case JP:
	case CN:
	default:
		return nil, errors.New(fmt.Sprintf("Unknown kii.Site:%d", info.Site))
	}
	cx = &Context{*info}
	err = nil
	return
}

func (cx *Context) Login(name, pass string) (ux *UserContext, err error) {
	// TODO:
	return nil, nil
}

func (cx *Context) Admin(info *AdminInfo) (ax *AdminContext, err error) {
	// Build request JSON.
	b, err := json.Marshal(oauthTokenRequest{
		ClientId:     info.Id,
		ClientSecret: info.Secret,
	})
	if err != nil {
		return nil, err
	}

	// Build a request.
	req, err := http.NewRequest("POST", cx.restUrl("/oauth2/token"), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Kii-AppID", cx.app.Id)
	req.Header.Add("X-Kii-AppKey", cx.app.Key)
	req.Header.Add("Content-Type", "application/vnd.kii.OauthTokenRequest+json")

	// Do the request and read its response's body.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Dispatch with status code.
	switch resp.StatusCode {
	case 200:
		var ar oauthTokenResponse
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}
		ax = &AdminContext{
			Context: cx,
			Token: Token{
				Id:          ar.Id,
				AccessToken: ar.AccessToken,
				ExpiresIn:   ar.ExpiresIn,
				TokenType:   ar.TokenType,
			},
		}
	default:
		var er ErrorResponse
		er.StatusCode = resp.StatusCode
		err = json.Unmarshal(body, &er)
		if err != nil {
			return nil, err
		}
		err = &er
	}
	return
}

func (cx *Context) Bucket(name string) (b *Bucket, err error) {
	// TODO:
	return nil, nil
}

func (cx *Context) restUrl(path string) string {
	host := ""
	switch cx.app.Site {
	case US:
		host = "api.kii.com"
	case JP:
		host = "api-jp.kii.com"
	case CN:
		host = "api-cn2.kii.com"
	}
	return "https://" + host + "/api" + path
}

type UserContext struct {
	// TODO:
}

type AdminContext struct {
	Context *Context
	Token
}

type Bucket struct {
	// TODO:
}
