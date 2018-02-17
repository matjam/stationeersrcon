package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Config stores the configuration for the tool.
type Config struct {
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

func homePath() string {
	var home string

	if runtime.GOOS == "windows" {
		home = os.Getenv("UserProfile")
	} else {
		home = os.Getenv("HOME")
	}

	return home
}

func loadConfig() *Config {
	config := new(Config)

	configPath := *configFlag

	if len(*configFlag) == 0 {
		configPath = filepath.Join(homePath(), ".srcon")
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config
	}

	err = json.Unmarshal(configData, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func saveConfig(config *Config) {
	configPath := *configFlag

	if len(*configFlag) == 0 {
		configPath = filepath.Join(homePath(), ".srcon")
	}

	configData, err := json.MarshalIndent(*config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(configPath), configData, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
