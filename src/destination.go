package main

import (
	"regexp"
)

type Destination struct {
	Hostnames []*regexp.Regexp
	Ports uint32
}

type DestinationConfig struct {
	Hostnames []string
	Ports []int64 // Maybe have this as port range
}

func NewDestination(conf *DestinationConfig) (*Destination, error) {
	destination := &Destination{
		Hostnames: make([]*regexp.Regexp, 0),
		Ports: make([]uint32, 0),
	}

	for _, hostname := range conf.Hostnames {
		re, err := regexp.Compile(host)
		if err != nil {
			return nil, errors.Wrapf(err, "Error compiling regex for hostname: %q", hostname)
		}

		source.Hostnames = append(source.Hostnames, re)
	}

	return destination, nil
}
