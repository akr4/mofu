package mofu

import "strings"

type Key []string

const mapKeySeparator = "$"

func (k Key) HasPrefix(prefix Key) bool {
	for i := range prefix {
		if k[i] != prefix[i] {
			return false
		}
	}
	return true
}

func (k *Key) mapKey() string {
	return strings.Join(*k, mapKeySeparator)
}

func NewKey(s string) Key {
	var key Key
	for _, x := range strings.Split(s, mapKeySeparator) {
		key = append(key, x)
	}
	return key
}

type Value string

type Config struct {
	data map[string]Value
}

func NewConfig() *Config {
	return &Config{
		data: make(map[string]Value),
	}
}

func (c Config) Put(key *Key, value *Value) Config {
	c.data[key.mapKey()] = *value

	return Config{
		data: c.data,
	}
}

func (c *Config) Filter(prefix Key) Config {
	m := make(map[string]Value)
	for k, v := range c.data {
		kk := NewKey(k)
		if kk.HasPrefix(prefix) {
			m[k] = v
		}
	}
	return Config{data: m}
}

func (c *Config) Merge(that Config) Config {
	m := make(map[string]Value)
	for k, v := range c.data {
		m[k] = v
	}
	for k, v := range that.data {
		m[k] = v
	}
	return Config{data: m}
}
