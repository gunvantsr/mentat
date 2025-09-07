package main

import (
	"flag"
	"fmt"
)

// Config holds all command-line flags
type Config struct {
	Interactive bool
}

func ParseFlags() *Config {
	config := &Config{}

	// Define command-line flags
	flag.BoolVar(&config.Interactive, "it", false, "Enable interactive mode")

	// Parse the command-line arguments
	flag.Parse()

	// Optional: Print help if needed
	if config.Interactive {
		fmt.Println("Interactive mode enabled!")
	}

	return config
}
