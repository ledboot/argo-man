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

func TestLoadConfig(t *testing.T) {
	GetBuildCfg().ConfigFile = "/home/dofun/syncthing/work/github/argo-man/config/simple.toml"
	if err := LoadConfig(&Cfg); err != nil {
		t.Fatal(err)
	}
	fmt.Println(Cfg)
}
