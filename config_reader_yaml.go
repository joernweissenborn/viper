package viper

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v2"
)

type Yaml struct{}

func (Yaml) Unmarshal(in io.Reader, c map[string]interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	return yaml.Unmarshal(buf.Bytes(), &c)
}

func (Yaml) SupportedExtensions() []string {
	return []string{"yaml", "yml"}
}

func init() {
	RegisterConfigReader(Yaml{})
}
