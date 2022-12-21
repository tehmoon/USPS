package main

import (
	"regexp"
	"net"
	"github.com/tehmoon/errors"
)

type Source struct {
	CIDR []*net.IPNet
	Hosts []*regexp.Regexp
	Modules []string
}

type SourceConfig struct {
	CIDR []string `yaml:"cidr"`
	Hosts []string `yaml:"hosts"`
	Modules []string `yaml:"modules"`
}

func NewSource(conf *SourceConfig) (*Source, error) {
	source := &Source{
		CIDR: make([]*net.IPNet, 0),
		Hosts: make([]*regexp.Regexp, 0),
		Modules: make([]string, 0),
	}

	for _, module := range conf.Modules {
		source.Modules = append(source.Modules, module)
	}

	for _, cidr := range conf.CIDR {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, errors.Wrapf(err, "Error parsing cidr for: %q", cidr)
		}

		source.CIDR = append(source.CIDR, ipnet)
	}

	for _, host := range conf.Hosts {
		re, err := regexp.Compile(host)
		if err != nil {
			return nil, errors.Wrapf(err, "Error compiling regex for host: %q", host)
		}

		source.Hosts = append(source.Hosts, re)
	}

	return source, nil
}
