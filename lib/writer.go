package mofu

import "io"

type Writer interface {
	Identifiers() []string
	AcceptFile(name string) bool
	Write(c *Config, w io.Writer)
}
