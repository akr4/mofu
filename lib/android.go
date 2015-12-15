package mofu

import (
	"fmt"
	"io"
	"strings"
)

type AndroidStrings struct{}

func (a AndroidStrings) Identifiers() []string {
	return []string{"android"}
}

func (a AndroidStrings) AcceptFile(name string) bool {
	return name == "strings.xml"
}

func (a AndroidStrings) Write(c *Config, w io.Writer) {
	fmt.Fprintln(w, `<?xml version="1.0" encoding="utf-8"?>
<resources>`)
	for k, v := range c.data {
		fmt.Fprintf(w, "    <string name=\"%v\">%v</string>\n", strings.Join(NewKey(k), "__"), v)
	}
	fmt.Fprintln(w, `</resources>`)
}
