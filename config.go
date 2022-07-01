package main

import (
	"os"

	"github.com/kovetskiy/ko"
)

type Config struct {
	Token   string `env:"MARK_TOKEN" toml:"token"`
	BaseURL  string `env:"MARK_BASE_URL" toml:"base_url"`
	H1Title  bool   `env:"MARK_H1_TITLE" toml:"h1_title"`
	H1Drop   bool   `env:"MARK_H1_DROP"  toml:"h1_drop"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	err := ko.Load(path, config)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}

		return nil, err
	}

	return config, nil
}
