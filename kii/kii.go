package kii

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Site int

const (
	US     Site = 0
	JP          = 1
	CN          = 2
	CUSTOM      = -1
)

type AppInfo struct {
	Site           Site
	Id             string
	Key            string
	CustomEndpoint string
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
	switch info.Site {
	case US:
	case JP:
	case CN:
	case CUSTOM:
	default:
		return nil, errors.New(fmt.Sprintf("Unknown kii.Site:%d", info.Site))
	}
	cx = &Context{*info}
	err = nil
	return
}

func (cx *Context) newRequest(method, path string, v interface{}, ctype string) (req *http.Request, err error) {
	// Build request JSON.
	var body *bytes.Buffer
	if v != nil {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}
	req, err = http.NewRequest(method, cx.endpoint(path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Kii-AppID", cx.app.Id)
	req.Header.Add("X-Kii-AppKey", cx.app.Key)
	if ctype != "" {
		req.Header.Add("Content-Type", ctype)
	}
	return
}

func (cx *Context) Login(name, pass string) (ux *UserContext, err error) {
	// TODO:
	return nil, nil
}

func (cx *Context) Admin(info *AdminInfo) (ax *AdminContext, err error) {
	// Build a request.
	req, err := cx.newRequest("POST", "/oauth2/token", oauthTokenRequest{
		ClientId:     info.Id,
		ClientSecret: info.Secret,
	}, "application/vnd.kii.OauthTokenRequest+json")
	if err != nil {
		return nil, err
	}

	// Do the request and read its response's body.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Dispatch with status code.
	switch resp.StatusCode {
	case 200:
		var ar oauthTokenResponse
		err = parseJson(resp.Body, &ar)
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
		err = parseJson(resp.Body, &er)
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

func (cx *Context) endpoint(path string) string {
	host := ""
	switch cx.app.Site {
	case US:
		host = "api.kii.com"
	case JP:
		host = "api-jp.kii.com"
	case CN:
		host = "api-cn2.kii.com"
	case CUSTOM:
		return cx.app.CustomEndpoint + path
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

// newRequest is alias of Context.newRequest
func (c *AdminContext) newRequest(method, path string, v interface{}, ctype string) (req *http.Request, err error) {
	return c.Context.newRequest(method, path, v, ctype)
}

func to_unixmsec(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// new_map create a map of clone, removed keys which starts with "_".
func new_map(src *map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{})
	if src == nil {
		return dst
	}
	for k, v := range *src {
		if strings.HasPrefix(k, "_") {
			continue
		}
		dst[k] = v
	}
	return dst
}

func (c *AdminContext) SendEvent(eventType, deviceID string, triggeredAt time.Time, values *map[string]interface{}) (success bool, err error) {
	// Create a request object.
	reqobj := new_map(values)
	reqobj["_type"] = eventType
	reqobj["_deviceID"] = eventType
	reqobj["_triggeredAt"] = to_unixmsec(triggeredAt)
	reqobj["_uploadedAt"] = to_unixmsec(time.Now())

	// Create a HTTP request.
	req, err := c.newRequest("POST", "/apps/{APP_ID}/events", reqobj, "application/vnd.kii.EventRecord+json")
	if err != nil {
		return false, err
	}

	// Do the request and read its response's body.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 204:
		return true, nil
	default:
		var er ErrorResponse
		er.StatusCode = resp.StatusCode
		err = parseJson(resp.Body, &er)
		if err != nil {
			return false, err
		}
		return false, err
	}
}

type Bucket struct {
	// TODO:
}
