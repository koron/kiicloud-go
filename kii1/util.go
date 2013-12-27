package kii

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

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
