package kiicloud

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"time"
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
func writeJson(v interface{}) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
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

// toUnixMsec converts time.Time to milliseconds from 1970/01/01 00:00:00 UTC.
func toUnixMsec(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
