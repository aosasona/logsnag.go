package utils

import (
	"encoding/json"
	"io"
)

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

func StructToMap(data interface{}) (map[string]interface{}, error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	err = json.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
