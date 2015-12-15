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
	for _, v := range c.data {
		var item Item = v
		fmt.Fprintf(w, "    <string name=\"%v\">%v</string>\n", strings.Join(item.key, "__"), item.value)
	}
	fmt.Fprintln(w, `</resources>`)
}
