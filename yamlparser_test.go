package gosk

import (
	"github.com/scottkiss/go-gypsy/yaml"
	"testing"
)

func TestParse(t *testing.T) {
	yp := new(YamlParser)
	configs := yp.parse("root")
	if configs == nil {
		t.Error("yaml parser error!")
	}
	cfg := configs["config.yml"]

	var c = cfg.(*yaml.File)
	c.Get("meta")

}
