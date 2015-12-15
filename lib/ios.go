package mofu

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type IOSStrings struct{}

func (i IOSStrings) Identifiers() []string {
	return []string{"ios", "osx"}
}

func (i IOSStrings) AcceptFile(name string) bool {
	return filepath.Ext(name) == ".strings"
}

func (i IOSStrings) Write(c *Config, w io.Writer) {
	for k, v := range c.data {
		fmt.Fprintf(w, "\"%v\" = \"%v\";\n", strings.Join(NewKey(k), "."), escape(string(v)))
	}
}

func escape(s string) string {
	return strings.Replace(s, "\"", "\\\"", -1)
}
