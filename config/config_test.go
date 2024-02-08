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
		os.Setenv("POSTGRES_HOST", "127.0.0.1")
		os.Setenv("POSTGRES_PORT", "1234")
		os.Setenv("POSTGRES_DB", "test")
		os.Setenv("POSTGRES_USER", "root")
		os.Setenv("POSTGRES_PASSWORD", "password")

		err := FillConfig()
		assert.NoError(t, err)

		expectedConfig := config{
			OpenAI: OpenAIConfig{
				ApiKey:    "test-api-key",
				MaxTokens: 30,
			},
			PostgreSQL: PostgreSqlConfig{
				Host:     "127.0.0.1",
				Port:     "1234",
				Database: "test",
				User:     "root",
				Password: "password",
			},
			Port: "8080",
		}

		assert.Equal(t, expectedConfig, Config)
	})
}
