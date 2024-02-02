package config

import (
	"testing"
	"time"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
)

const (
	mockDir              = "./mock"
	mockJson             = mockDir + "/mock.json"
	mockNotSupportedYaml = mockDir + "/mock.yml"
	mockInvalidJson      = mockDir + "/invalid.json"
)

func TestConfig(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected *Config
		isErr    bool
	}{
		{
			name:     "Should load default values when filepath not provided",
			filePath: "",
			expected: &Config{
				OpenAIAPIKey:   "",
				ContextTimeout: 5 * time.Second,
				IgnoreFiles:    "go.mod,go.sum",
			},
			isErr: false,
		},
		{
			name:     "Should load default values when filepath does not exist",
			filePath: "/doesnotexist.json",
			expected: &Config{
				OpenAIAPIKey:   "",
				ContextTimeout: 5 * time.Second,
				IgnoreFiles:    "go.mod,go.sum",
			},
			isErr: false,
		},
		{
			name:     "Should load values from file when filepath exists",
			filePath: mockJson,
			expected: &Config{
				OpenAIAPIKey:   "some-api-key",
				ContextTimeout: 10 * time.Second,
				IgnoreFiles:    "go.mod,go.sum,.env",
			},
			isErr: false,
		},
		{
			name:     "Should return error when filepath exists but is not a JSON file",
			filePath: mockNotSupportedYaml,
			expected: nil,
			isErr:    true,
		},
		{
			name:     "Should return error when config is invalid format",
			filePath: mockInvalidJson,
			expected: nil,
			isErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new koanf instance
			k := koanf.New(".")

			config, err := LoadConfig(k, tc.filePath)
			if tc.isErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tc.expected, config)
			}
		})
	}
}

func TestConversionOfContextTimeout(t *testing.T) {
	expected := time.Duration(10 * time.Second)
	cfg, err := LoadConfig(koanf.New("."), mockJson)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, cfg.ContextTimeout)
}
