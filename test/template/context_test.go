package template

import (
	"testing"
)

func TestAdd(t *testing.T) {
	e := map[string]string{"new": "value"}
	c := Context{}
	c.Add("new", "value")
	if c.Values["new"] != e["new"] {
		t.Errorf("Expected %v, got %v", e, c.Values)
	}
}
