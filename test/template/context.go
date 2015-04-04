package template

type Context struct {
	Values map[string]interface{}
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Add(key string, value interface{}) {
	if c.Values == nil {
		c.Values = make(map[string]interface{}, 1)
	}
	c.Values[key] = value
}
