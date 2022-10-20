package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

func serializeBodyToBuffer(body interface{}) (*bytes.Buffer, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("cannot marshal body to json")
	}

	return bytes.NewBuffer(jsonBody), nil
}

func serializeResponseToStruct(body io.ReadCloser, res *map[string]any) error {
	raw, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &res)
	if err != nil {
		return err
	}

	return nil
}
