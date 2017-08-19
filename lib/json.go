package mofu

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type Json struct{}

func (j Json) Identifiers() []string {
	return []string{"json"}
}

func (j Json) AcceptFile(name string) bool {
	return filepath.Ext(name) == ".json"
}

func (j Json) Write(c *Config, w io.Writer) {
	lines := make([]string, len(c.data))
	idx := 0
	for k, v := range c.data {
		lines[idx] = fmt.Sprintf("  \"%v\": \"%v\"", strings.Join(NewKey(k), "."), v)
		idx++
	}

	fmt.Fprintln(w, "{")
	fmt.Fprintln(w, strings.Join(lines, ",\n"))
	fmt.Fprintln(w, "}")
}
