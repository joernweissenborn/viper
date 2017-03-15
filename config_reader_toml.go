package viper

import (
	"io"

	"github.com/pelletier/go-toml"
)

type Toml struct{}

func (Toml) Unmarshal(in io.Reader, c map[string]interface{}) (err error) {
	tree, err := toml.LoadReader(in)
	if err != nil {
		return
	}
	tmap := tree.ToMap()
	for k, v := range tmap {
		c[k] = v
	}
	return

}
func (Toml) SupportedExtensions() []string {
	return []string{"toml"}
}

func init() {
	RegisterConfigReader(Toml{})
}
