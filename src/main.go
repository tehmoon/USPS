package main

import (
//	"./modules/httpProxy"
	"fmt"
	"os"
)

type ConfigFile struct {
	Modules map[string]interface{}
	StaticHosts map[string]string
	Sources map[string]*SourceConfig
	Destinations map[string]*DestinationConfig
//	Allow []*AllowRule
//	Block []*BlockRule
}

type Config struct {
	Sources map[string]*Source
	StaticHosts map[string]string
}

func main() {
	configFile := &ConfigFile{
		StaticHosts: map[string]string{
			"web1": "172.17.0.5",
			"web2": "172.17.0.6",
		},
		Sources: map[string]*SourceConfig{
			"marketing": &SourceConfig{
				Hosts: []string{"mark[1-10]",},
				Modules: []string{"http_proxy"},
			},
			"prod": &SourceConfig{
				CIDR: []string{"10.0.0.0/16", "10.1.0.0/16",},
				Hosts: []string{"web[1-2]",},
				Modules: []string{"http_proxy"},
			},
		},
	}

	config := &Config{
		Sources: make(map[string]*Source),
		StaticHosts: make(map[string]string),
	}

	for name, cfg := range configFile.Sources {
		source, err := NewSource(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating new source for %s: %v\n", name, err.Error())
			os.Exit(2)
		}

		config.Sources[name] = source
	}

	for name, host := range configFile.StaticHosts {
		config.StaticHosts[name] = host
	}

	fmt.Println(config)
	fmt.Println(config.Sources["marketing"])
	fmt.Println(config.Sources["prod"])
}
