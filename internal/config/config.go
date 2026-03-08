package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config holds all runtime configuration loaded from config.toml.
type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
	Site     SiteConfig     `toml:"site"`
	Admin    AdminConfig    `toml:"admin"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type DatabaseConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}

type SiteConfig struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

type AdminConfig struct {
	Prefix string `toml:"prefix"`
}

// Load reads and parses the TOML config file at path.
// Missing file is not an error — defaults are returned instead.
func Load(path string) (*Config, error) {
	cfg := defaults()

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("config read %s: %w", path, err)
	}

	if err := toml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("config parse %s: %w", path, err)
	}

	return cfg, nil
}

func defaults() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Database: DatabaseConfig{
			Driver: "sqlite",
			DSN:    "./cms.db",
		},
		Site: SiteConfig{
			Name: "My CMS",
			URL:  "http://localhost:8080",
		},
		Admin: AdminConfig{
			Prefix: "/admin",
		},
	}
}
