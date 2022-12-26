package config

import (
	"bytes"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"testing"
)

func TestGenFile(t *testing.T) {
	var buff bytes.Buffer
	e := toml.NewEncoder(&buff)
	_ = e.Encode(Cfg)
	fmt.Println(buff.String())
}
