package config

import (
	"errors"
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
)

func LoadConfig(v interface{}) error {
	if GetBuildCfg().ConfigFile == "" {
		return errors.New("path is empty")
	}
	fi, err := os.Open(GetBuildCfg().ConfigFile)
	if err != nil {
		return err
	}
	defer fi.Close()
	data, err := io.ReadAll(fi)
	if err != nil {
		return err
	}
	if err = toml.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
