package main

import (
	"regexp"
	"github.com/tehmoon/errors"
)

type Destination struct {
	Hostnames []*regexp.Regexp
	Ports []uint16
}

type DestinationConfig struct {
	Hostnames []string `yaml:"hostnames"`
	Ports []int64 `yaml:"ports"` // Maybe have this as port range
}

const (
	MAX_PORT = 1<<16 - 1
)

func NewDestination(conf *DestinationConfig) (*Destination, error) {
	destination := &Destination{
		Hostnames: make([]*regexp.Regexp, 0),
		Ports: make([]uint16, 0),
	}

	for _, hostname := range conf.Hostnames {
		re, err := regexp.Compile(hostname)
		if err != nil {
			return nil, errors.Wrapf(err, "Error compiling regex for hostname: %q", hostname)
		}

		destination.Hostnames = append(destination.Hostnames, re)
	}

	for _, port := range conf.Ports {
		if port < 0 {
			return nil, errors.Errorf("Port %d cannot be lower than 0", port)
		}

		if port > MAX_PORT {
			return nil, errors.Errorf("Port %d cannot be higher than %d", port, MAX_PORT)
		}

		destination.Ports = append(destination.Ports, uint16(port))
	}

	return destination, nil
}
