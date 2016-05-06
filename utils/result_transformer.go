package utils

import (
	"encoding/json"
)

type ResultTransformer struct {
	value interface{}
}

// NewResultTransformer return new ResultTransformer object
func NewResultTransformer(value interface{}) *ResultTransformer {
	return &ResultTransformer{value}
}

// Set value
func (rt *ResultTransformer) Set(value interface{}) {
	rt.value = value
}

// Get value
func (rt *ResultTransformer) Get() interface{} {
	return rt.value
}

// ToJson return json
func (rt *ResultTransformer) ToJson() (string, error) {

	json, err := json.MarshalIndent(rt.value, "", "  ")
	if err != nil {
		return "", err
	}

	return string(json), nil
}
