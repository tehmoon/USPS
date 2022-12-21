package main

import (
	"github.com/tehmoon/errors"
)

type Rule struct {
	Sources []*Source
	Destinations []*Destination
}

type RuleConfig struct {
	Sources []interface{} `yaml:"sources"` // Can either be a string or a Source
	Destinations []interface{} `yaml:"destinations"` // Can either be a string or a Destination
	Modules []string `yaml:"modules"`
}

func NewRule(config *Config, rule *RuleConfig) (*Rule, error) {
	r := &Rule{
		Sources: make([]*Source, 0),
		Destinations: make([]*Destination, 0),
	}

	for i, source := range rule.Sources {
		switch t := source.(type) {
			case string:
				name, _ := source.(string)
				src, found := config.Sources[name]
				if ! found {
					return nil, errors.Errorf("Allow source %q #%d is not found", name, i)
				}

				r.Sources = append(r.Sources, src)
			//TODO: handle this
			//case []*Source:
			default:
				return nil, errors.Errorf("Allow source #%d type: %T is not supported", i, t)
		}
	}

	//if len(rule.Sources) == 0 {
//TODO: add default rule when not found

	for i, destination := range rule.Destinations {
		switch t := destination.(type) {
			case string:
				name, _ := destination.(string)
				dst, found := config.Destinations[name]
				if ! found {
					return nil, errors.Errorf("Allow destination %q #%d is not found", name, i)
				}

				r.Destinations = append(r.Destinations, dst)
			//TODO: handle this
			//case []*Destination:
			default:
				return nil, errors.Errorf("Allow destination #%d type: %T is not supported", i, t)
		}
	}

	return r, nil
}
