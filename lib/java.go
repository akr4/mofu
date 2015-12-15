package mofu

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type JavaProperties struct{}

func (j JavaProperties) Identifiers() []string {
	return []string{"java"}
}

func (j JavaProperties) AcceptFile(name string) bool {
	return filepath.Ext(name) == ".properties"
}

func (j JavaProperties) Write(c *Config, w io.Writer) {
	for k, v := range c.data {
		fmt.Fprintf(w, "%v = %v\n", strings.Join(NewKey(k), "."), v)
	}
}
