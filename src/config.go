package main

import (
	"gopkg.in/yaml.v2"
	"github.com/tehmoon/errors"
	"os"
	"fmt"
)

type ConfigFile struct {
	Modules map[string]interface{} `yaml:"modules"`
	StaticHosts map[string]string `yaml:"static_hosts"`
	Sources map[string]*SourceConfig `yaml:"sources"`
	Destinations map[string]*DestinationConfig `yaml:"destinations"`
	Allow []*RuleConfig `yaml:"allow"`
	Deny []*RuleConfig `yaml:"deny"`
}

type Config struct {
	Sources map[string]*Source
	StaticHosts map[string]string
	Destinations map[string]*Destination
}

func OpenConfigFile(p string) (*Config, error) {
	content, err := os.ReadFile(p)
	if err != nil {
		return nil, errors.Wrap(err, "Err reading config file")
	}

	cf := &ConfigFile{}
	err = yaml.UnmarshalStrict(content, cf)
	if err != nil {
		return nil, errors.Wrap(err, "Err parsing config file to yaml")
	}

	fmt.Println(cf)

	config := &Config{}

	return config, nil
}
