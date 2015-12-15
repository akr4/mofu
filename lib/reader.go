package mofu

import "io"

type ConfigReader interface {
	Read(r io.Reader) (*Config, error)
}
