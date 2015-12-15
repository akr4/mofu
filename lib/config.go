package mofu

type Key []string
type Value string
type Item struct {
	key   Key
	value Value
}

type Config struct {
	data []Item
}

func (c *Config) Put(key Key, value Value) {
	c.data = append(c.data, Item{key, value})
}
