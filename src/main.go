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
	Allow []*AllowRuleConfig
//	Block []*BlockRule
}

type Config struct {
	Sources map[string]*Source
	StaticHosts map[string]string
	Destinations map[string]*Destination
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
		Destinations: map[string]*DestinationConfig{
			"facebook_443": &DestinationConfig{
				Hostnames: []string{`\.facebook\.com^`, `\.faceblick\.com^`},
				Ports: []int64{443,},
			},
		},
		Allow: []*AllowRuleConfig{
			&AllowRuleConfig{
				Sources: []interface{}{
					"marketing",
				},
				Destinations: []interface{}{
					"facebook_443",
				},
			},
		},
	}

	config := &Config{
		Sources: make(map[string]*Source),
		StaticHosts: make(map[string]string),
		Destinations: make(map[string]*Destination),
	}

	for name, cfg := range configFile.Sources {
		source, err := NewSource(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating new source for %s: %s\n", name, err.Error())
			os.Exit(2)
		}

		config.Sources[name] = source
	}

	for name, cfg := range configFile.Destinations {
		destination, err := NewDestination(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating new destination for %s: %s\n", name, err.Error())
			os.Exit(2)
		}

		config.Destinations[name] = destination
	}

	for name, host := range configFile.StaticHosts {
		config.StaticHosts[name] = host
	}

	allowRules := make([]*AllowRule, 0)
	for i, rule := range configFile.Allow {
		ar, err := NewAllowRule(config, rule)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing allow rule #%d: %s\n", i+1, err.Error())
			os.Exit(2)
		}

		allowRules = append(allowRules, ar)
	}

	fmt.Println(config)
	fmt.Println(config.Sources["marketing"])
	fmt.Println(config.Sources["prod"])
	fmt.Println(config.Destinations["facebook_443"])
	for _, ar := range allowRules {
		for _, src := range ar.Sources {
			fmt.Printf("Source: %v\n", src)
		}
		for _, dst := range ar.Destinations {
			fmt.Printf("Destination: %v\n", dst)
		}
	}
}
