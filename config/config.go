package config

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

const EnvFileName = "env.json"

type Config struct {
	OpenAIAPIKey   string        `koanf:"OPENAI_API_KEY"`
	ContextTimeout time.Duration `koanf:"CONTEXT_TIMEOUT"`
	// IgnoreFiles is a comma-separated list of file names to ignore.
	IgnoreFiles string `koanf:"IGNORE_FILES"`
}

// DefaultConfig is the default configuration for the application.
var defaultConfig = Config{
	OpenAIAPIKey:   "",
	ContextTimeout: 5 * time.Second,
	IgnoreFiles:    "go.mod,go.sum",
}

// LoadConfig loads the configuration from the given path.
// If the path is empty, the default configuration is used.
func LoadConfig(k *koanf.Koanf, filePath string) (*Config, error) {
	if err := loadDefaultConfig(k); err != nil {
		return nil, err
	}

	if filePath != "" && fileExists(filePath) {
		// loadFileConfig currently supports JSON files.
		if err := loadFileConfig(k, filePath); err != nil {
			return nil, err
		}
	}

	config := &Config{}
	if err := k.Unmarshal("", config); err != nil {
		return nil, err
	}

	return config, nil
}

// loadDefaultConfig loads the default configuration
// into the given koanf instance.
func loadDefaultConfig(k *koanf.Koanf) error {
	return k.Load(structs.Provider(defaultConfig, "koanf"), nil)
}

// loadFileConfig loads the configuration from the given file
// it currently supports JSON files.
func loadFileConfig(k *koanf.Koanf, filePath string) error {
	if !strings.HasSuffix(filePath, ".json") {
		return errors.New("invalid file extension, only .json files are supported")
	}

	if err := k.Load(file.Provider(filePath), json.Parser()); err != nil {
		return err
	}

	return nil
}

// fileExists checks if the given file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}
