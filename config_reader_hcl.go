package viper

import (
	"bytes"
	"io"

	"github.com/hashicorp/hcl"
)

// Hcl provides support for the HCL format
type Hcl struct{}

func (Hcl) Unmarshal(in io.Reader, c map[string]interface{}) (err error) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	obj, err := hcl.Parse(string(buf.Bytes()))
	if err != nil {
		return
	}
	err = hcl.DecodeObject(&c, obj)
	return err
}

func (Hcl) SupportedExtensions() []string {
	return []string{"hcl"}
}

func init() {
	RegisterConfigReader(Hcl{})
}
