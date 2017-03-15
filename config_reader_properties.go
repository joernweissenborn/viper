package viper

import (
	"bytes"
	"io"
	"strings"

	"github.com/magiconair/properties"
)

type Properties struct{}

func (Properties) Unmarshal(in io.Reader, c map[string]interface{}) (err error) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	var p *properties.Properties
	if p, err = properties.Load(buf.Bytes(), properties.UTF8); err != nil {
		return ConfigParseError{err}
	}
	for _, key := range p.Keys() {
		value, _ := p.Get(key)
		// recursively build nested maps
		path := strings.Split(key, ".")
		lastKey := strings.ToLower(path[len(path)-1])
		deepestMap := deepSearch(c, path[0:len(path)-1])
		// set innermost value
		deepestMap[lastKey] = value
	}
	return

}

func (Properties) SupportedExtensions() []string {
	return []string{"properties", "props", "prop"}
}

func init() {
	RegisterConfigReader(Properties{})
}
