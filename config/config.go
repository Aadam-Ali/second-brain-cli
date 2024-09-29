package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	Root    string
	Inbox   string
	Version string
}

var version string

func GetConfig() Configuration {
	homeDir, _ := os.UserHomeDir()

	root := getEnv("SB", fmt.Sprintf("%s/second-brain", homeDir))
	inbox := fmt.Sprintf("%s/inbox", root)

	return Configuration{
		Root:    root,
		Inbox:   inbox,
		Version: version,
	}
}

func getEnv(key string, defaultValue string) string {
	value, varExists := os.LookupEnv(key)
	if varExists {
		return value
	}
	return defaultValue
}
