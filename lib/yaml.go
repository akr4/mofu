package mofu

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type Yaml struct{}
type tree map[interface{}]interface{}

func (y *Yaml) Read(r io.Reader) (*Config, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	m := make(tree)
	err = yaml.Unmarshal(bytes, m)
	if err != nil {
		return nil, err
	}

	return visit(*new(Key), &m), nil
}

func visit(key Key, t *tree) *Config {
	c := new(Config)
	visitR(c, key, t)
	return c
}

func visitR(c *Config, key Key, t *tree) {
	for k, v := range *t {
		var kk string = k.(string)
		key2 := append(key, kk)
		switch v.(type) {
		case tree:
			var tt tree = v.(tree)
			visitR(c, key2, &tt)
		case string:
			var vv string = v.(string)
			c.Put(key2, Value(vv))
		default:
			fmt.Printf("unexpected: %v\n", v)
		}
	}
}
