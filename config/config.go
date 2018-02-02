package config

import "path/filepath"

// Config ...
type Config struct {
	Title   string
	Version string
	GameDir string
	FontDir string
}

var config *Config

// Setup ...
func Setup(GameDir string) {

	config = &Config{
		Title:   "Factorigo",
		Version: "0.0.1",
		GameDir: GameDir,
		FontDir: filepath.Join(GameDir, "data", "core", "fonts"),
	}

}

// Get ...
func Get() *Config {
	return config
}
