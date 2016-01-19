package config

import (
	"errors"

	"github.com/pelletier/go-toml"
)

// Config holds configuration values
type Config struct {
	DbURL    string
	LoginKey []byte
}

var currentConfiguration Config

// LoadConfig loads configuration from a file and validates
func LoadConfig(file string) error {
	toml, err := toml.LoadFile("config.toml")
	if err != nil {
		return err
	}
	currentConfiguration.DbURL = toml.Get("database.connectUrl").(string)
	l := toml.Get("crypto.loginKey").([]interface{})
	currentConfiguration.LoginKey = make([]byte, len(l))
	for i := range l {
		currentConfiguration.LoginKey[i] = byte(l[i].(int64))
	}
	if len(currentConfiguration.DbURL) == 0 || currentConfiguration.LoginKey == nil {
		return errors.New("Invalid configuration")
	}
	return nil
}

// GetConfig returns the loaded configuration
func GetConfig() Config {
	return currentConfiguration
}
