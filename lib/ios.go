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
	for _, v := range c.data {
		var item Item = v
		fmt.Fprintf(w, "\"%v\" = \"%v\";\n", strings.Join(item.key, "."), escape(string(item.value)))
	}
}

func escape(s string) string {
	return strings.Replace(s, "\"", "\\\"", -1)
}
