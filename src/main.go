package main

import (
//	"./modules/httpProxy"
	"fmt"
	"os"
	"github.com/tehmoon/errors"
)

func main() {
	cf := "./config.yaml"

	_, err := OpenConfigFile(cf)
	if err != nil {
		err = errors.Wrapf(err, "Error parsing the config file %q", cf)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	//config := &Config{
	//	Sources: make(map[string]*Source),
	//	StaticHosts: make(map[string]string),
	//	Destinations: make(map[string]*Destination),
	//}

	//for name, cfg := range configFile.Sources {
	//	source, err := NewSource(cfg)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error creating new source for %s: %s\n", name, err.Error())
	//		os.Exit(2)
	//	}

	//	config.Sources[name] = source
	//}

	//for name, cfg := range configFile.Destinations {
	//	destination, err := NewDestination(cfg)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error creating new destination for %s: %s\n", name, err.Error())
	//		os.Exit(2)
	//	}

	//	config.Destinations[name] = destination
	//}

	//for name, host := range configFile.StaticHosts {
	//	config.StaticHosts[name] = host
	//}

	//allowRules := make([]*Rule, 0)
	//for i, rule := range configFile.Allow {
	//	r, err := NewRule(config, rule)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error parsing allow rule #%d: %s\n", i+1, err.Error())
	//		os.Exit(2)
	//	}

	//	allowRules = append(allowRules, r)
	//}

	//fmt.Println(config)
	//fmt.Println(config.Sources["marketing"])
	//fmt.Println(config.Sources["prod"])
	//fmt.Println(config.Destinations["facebook_443"])
	//for _, ar := range allowRules {
	//	for _, src := range ar.Sources {
	//		fmt.Printf("Source: %v\n", src)
	//	}
	//	for _, dst := range ar.Destinations {
	//		fmt.Printf("Destination: %v\n", dst)
	//	}
	//}
}
