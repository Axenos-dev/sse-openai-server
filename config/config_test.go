package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFillConfig(t *testing.T) {
	// Save the current environment variables
	originalEnv := os.Environ()

	// ensure the original environment variables are restored after the test
	defer func() {
		os.Clearenv()
		for _, e := range originalEnv {
			pair := strings.SplitN(e, "=", 2)
			os.Setenv(pair[0], pair[1])
		}
	}()

	t.Run("Fill_Config", func(t *testing.T) {
		os.Setenv("PORT", "8080")
		os.Setenv("OPEN_AI_API_KEY", "test-api-key")
		os.Setenv("MAX_TOKENS", "30")

		err := FillConfig()
		assert.NoError(t, err)

		expectedConfig := config{
			OpenAI: OpenAIConfig{
				ApiKey:    "test-api-key",
				MaxTokens: 30,
			},
			Port: "8080",
		}

		assert.Equal(t, expectedConfig, Config)
	})
}
