package config

import (
	"encoding/json"
	"io"
	"os"
)

func LoadConfig(path string) Config {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	cfg := Config{}
	_ = json.Unmarshal(b, &cfg)
	return cfg
}

func LoadChains(path string) ([]*Chain, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	list := make([]*Chain, 0, 10)
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
