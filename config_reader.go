package viper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

var configReaders = map[string]ConfigReader{}


// RegisterConfigReader registers a ConfigReader for its supported extensions.
//
// If there is already a reader registered for an extension, the newly registered reader will be used for this extension.
func RegisterConfigReader(r ConfigReader) {
	for _, ext := range r.SupportedExtensions() {
		configReaders[ext] = r
	}
}

// ConfigReader is an interface for a configuration file reader.
//
// Unmarshal must unmarshal an io.Reader into the map c.
// SupportedExtensions must deliver a list of file extensions supported by the reader.
type ConfigReader interface {
	Unmarshal(in io.Reader, c map[string]interface{}) error
	SupportedExtensions() []string
}

func unmarshallConfigReader(in io.Reader, c map[string]interface{}, configType string) error {
	configType = strings.ToLower(configType)
	configReader, supported := configReaders[configType]
	if !supported {
		return ConfigParseError{errors.New(fmt.Sprintf("Unspported file format '%s'.", configType))}
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	if err := configReader.Unmarshal(buf, c); err != nil {
		return ConfigParseError{err}

	}
	insensitiviseMap(c)
	return nil
}
