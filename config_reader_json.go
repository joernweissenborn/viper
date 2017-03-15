package viper

import (
	"encoding/json"
	"io"
)

type Json struct{}

func (Json) Unmarshal(in io.Reader, c map[string]interface{}) error {
	dec := json.NewDecoder(in)
	return dec.Decode(&c)
}

func (Json) SupportedExtensions() []string {
	return []string{"json"}
}

func init() {
	RegisterConfigReader(Json{})
}
