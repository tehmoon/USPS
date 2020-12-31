package main

import (
	"github.com/tehmoon/errors"
)

type AllowRule struct {
	Sources []*Source
	Destinations []*Destination
}

type AllowRuleConfig struct {
	Sources []interface{} // Can either be a string or a Source
	Destinations []interface{} // Can either be a string or a Destination
}

func NewAllowRule(config *Config, rule *AllowRuleConfig) (*AllowRule, error) {
	ar := &AllowRule{
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

				ar.Sources = append(ar.Sources, src)
			//TODO: handle this
			//case []*Source:
			default:
				return nil, errors.Errorf("Allow source #%d type: %T is not supported", i, t)
		}
	}

	if len(rule.Sources) == 0 {
//TODO: add default rule when not found

	for i, destination := range rule.Destinations {
		switch t := destination.(type) {
			case string:
				name, _ := destination.(string)
				dst, found := config.Destinations[name]
				if ! found {
					return nil, errors.Errorf("Allow destination %q #%d is not found", name, i)
				}

				ar.Destinations = append(ar.Destinations, dst)
			//TODO: handle this
			//case []*Destination:
			default:
				return nil, errors.Errorf("Allow destination #%d type: %T is not supported", i, t)
		}
	}

	return ar, nil
}
