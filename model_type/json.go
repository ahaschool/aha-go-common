package model_type

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JsonType map[string]interface{}

func (jt JsonType) Value() (driver.Value, error) {
	j, err := json.Marshal(jt)
	return j, err
}

func (jt *JsonType) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	*jt, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
