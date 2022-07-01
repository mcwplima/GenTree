package config

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type key int

const configKey key = 98

// NewContext generates a new Context storing the Config into its values.
// Thats helpfull if you need to transfer the config inside the context
// to another function.
func NewContext(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, configKey, c)
}

// FromContext retrieves a *Store previously added to the context by the NewContext func.
// It returns nil if no Store is found.
func FromContext(ctx context.Context) (*Config, bool) {
	s, ok := ctx.Value(configKey).(*Config)
	return s, ok
}

// Config represents the configuration file
type Config struct {
	Database Database
	Core     Core
	Log      Log
}

// Core has the configuration of webservice core.
type Core struct {
	Port int
}

// Log has the configuration of the facility and program.
type Log struct {
	Logger   string
	Facility string
	Program  string
}

// Database has the configuration to connection in the PostgreSQL.
type Database struct {
	Server   string
	Port     int
	Database string
	User     string
	Password string
}

// ReadConfig takes the configurations itens from file.
func ReadConfig(configfile string) *Config {
	var c Config
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	if _, err := toml.DecodeFile(configfile, &c); err != nil {
		log.Fatal(err)
	}
	return &c
}
