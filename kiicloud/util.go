package kiicloud

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// parseJson read from reader and parse it as JSON.
func parseJson(r io.Reader, v interface{}) (err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

// writeJson serialize v as JSON.
func writeJson(v interface{}) (*bytes.Buffer, error) {
	var body *bytes.Buffer
	if v != nil {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}
	return body, nil
}

// newMap clone a map, without keys which start with "_".
func newMap(src *map[string]interface{}) map[string]interface{} {
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
