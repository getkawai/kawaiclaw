package internal

import (
	"os"
	"path/filepath"

	"github.com/sipeed/picoclaw/pkg/config"
)

const Logo = "🦞"

// GetKawaiclawHome returns the Kawaiclaw home directory.
// Priority: $KAWAICLAW_HOME > $PICOCLAW_HOME > ~/.kawaiclaw
func GetKawaiclawHome() string {
	if home := os.Getenv("KAWAICLAW_HOME"); home != "" {
		return home
	}
	if home := os.Getenv("PICOCLAW_HOME"); home != "" {
		return home
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".kawaiclaw")
}

func GetConfigPath() string {
	if configPath := os.Getenv("KAWAICLAW_CONFIG"); configPath != "" {
		return configPath
	}
	if configPath := os.Getenv("PICOCLAW_CONFIG"); configPath != "" {
		return configPath
	}
	if home := os.Getenv("KAWAICLAW_HOME"); home != "" {
		return filepath.Join(home, "config.json")
	}
	if home := os.Getenv("PICOCLAW_HOME"); home != "" {
		return filepath.Join(home, "config.json")
	}

	home, _ := os.UserHomeDir()
	defaultConfig := filepath.Join(home, ".kawaiclaw", "config.json")
	if _, err := os.Stat(defaultConfig); err == nil {
		return defaultConfig
	}
	legacyConfig := filepath.Join(home, ".picoclaw", "config.json")
	if _, err := os.Stat(legacyConfig); err == nil {
		return legacyConfig
	}
	return defaultConfig
}

func LoadConfig() (*config.Config, error) {
	return config.LoadConfig(GetConfigPath())
}

// FormatVersion returns the version string with optional git commit
// Deprecated: Use pkg/config.FormatVersion instead
func FormatVersion() string {
	return config.FormatVersion()
}

// FormatBuildInfo returns build time and go version info
// Deprecated: Use pkg/config.FormatBuildInfo instead
func FormatBuildInfo() (string, string) {
	return config.FormatBuildInfo()
}

// GetVersion returns the version string
// Deprecated: Use pkg/config.GetVersion instead
func GetVersion() string {
	return config.GetVersion()
}
