package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Config struct for webapp config TODO вынести в internal
type Config struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"database_name"`
	} `yaml:"postgres"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
