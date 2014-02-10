package kiicloud

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
